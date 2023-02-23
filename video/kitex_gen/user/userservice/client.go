// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	user "douyinProject/video/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	InfoMethod(ctx context.Context, request *user.InfoReq, callOptions ...callopt.Option) (r *user.InfoResp, err error)
	AddWorkNumMethod(ctx context.Context, request *user.AddWorkNumReq, callOptions ...callopt.Option) (r *user.AddWorkNumResp, err error)
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
	return &kUserServiceClient{
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

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) InfoMethod(ctx context.Context, request *user.InfoReq, callOptions ...callopt.Option) (r *user.InfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InfoMethod(ctx, request)
}

func (p *kUserServiceClient) AddWorkNumMethod(ctx context.Context, request *user.AddWorkNumReq, callOptions ...callopt.Option) (r *user.AddWorkNumResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddWorkNumMethod(ctx, request)
}
