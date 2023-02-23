package consts

import "time"

var (
	HttpServicePort = ":9999"
	HttpServiceHost = ""

	// 授权认证
	JwtIdentityKey = "userid"
	JwtRealm       = "test zone"
	JwtKey         = "ujdfiweuophhfre432538f43hfb8932rf98234uh89r"
	JwtTimeOut     = time.Hour * 24
	JwtMaxRefresh  = time.Hour
	JwtTokenLookup = "query: token,param: token,form: token"

	// 远程rpc服务
	ETCDAddr            = ":2379"
	ActionServiceName   = "actionservice"
	ChatServiceName     = "chatservice"
	RelationServiceName = "relationservice"
	UserServiceName     = "userservice"
	VideoServiceName    = "videoservice"
)
