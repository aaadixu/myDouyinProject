// Code generated by hertz generator.

package http

import (
	"context"
	http "douyinProject/http/biz/model/http"
	jwt "douyinProject/http/biz/mw"
	rpc "douyinProject/http/biz/rpc"
	"douyinProject/http/kitex_gen/httprpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"io/ioutil"
)

// FeedMethod .
// @router /douyin/feed/ [GET]
func FeedMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req http.FeedReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(http.FeedResp)

	res, err := jwt.AuthToken(ctx, c)
	//if err != nil || res == -1 { // token校验失败
	//	resp.StatusCode = 2
	//	resp.StatusMsg = "auth fail"
	//	c.JSON(consts.StatusBadRequest, resp)
	//	return
	//}

	rpcResp, err := rpc.VideoClient.FeedMethod(ctx, &httprpc.FeedReq{LatestTime: req.LatestTime, UserId: res})
	if err != nil {
		resp.StatusCode = 3
		resp.StatusMsg = "get video list fail"
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	resp.StatusCode = rpcResp.StatusCode
	resp.StatusMsg = rpcResp.StatusMsg

	vList, err := rpc.PackVideos(rpcResp.VideoList)

	resp.VideoList = vList

	resp.NextTime = rpcResp.NextTime
	c.JSON(consts.StatusOK, resp)
}

// PublishActionMethod .
// @router /douyin/publish/action/ [POST]
func PublishActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error

	//token := string(c.FormValue("token"))
	title := string(c.FormValue("title"))
	fileHeader, err := c.FormFile("data")
	if err != nil {
		panic(err)
	}
	open, err := fileHeader.Open()
	if err != nil {
		panic(err)
	}
	// 读取文件到字节数组
	fileRaw, err := ioutil.ReadAll(open)

	resp := new(http.PublishActionResp)

	res, err := jwt.AuthToken(ctx, c)
	if err != nil || res == -1 { // token校验失败
		resp.StatusCode = 2
		msg := "auth fail"
		resp.StatusMsg = &msg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	rpcResp, err := rpc.VideoClient.PublishActionMethod(ctx, &httprpc.PublishActionReq{
		UserId: res,
		Data:   fileRaw,
		Title:  title,
	})
	if err != nil {
		resp.StatusCode = 2
		msg := "auth fail"
		resp.StatusMsg = &msg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	resp.StatusCode = rpcResp.StatusCode
	resp.StatusMsg = rpcResp.StatusMsg

	if err != nil {
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// PublishListMethod .
// @router /douyin/publish/list/ [GET]
func PublishListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req http.PublishListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(http.PublishListResp)

	res, err := jwt.AuthToken(ctx, c)
	if err != nil || res == -1 { // token校验失败
		msg := "auth fail"
		resp.StatusCode = 2
		resp.StatusMsg = &msg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	rpcResp, err := rpc.VideoClient.PublishListMethod(ctx, &httprpc.PublishListReq{UserId: req.UserID})
	resp.StatusCode = rpcResp.StatusCode
	resp.StatusMsg = rpcResp.StatusMsg

	vs, err := rpc.PackVideos(rpcResp.VideoList)

	resp.VideoList = vs
	c.JSON(consts.StatusOK, resp)
}
