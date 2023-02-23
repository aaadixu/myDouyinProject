// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	video "douyinProject/video/kitex_gen/video"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FeedMethod":          kitex.NewMethodInfo(feedMethodHandler, newVideoServiceFeedMethodArgs, newVideoServiceFeedMethodResult, false),
		"PublishActionMethod": kitex.NewMethodInfo(publishActionMethodHandler, newVideoServicePublishActionMethodArgs, newVideoServicePublishActionMethodResult, false),
		"PublishListMethod":   kitex.NewMethodInfo(publishListMethodHandler, newVideoServicePublishListMethodArgs, newVideoServicePublishListMethodResult, false),
		"FavoriteCountMethod": kitex.NewMethodInfo(favoriteCountMethodHandler, newVideoServiceFavoriteCountMethodArgs, newVideoServiceFavoriteCountMethodResult, false),
		"CommentCountMethod":  kitex.NewMethodInfo(commentCountMethodHandler, newVideoServiceCommentCountMethodArgs, newVideoServiceCommentCountMethodResult, false),
		"VideoListMethod":     kitex.NewMethodInfo(videoListMethodHandler, newVideoServiceVideoListMethodArgs, newVideoServiceVideoListMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
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

func feedMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceFeedMethodArgs)
	realResult := result.(*video.VideoServiceFeedMethodResult)
	success, err := handler.(video.VideoService).FeedMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFeedMethodArgs() interface{} {
	return video.NewVideoServiceFeedMethodArgs()
}

func newVideoServiceFeedMethodResult() interface{} {
	return video.NewVideoServiceFeedMethodResult()
}

func publishActionMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServicePublishActionMethodArgs)
	realResult := result.(*video.VideoServicePublishActionMethodResult)
	success, err := handler.(video.VideoService).PublishActionMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePublishActionMethodArgs() interface{} {
	return video.NewVideoServicePublishActionMethodArgs()
}

func newVideoServicePublishActionMethodResult() interface{} {
	return video.NewVideoServicePublishActionMethodResult()
}

func publishListMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServicePublishListMethodArgs)
	realResult := result.(*video.VideoServicePublishListMethodResult)
	success, err := handler.(video.VideoService).PublishListMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePublishListMethodArgs() interface{} {
	return video.NewVideoServicePublishListMethodArgs()
}

func newVideoServicePublishListMethodResult() interface{} {
	return video.NewVideoServicePublishListMethodResult()
}

func favoriteCountMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceFavoriteCountMethodArgs)
	realResult := result.(*video.VideoServiceFavoriteCountMethodResult)
	success, err := handler.(video.VideoService).FavoriteCountMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFavoriteCountMethodArgs() interface{} {
	return video.NewVideoServiceFavoriteCountMethodArgs()
}

func newVideoServiceFavoriteCountMethodResult() interface{} {
	return video.NewVideoServiceFavoriteCountMethodResult()
}

func commentCountMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceCommentCountMethodArgs)
	realResult := result.(*video.VideoServiceCommentCountMethodResult)
	success, err := handler.(video.VideoService).CommentCountMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceCommentCountMethodArgs() interface{} {
	return video.NewVideoServiceCommentCountMethodArgs()
}

func newVideoServiceCommentCountMethodResult() interface{} {
	return video.NewVideoServiceCommentCountMethodResult()
}

func videoListMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceVideoListMethodArgs)
	realResult := result.(*video.VideoServiceVideoListMethodResult)
	success, err := handler.(video.VideoService).VideoListMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoListMethodArgs() interface{} {
	return video.NewVideoServiceVideoListMethodArgs()
}

func newVideoServiceVideoListMethodResult() interface{} {
	return video.NewVideoServiceVideoListMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FeedMethod(ctx context.Context, request *video.FeedReq) (r *video.FeedResp, err error) {
	var _args video.VideoServiceFeedMethodArgs
	_args.Request = request
	var _result video.VideoServiceFeedMethodResult
	if err = p.c.Call(ctx, "FeedMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishActionMethod(ctx context.Context, request *video.PublishActionReq) (r *video.PublishActionResp, err error) {
	var _args video.VideoServicePublishActionMethodArgs
	_args.Request = request
	var _result video.VideoServicePublishActionMethodResult
	if err = p.c.Call(ctx, "PublishActionMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishListMethod(ctx context.Context, request *video.PublishListReq) (r *video.PublishListResp, err error) {
	var _args video.VideoServicePublishListMethodArgs
	_args.Request = request
	var _result video.VideoServicePublishListMethodResult
	if err = p.c.Call(ctx, "PublishListMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteCountMethod(ctx context.Context, request *video.FavoriteCountReq) (r *video.FavoriteCountResp, err error) {
	var _args video.VideoServiceFavoriteCountMethodArgs
	_args.Request = request
	var _result video.VideoServiceFavoriteCountMethodResult
	if err = p.c.Call(ctx, "FavoriteCountMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentCountMethod(ctx context.Context, request *video.CommentCountReq) (r *video.CommentCountResp, err error) {
	var _args video.VideoServiceCommentCountMethodArgs
	_args.Request = request
	var _result video.VideoServiceCommentCountMethodResult
	if err = p.c.Call(ctx, "CommentCountMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoListMethod(ctx context.Context, request *video.VideoListReq) (r *video.VideoListResp, err error) {
	var _args video.VideoServiceVideoListMethodArgs
	_args.Request = request
	var _result video.VideoServiceVideoListMethodResult
	if err = p.c.Call(ctx, "VideoListMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
