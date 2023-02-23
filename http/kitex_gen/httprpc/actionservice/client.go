// Code generated by Kitex v0.4.4. DO NOT EDIT.

package actionservice

import (
	"context"
	httprpc "douyinProject/http/kitex_gen/httprpc"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FavoriteActionMethod(ctx context.Context, request *httprpc.FavoriteActionReq, callOptions ...callopt.Option) (r *httprpc.FavoriteActionResp, err error)
	FavoriteListMethod(ctx context.Context, request *httprpc.FavoriteListReq, callOptions ...callopt.Option) (r *httprpc.FavoriteListResp, err error)
	CommentActionMethod(ctx context.Context, request *httprpc.CommentActionReq, callOptions ...callopt.Option) (r *httprpc.CommentActionResp, err error)
	CommentListMethod(ctx context.Context, request *httprpc.CommentListReq, callOptions ...callopt.Option) (r *httprpc.CommentListResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kActionServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kActionServiceClient struct {
	*kClient
}

func (p *kActionServiceClient) FavoriteActionMethod(ctx context.Context, request *httprpc.FavoriteActionReq, callOptions ...callopt.Option) (r *httprpc.FavoriteActionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteActionMethod(ctx, request)
}

func (p *kActionServiceClient) FavoriteListMethod(ctx context.Context, request *httprpc.FavoriteListReq, callOptions ...callopt.Option) (r *httprpc.FavoriteListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteListMethod(ctx, request)
}

func (p *kActionServiceClient) CommentActionMethod(ctx context.Context, request *httprpc.CommentActionReq, callOptions ...callopt.Option) (r *httprpc.CommentActionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentActionMethod(ctx, request)
}

func (p *kActionServiceClient) CommentListMethod(ctx context.Context, request *httprpc.CommentListReq, callOptions ...callopt.Option) (r *httprpc.CommentListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentListMethod(ctx, request)
}