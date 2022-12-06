// Code generated by goctl. DO NOT EDIT.
// Source: post.proto

package post

import (
	"context"

	"github.com/xh-polaris/meowchat-post-rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddPostReq      = pb.AddPostReq
	AddPostResp     = pb.AddPostResp
	DeletePostReq   = pb.DeletePostReq
	DeletePostResp  = pb.DeletePostResp
	GetManyPostReq  = pb.GetManyPostReq
	GetManyPostResp = pb.GetManyPostResp
	GetPostReq      = pb.GetPostReq
	GetPostResp     = pb.GetPostResp
	Post            = pb.Post
	Tag             = pb.Tag
	UpdatePostReq   = pb.UpdatePostReq
	UpdatePostResp  = pb.UpdatePostResp
	User            = pb.User

	Post interface {
		GetManyPost(ctx context.Context, in *GetManyPostReq, opts ...grpc.CallOption) (*GetManyPostResp, error)
		GetPost(ctx context.Context, in *GetPostReq, opts ...grpc.CallOption) (*GetPostResp, error)
		AddPost(ctx context.Context, in *AddPostReq, opts ...grpc.CallOption) (*AddPostResp, error)
		UpdatePost(ctx context.Context, in *UpdatePostReq, opts ...grpc.CallOption) (*UpdatePostResp, error)
		DeletePost(ctx context.Context, in *DeletePostReq, opts ...grpc.CallOption) (*DeletePostResp, error)
	}

	defaultPost struct {
		cli zrpc.Client
	}
)

func NewPost(cli zrpc.Client) Post {
	return &defaultPost{
		cli: cli,
	}
}

func (m *defaultPost) GetManyPost(ctx context.Context, in *GetManyPostReq, opts ...grpc.CallOption) (*GetManyPostResp, error) {
	client := pb.NewPostClient(m.cli.Conn())
	return client.GetManyPost(ctx, in, opts...)
}

func (m *defaultPost) GetPost(ctx context.Context, in *GetPostReq, opts ...grpc.CallOption) (*GetPostResp, error) {
	client := pb.NewPostClient(m.cli.Conn())
	return client.GetPost(ctx, in, opts...)
}

func (m *defaultPost) AddPost(ctx context.Context, in *AddPostReq, opts ...grpc.CallOption) (*AddPostResp, error) {
	client := pb.NewPostClient(m.cli.Conn())
	return client.AddPost(ctx, in, opts...)
}

func (m *defaultPost) UpdatePost(ctx context.Context, in *UpdatePostReq, opts ...grpc.CallOption) (*UpdatePostResp, error) {
	client := pb.NewPostClient(m.cli.Conn())
	return client.UpdatePost(ctx, in, opts...)
}

func (m *defaultPost) DeletePost(ctx context.Context, in *DeletePostReq, opts ...grpc.CallOption) (*DeletePostResp, error) {
	client := pb.NewPostClient(m.cli.Conn())
	return client.DeletePost(ctx, in, opts...)
}