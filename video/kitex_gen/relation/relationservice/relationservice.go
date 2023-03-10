// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"
	relation "douyinProject/video/kitex_gen/relation"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return relationServiceServiceInfo
}

var relationServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "RelationService"
	handlerType := (*relation.RelationService)(nil)
	methods := map[string]kitex.MethodInfo{
		"IsFollowingMethod": kitex.NewMethodInfo(isFollowingMethodHandler, newRelationServiceIsFollowingMethodArgs, newRelationServiceIsFollowingMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "relation",
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

func isFollowingMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceIsFollowingMethodArgs)
	realResult := result.(*relation.RelationServiceIsFollowingMethodResult)
	success, err := handler.(relation.RelationService).IsFollowingMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceIsFollowingMethodArgs() interface{} {
	return relation.NewRelationServiceIsFollowingMethodArgs()
}

func newRelationServiceIsFollowingMethodResult() interface{} {
	return relation.NewRelationServiceIsFollowingMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) IsFollowingMethod(ctx context.Context, request *relation.IsFollowingReq) (r *relation.IsFollowingResp, err error) {
	var _args relation.RelationServiceIsFollowingMethodArgs
	_args.Request = request
	var _result relation.RelationServiceIsFollowingMethodResult
	if err = p.c.Call(ctx, "IsFollowingMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
