package utils

import (
	"context"
	"douyinProject/http/kitex_gen/httprpc"

	"douyinProject/http/biz/consts"
	bean "douyinProject/http/biz/model/http"
	"douyinProject/http/biz/rpc"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	//"github.com/golang-jwt/jwt/v4"
	"github.com/hertz-contrib/jwt"
	"log"
	"net/http"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       consts.JwtRealm,
		Key:         []byte(consts.JwtKey),
		Timeout:     consts.JwtTimeOut,
		MaxRefresh:  consts.JwtMaxRefresh,
		TokenLookup: consts.JwtTokenLookup,

		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct bean.LoginReq

			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			userRes, err := rpc.UserClient.LoginMethod(ctx, &httprpc.LoginReq{
				Username: loginStruct.Name,
				Password: loginStruct.Password,
			})
			if err != nil {
				return nil, err
			}
			if userRes.UserId == 0 {
				return nil, errors.New("user already exists or wrong password")
			}
			return &bean.User{
				ID: userRes.UserId,
			}, nil
		},
		IdentityKey: consts.JwtIdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &bean.User{
				ID: claims[consts.JwtIdentityKey].(int64),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {

			if v, ok := data.(*bean.User); ok {
				return jwt.MapClaims{
					consts.JwtIdentityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := JwtMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
}

// 时间过期之后，map会置空
func AuthToken(ctx context.Context, c *app.RequestContext) (int64, error) {
	res, err := JwtMiddleware.GetClaimsFromJWT(ctx, c)

	if err != nil {
		return -1, err
	}
	//fmt.Println(res["exp"].(float64)) // 过期时间
	//fmt.Println(res["orig_iat"].(float64)) // 开始时间

	value, ok := res[consts.JwtIdentityKey]
	if ok {
		return int64(value.(float64)), nil
	}
	return -1, err
}
