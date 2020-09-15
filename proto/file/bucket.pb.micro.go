// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/file/bucket.proto

package omo_msa_file

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Bucket service

type BucketService interface {
	// 创建一个存储桶
	Make(ctx context.Context, in *BucketMakeRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 列举存储桶
	List(ctx context.Context, in *BucketListRequest, opts ...client.CallOption) (*BucketListResponse, error)
	// 删除一个存储桶
	Remove(ctx context.Context, in *BucketRemoveRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 获取一个存储桶信息
	Get(ctx context.Context, in *BucketGetRequest, opts ...client.CallOption) (*BucketGetResponse, error)
	// 更新一个存储桶的引擎
	UpdateEngine(ctx context.Context, in *BucketUpdateEngineRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 更新一个存储桶的容量
	UpdateCapacity(ctx context.Context, in *BucketUpdateCapacityRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 重置一个存储桶的访问令牌
	ResetToken(ctx context.Context, in *BucketResetTokenRequest, opts ...client.CallOption) (*BlankResponse, error)
}

type bucketService struct {
	c    client.Client
	name string
}

func NewBucketService(name string, c client.Client) BucketService {
	return &bucketService{
		c:    c,
		name: name,
	}
}

func (c *bucketService) Make(ctx context.Context, in *BucketMakeRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Bucket.Make", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketService) List(ctx context.Context, in *BucketListRequest, opts ...client.CallOption) (*BucketListResponse, error) {
	req := c.c.NewRequest(c.name, "Bucket.List", in)
	out := new(BucketListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketService) Remove(ctx context.Context, in *BucketRemoveRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Bucket.Remove", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketService) Get(ctx context.Context, in *BucketGetRequest, opts ...client.CallOption) (*BucketGetResponse, error) {
	req := c.c.NewRequest(c.name, "Bucket.Get", in)
	out := new(BucketGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketService) UpdateEngine(ctx context.Context, in *BucketUpdateEngineRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Bucket.UpdateEngine", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketService) UpdateCapacity(ctx context.Context, in *BucketUpdateCapacityRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Bucket.UpdateCapacity", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketService) ResetToken(ctx context.Context, in *BucketResetTokenRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Bucket.ResetToken", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Bucket service

type BucketHandler interface {
	// 创建一个存储桶
	Make(context.Context, *BucketMakeRequest, *BlankResponse) error
	// 列举存储桶
	List(context.Context, *BucketListRequest, *BucketListResponse) error
	// 删除一个存储桶
	Remove(context.Context, *BucketRemoveRequest, *BlankResponse) error
	// 获取一个存储桶信息
	Get(context.Context, *BucketGetRequest, *BucketGetResponse) error
	// 更新一个存储桶的引擎
	UpdateEngine(context.Context, *BucketUpdateEngineRequest, *BlankResponse) error
	// 更新一个存储桶的容量
	UpdateCapacity(context.Context, *BucketUpdateCapacityRequest, *BlankResponse) error
	// 重置一个存储桶的访问令牌
	ResetToken(context.Context, *BucketResetTokenRequest, *BlankResponse) error
}

func RegisterBucketHandler(s server.Server, hdlr BucketHandler, opts ...server.HandlerOption) error {
	type bucket interface {
		Make(ctx context.Context, in *BucketMakeRequest, out *BlankResponse) error
		List(ctx context.Context, in *BucketListRequest, out *BucketListResponse) error
		Remove(ctx context.Context, in *BucketRemoveRequest, out *BlankResponse) error
		Get(ctx context.Context, in *BucketGetRequest, out *BucketGetResponse) error
		UpdateEngine(ctx context.Context, in *BucketUpdateEngineRequest, out *BlankResponse) error
		UpdateCapacity(ctx context.Context, in *BucketUpdateCapacityRequest, out *BlankResponse) error
		ResetToken(ctx context.Context, in *BucketResetTokenRequest, out *BlankResponse) error
	}
	type Bucket struct {
		bucket
	}
	h := &bucketHandler{hdlr}
	return s.Handle(s.NewHandler(&Bucket{h}, opts...))
}

type bucketHandler struct {
	BucketHandler
}

func (h *bucketHandler) Make(ctx context.Context, in *BucketMakeRequest, out *BlankResponse) error {
	return h.BucketHandler.Make(ctx, in, out)
}

func (h *bucketHandler) List(ctx context.Context, in *BucketListRequest, out *BucketListResponse) error {
	return h.BucketHandler.List(ctx, in, out)
}

func (h *bucketHandler) Remove(ctx context.Context, in *BucketRemoveRequest, out *BlankResponse) error {
	return h.BucketHandler.Remove(ctx, in, out)
}

func (h *bucketHandler) Get(ctx context.Context, in *BucketGetRequest, out *BucketGetResponse) error {
	return h.BucketHandler.Get(ctx, in, out)
}

func (h *bucketHandler) UpdateEngine(ctx context.Context, in *BucketUpdateEngineRequest, out *BlankResponse) error {
	return h.BucketHandler.UpdateEngine(ctx, in, out)
}

func (h *bucketHandler) UpdateCapacity(ctx context.Context, in *BucketUpdateCapacityRequest, out *BlankResponse) error {
	return h.BucketHandler.UpdateCapacity(ctx, in, out)
}

func (h *bucketHandler) ResetToken(ctx context.Context, in *BucketResetTokenRequest, out *BlankResponse) error {
	return h.BucketHandler.ResetToken(ctx, in, out)
}
