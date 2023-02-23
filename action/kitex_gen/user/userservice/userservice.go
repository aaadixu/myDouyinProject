// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	user "douyinProject/action/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteCountMethod":  kitex.NewMethodInfo(favoriteCountMethodHandler, newUserServiceFavoriteCountMethodArgs, newUserServiceFavoriteCountMethodResult, false),
		"TotalFavoritedMethod": kitex.NewMethodInfo(totalFavoritedMethodHandler, newUserServiceTotalFavoritedMethodArgs, newUserServiceTotalFavoritedMethodResult, false),
		"InfoMethod":           kitex.NewMethodInfo(infoMethodHandler, newUserServiceInfoMethodArgs, newUserServiceInfoMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func favoriteCountMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceFavoriteCountMethodArgs)
	realResult := result.(*user.UserServiceFavoriteCountMethodResult)
	success, err := handler.(user.UserService).FavoriteCountMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceFavoriteCountMethodArgs() interface{} {
	return user.NewUserServiceFavoriteCountMethodArgs()
}

func newUserServiceFavoriteCountMethodResult() interface{} {
	return user.NewUserServiceFavoriteCountMethodResult()
}

func totalFavoritedMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceTotalFavoritedMethodArgs)
	realResult := result.(*user.UserServiceTotalFavoritedMethodResult)
	success, err := handler.(user.UserService).TotalFavoritedMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceTotalFavoritedMethodArgs() interface{} {
	return user.NewUserServiceTotalFavoritedMethodArgs()
}

func newUserServiceTotalFavoritedMethodResult() interface{} {
	return user.NewUserServiceTotalFavoritedMethodResult()
}

func infoMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceInfoMethodArgs)
	realResult := result.(*user.UserServiceInfoMethodResult)
	success, err := handler.(user.UserService).InfoMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceInfoMethodArgs() interface{} {
	return user.NewUserServiceInfoMethodArgs()
}

func newUserServiceInfoMethodResult() interface{} {
	return user.NewUserServiceInfoMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteCountMethod(ctx context.Context, request *user.FavoriteCountReq) (r *user.FavoriteCountResp, err error) {
	var _args user.UserServiceFavoriteCountMethodArgs
	_args.Request = request
	var _result user.UserServiceFavoriteCountMethodResult
	if err = p.c.Call(ctx, "FavoriteCountMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TotalFavoritedMethod(ctx context.Context, request *user.TotalFavoritedReq) (r *user.TotalFavoritedResp, err error) {
	var _args user.UserServiceTotalFavoritedMethodArgs
	_args.Request = request
	var _result user.UserServiceTotalFavoritedMethodResult
	if err = p.c.Call(ctx, "TotalFavoritedMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) InfoMethod(ctx context.Context, request *user.InfoReq) (r *user.InfoResp, err error) {
	var _args user.UserServiceInfoMethodArgs
	_args.Request = request
	var _result user.UserServiceInfoMethodResult
	if err = p.c.Call(ctx, "InfoMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
