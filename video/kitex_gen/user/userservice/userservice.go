// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	user "douyinProject/video/kitex_gen/user"
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
		"InfoMethod":       kitex.NewMethodInfo(infoMethodHandler, newUserServiceInfoMethodArgs, newUserServiceInfoMethodResult, false),
		"AddWorkNumMethod": kitex.NewMethodInfo(addWorkNumMethodHandler, newUserServiceAddWorkNumMethodArgs, newUserServiceAddWorkNumMethodResult, false),
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

func addWorkNumMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceAddWorkNumMethodArgs)
	realResult := result.(*user.UserServiceAddWorkNumMethodResult)
	success, err := handler.(user.UserService).AddWorkNumMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceAddWorkNumMethodArgs() interface{} {
	return user.NewUserServiceAddWorkNumMethodArgs()
}

func newUserServiceAddWorkNumMethodResult() interface{} {
	return user.NewUserServiceAddWorkNumMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
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

func (p *kClient) AddWorkNumMethod(ctx context.Context, request *user.AddWorkNumReq) (r *user.AddWorkNumResp, err error) {
	var _args user.UserServiceAddWorkNumMethodArgs
	_args.Request = request
	var _result user.UserServiceAddWorkNumMethodResult
	if err = p.c.Call(ctx, "AddWorkNumMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}