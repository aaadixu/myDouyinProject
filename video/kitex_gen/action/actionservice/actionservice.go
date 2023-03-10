// Code generated by Kitex v0.4.4. DO NOT EDIT.

package actionservice

import (
	"context"
	action "douyinProject/video/kitex_gen/action"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return actionServiceServiceInfo
}

var actionServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ActionService"
	handlerType := (*action.ActionService)(nil)
	methods := map[string]kitex.MethodInfo{
		"IsUserFavoriteVideoMethod": kitex.NewMethodInfo(isUserFavoriteVideoMethodHandler, newActionServiceIsUserFavoriteVideoMethodArgs, newActionServiceIsUserFavoriteVideoMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "action",
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

func isUserFavoriteVideoMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*action.ActionServiceIsUserFavoriteVideoMethodArgs)
	realResult := result.(*action.ActionServiceIsUserFavoriteVideoMethodResult)
	success, err := handler.(action.ActionService).IsUserFavoriteVideoMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newActionServiceIsUserFavoriteVideoMethodArgs() interface{} {
	return action.NewActionServiceIsUserFavoriteVideoMethodArgs()
}

func newActionServiceIsUserFavoriteVideoMethodResult() interface{} {
	return action.NewActionServiceIsUserFavoriteVideoMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) IsUserFavoriteVideoMethod(ctx context.Context, request *action.IsUserFavoriteVideoReq) (r *action.IsUserFavoriteVideoResp, err error) {
	var _args action.ActionServiceIsUserFavoriteVideoMethodArgs
	_args.Request = request
	var _result action.ActionServiceIsUserFavoriteVideoMethodResult
	if err = p.c.Call(ctx, "IsUserFavoriteVideoMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
