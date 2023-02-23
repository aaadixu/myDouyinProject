// Code generated by Kitex v0.4.4. DO NOT EDIT.

package actionservice

import (
	"context"
	httprpc "douyinProject/http/kitex_gen/httprpc"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return actionServiceServiceInfo
}

var actionServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ActionService"
	handlerType := (*httprpc.ActionService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteActionMethod": kitex.NewMethodInfo(favoriteActionMethodHandler, newActionServiceFavoriteActionMethodArgs, newActionServiceFavoriteActionMethodResult, false),
		"FavoriteListMethod":   kitex.NewMethodInfo(favoriteListMethodHandler, newActionServiceFavoriteListMethodArgs, newActionServiceFavoriteListMethodResult, false),
		"CommentActionMethod":  kitex.NewMethodInfo(commentActionMethodHandler, newActionServiceCommentActionMethodArgs, newActionServiceCommentActionMethodResult, false),
		"CommentListMethod":    kitex.NewMethodInfo(commentListMethodHandler, newActionServiceCommentListMethodArgs, newActionServiceCommentListMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "httprpc",
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

func favoriteActionMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*httprpc.ActionServiceFavoriteActionMethodArgs)
	realResult := result.(*httprpc.ActionServiceFavoriteActionMethodResult)
	success, err := handler.(httprpc.ActionService).FavoriteActionMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newActionServiceFavoriteActionMethodArgs() interface{} {
	return httprpc.NewActionServiceFavoriteActionMethodArgs()
}

func newActionServiceFavoriteActionMethodResult() interface{} {
	return httprpc.NewActionServiceFavoriteActionMethodResult()
}

func favoriteListMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*httprpc.ActionServiceFavoriteListMethodArgs)
	realResult := result.(*httprpc.ActionServiceFavoriteListMethodResult)
	success, err := handler.(httprpc.ActionService).FavoriteListMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newActionServiceFavoriteListMethodArgs() interface{} {
	return httprpc.NewActionServiceFavoriteListMethodArgs()
}

func newActionServiceFavoriteListMethodResult() interface{} {
	return httprpc.NewActionServiceFavoriteListMethodResult()
}

func commentActionMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*httprpc.ActionServiceCommentActionMethodArgs)
	realResult := result.(*httprpc.ActionServiceCommentActionMethodResult)
	success, err := handler.(httprpc.ActionService).CommentActionMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newActionServiceCommentActionMethodArgs() interface{} {
	return httprpc.NewActionServiceCommentActionMethodArgs()
}

func newActionServiceCommentActionMethodResult() interface{} {
	return httprpc.NewActionServiceCommentActionMethodResult()
}

func commentListMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*httprpc.ActionServiceCommentListMethodArgs)
	realResult := result.(*httprpc.ActionServiceCommentListMethodResult)
	success, err := handler.(httprpc.ActionService).CommentListMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newActionServiceCommentListMethodArgs() interface{} {
	return httprpc.NewActionServiceCommentListMethodArgs()
}

func newActionServiceCommentListMethodResult() interface{} {
	return httprpc.NewActionServiceCommentListMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteActionMethod(ctx context.Context, request *httprpc.FavoriteActionReq) (r *httprpc.FavoriteActionResp, err error) {
	var _args httprpc.ActionServiceFavoriteActionMethodArgs
	_args.Request = request
	var _result httprpc.ActionServiceFavoriteActionMethodResult
	if err = p.c.Call(ctx, "FavoriteActionMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteListMethod(ctx context.Context, request *httprpc.FavoriteListReq) (r *httprpc.FavoriteListResp, err error) {
	var _args httprpc.ActionServiceFavoriteListMethodArgs
	_args.Request = request
	var _result httprpc.ActionServiceFavoriteListMethodResult
	if err = p.c.Call(ctx, "FavoriteListMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentActionMethod(ctx context.Context, request *httprpc.CommentActionReq) (r *httprpc.CommentActionResp, err error) {
	var _args httprpc.ActionServiceCommentActionMethodArgs
	_args.Request = request
	var _result httprpc.ActionServiceCommentActionMethodResult
	if err = p.c.Call(ctx, "CommentActionMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentListMethod(ctx context.Context, request *httprpc.CommentListReq) (r *httprpc.CommentListResp, err error) {
	var _args httprpc.ActionServiceCommentListMethodArgs
	_args.Request = request
	var _result httprpc.ActionServiceCommentListMethodResult
	if err = p.c.Call(ctx, "CommentListMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}