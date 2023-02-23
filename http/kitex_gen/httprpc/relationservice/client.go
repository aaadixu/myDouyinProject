// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"
	httprpc "douyinProject/http/kitex_gen/httprpc"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RelationActionMethod(ctx context.Context, request *httprpc.RelationActionReq, callOptions ...callopt.Option) (r *httprpc.RelationActionResp, err error)
	FollowListMethod(ctx context.Context, request *httprpc.FollowListReq, callOptions ...callopt.Option) (r *httprpc.FollowListResp, err error)
	FollowerListMethod(ctx context.Context, request *httprpc.FollowerListReq, callOptions ...callopt.Option) (r *httprpc.FollowerListResp, err error)
	FriendListMethod(ctx context.Context, request *httprpc.FriendListReq, callOptions ...callopt.Option) (r *httprpc.FriendListResp, err error)
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
	return &kRelationServiceClient{
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

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) RelationActionMethod(ctx context.Context, request *httprpc.RelationActionReq, callOptions ...callopt.Option) (r *httprpc.RelationActionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationActionMethod(ctx, request)
}

func (p *kRelationServiceClient) FollowListMethod(ctx context.Context, request *httprpc.FollowListReq, callOptions ...callopt.Option) (r *httprpc.FollowListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowListMethod(ctx, request)
}

func (p *kRelationServiceClient) FollowerListMethod(ctx context.Context, request *httprpc.FollowerListReq, callOptions ...callopt.Option) (r *httprpc.FollowerListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowerListMethod(ctx, request)
}

func (p *kRelationServiceClient) FriendListMethod(ctx context.Context, request *httprpc.FriendListReq, callOptions ...callopt.Option) (r *httprpc.FriendListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FriendListMethod(ctx, request)
}
