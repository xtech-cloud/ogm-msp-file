// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/file/object.proto

package file

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Object service

func NewObjectEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Object service

type ObjectService interface {
	// 准备一个对象的元数据
	Prepare(ctx context.Context, in *ObjectPrepareRequest, opts ...client.CallOption) (*ObjectPrepareResponse, error)
	// 写入一个对象的元数据
	Flush(ctx context.Context, in *ObjectFlushRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 获取一个对象信息
	Get(ctx context.Context, in *ObjectGetRequest, opts ...client.CallOption) (*ObjectGetResponse, error)
	// 精确查找一个对象信息
	Find(ctx context.Context, in *ObjectFindRequest, opts ...client.CallOption) (*ObjectFindResponse, error)
	// 删除一个对象
	Remove(ctx context.Context, in *ObjectRemoveRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 列举一个存储桶中的所有对象
	List(ctx context.Context, in *ObjectListRequest, opts ...client.CallOption) (*ObjectListResponse, error)
	// 模糊查找存储桶中的对象
	Search(ctx context.Context, in *ObjectSearchRequest, opts ...client.CallOption) (*ObjectSearchResponse, error)
	// 发布一个对象
	// 生成指定有效期的公开链接, 对象的URL有值
	Publish(ctx context.Context, in *ObjectPublishRequest, opts ...client.CallOption) (*ObjectPublishResponse, error)
	// 预览一个对象
	// 生成临时的公开链接, 对象的URL无值
	Preview(ctx context.Context, in *ObjectPreviewRequest, opts ...client.CallOption) (*ObjectPreviewResponse, error)
	// 撤回一个对象
	// 撤回公开链接，对象的URL无值
	Retract(ctx context.Context, in *ObjectRetractRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 将Base64编码的内容转换为对象
	ConvertFromBase64(ctx context.Context, in *ObjectConvertFromBase64Request, opts ...client.CallOption) (*ObjectConvertFromBase64Response, error)
	// 将Url地址转换为对象
	ConvertFromUrl(ctx context.Context, in *ObjectConvertFromUrlRequest, opts ...client.CallOption) (*ObjectConvertFromUrlResponse, error)
}

type objectService struct {
	c    client.Client
	name string
}

func NewObjectService(name string, c client.Client) ObjectService {
	return &objectService{
		c:    c,
		name: name,
	}
}

func (c *objectService) Prepare(ctx context.Context, in *ObjectPrepareRequest, opts ...client.CallOption) (*ObjectPrepareResponse, error) {
	req := c.c.NewRequest(c.name, "Object.Prepare", in)
	out := new(ObjectPrepareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) Flush(ctx context.Context, in *ObjectFlushRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Object.Flush", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) Get(ctx context.Context, in *ObjectGetRequest, opts ...client.CallOption) (*ObjectGetResponse, error) {
	req := c.c.NewRequest(c.name, "Object.Get", in)
	out := new(ObjectGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) Find(ctx context.Context, in *ObjectFindRequest, opts ...client.CallOption) (*ObjectFindResponse, error) {
	req := c.c.NewRequest(c.name, "Object.Find", in)
	out := new(ObjectFindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) Remove(ctx context.Context, in *ObjectRemoveRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Object.Remove", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) List(ctx context.Context, in *ObjectListRequest, opts ...client.CallOption) (*ObjectListResponse, error) {
	req := c.c.NewRequest(c.name, "Object.List", in)
	out := new(ObjectListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) Search(ctx context.Context, in *ObjectSearchRequest, opts ...client.CallOption) (*ObjectSearchResponse, error) {
	req := c.c.NewRequest(c.name, "Object.Search", in)
	out := new(ObjectSearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) Publish(ctx context.Context, in *ObjectPublishRequest, opts ...client.CallOption) (*ObjectPublishResponse, error) {
	req := c.c.NewRequest(c.name, "Object.Publish", in)
	out := new(ObjectPublishResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) Preview(ctx context.Context, in *ObjectPreviewRequest, opts ...client.CallOption) (*ObjectPreviewResponse, error) {
	req := c.c.NewRequest(c.name, "Object.Preview", in)
	out := new(ObjectPreviewResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) Retract(ctx context.Context, in *ObjectRetractRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Object.Retract", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) ConvertFromBase64(ctx context.Context, in *ObjectConvertFromBase64Request, opts ...client.CallOption) (*ObjectConvertFromBase64Response, error) {
	req := c.c.NewRequest(c.name, "Object.ConvertFromBase64", in)
	out := new(ObjectConvertFromBase64Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) ConvertFromUrl(ctx context.Context, in *ObjectConvertFromUrlRequest, opts ...client.CallOption) (*ObjectConvertFromUrlResponse, error) {
	req := c.c.NewRequest(c.name, "Object.ConvertFromUrl", in)
	out := new(ObjectConvertFromUrlResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Object service

type ObjectHandler interface {
	// 准备一个对象的元数据
	Prepare(context.Context, *ObjectPrepareRequest, *ObjectPrepareResponse) error
	// 写入一个对象的元数据
	Flush(context.Context, *ObjectFlushRequest, *UuidResponse) error
	// 获取一个对象信息
	Get(context.Context, *ObjectGetRequest, *ObjectGetResponse) error
	// 精确查找一个对象信息
	Find(context.Context, *ObjectFindRequest, *ObjectFindResponse) error
	// 删除一个对象
	Remove(context.Context, *ObjectRemoveRequest, *UuidResponse) error
	// 列举一个存储桶中的所有对象
	List(context.Context, *ObjectListRequest, *ObjectListResponse) error
	// 模糊查找存储桶中的对象
	Search(context.Context, *ObjectSearchRequest, *ObjectSearchResponse) error
	// 发布一个对象
	// 生成指定有效期的公开链接, 对象的URL有值
	Publish(context.Context, *ObjectPublishRequest, *ObjectPublishResponse) error
	// 预览一个对象
	// 生成临时的公开链接, 对象的URL无值
	Preview(context.Context, *ObjectPreviewRequest, *ObjectPreviewResponse) error
	// 撤回一个对象
	// 撤回公开链接，对象的URL无值
	Retract(context.Context, *ObjectRetractRequest, *UuidResponse) error
	// 将Base64编码的内容转换为对象
	ConvertFromBase64(context.Context, *ObjectConvertFromBase64Request, *ObjectConvertFromBase64Response) error
	// 将Url地址转换为对象
	ConvertFromUrl(context.Context, *ObjectConvertFromUrlRequest, *ObjectConvertFromUrlResponse) error
}

func RegisterObjectHandler(s server.Server, hdlr ObjectHandler, opts ...server.HandlerOption) error {
	type object interface {
		Prepare(ctx context.Context, in *ObjectPrepareRequest, out *ObjectPrepareResponse) error
		Flush(ctx context.Context, in *ObjectFlushRequest, out *UuidResponse) error
		Get(ctx context.Context, in *ObjectGetRequest, out *ObjectGetResponse) error
		Find(ctx context.Context, in *ObjectFindRequest, out *ObjectFindResponse) error
		Remove(ctx context.Context, in *ObjectRemoveRequest, out *UuidResponse) error
		List(ctx context.Context, in *ObjectListRequest, out *ObjectListResponse) error
		Search(ctx context.Context, in *ObjectSearchRequest, out *ObjectSearchResponse) error
		Publish(ctx context.Context, in *ObjectPublishRequest, out *ObjectPublishResponse) error
		Preview(ctx context.Context, in *ObjectPreviewRequest, out *ObjectPreviewResponse) error
		Retract(ctx context.Context, in *ObjectRetractRequest, out *UuidResponse) error
		ConvertFromBase64(ctx context.Context, in *ObjectConvertFromBase64Request, out *ObjectConvertFromBase64Response) error
		ConvertFromUrl(ctx context.Context, in *ObjectConvertFromUrlRequest, out *ObjectConvertFromUrlResponse) error
	}
	type Object struct {
		object
	}
	h := &objectHandler{hdlr}
	return s.Handle(s.NewHandler(&Object{h}, opts...))
}

type objectHandler struct {
	ObjectHandler
}

func (h *objectHandler) Prepare(ctx context.Context, in *ObjectPrepareRequest, out *ObjectPrepareResponse) error {
	return h.ObjectHandler.Prepare(ctx, in, out)
}

func (h *objectHandler) Flush(ctx context.Context, in *ObjectFlushRequest, out *UuidResponse) error {
	return h.ObjectHandler.Flush(ctx, in, out)
}

func (h *objectHandler) Get(ctx context.Context, in *ObjectGetRequest, out *ObjectGetResponse) error {
	return h.ObjectHandler.Get(ctx, in, out)
}

func (h *objectHandler) Find(ctx context.Context, in *ObjectFindRequest, out *ObjectFindResponse) error {
	return h.ObjectHandler.Find(ctx, in, out)
}

func (h *objectHandler) Remove(ctx context.Context, in *ObjectRemoveRequest, out *UuidResponse) error {
	return h.ObjectHandler.Remove(ctx, in, out)
}

func (h *objectHandler) List(ctx context.Context, in *ObjectListRequest, out *ObjectListResponse) error {
	return h.ObjectHandler.List(ctx, in, out)
}

func (h *objectHandler) Search(ctx context.Context, in *ObjectSearchRequest, out *ObjectSearchResponse) error {
	return h.ObjectHandler.Search(ctx, in, out)
}

func (h *objectHandler) Publish(ctx context.Context, in *ObjectPublishRequest, out *ObjectPublishResponse) error {
	return h.ObjectHandler.Publish(ctx, in, out)
}

func (h *objectHandler) Preview(ctx context.Context, in *ObjectPreviewRequest, out *ObjectPreviewResponse) error {
	return h.ObjectHandler.Preview(ctx, in, out)
}

func (h *objectHandler) Retract(ctx context.Context, in *ObjectRetractRequest, out *UuidResponse) error {
	return h.ObjectHandler.Retract(ctx, in, out)
}

func (h *objectHandler) ConvertFromBase64(ctx context.Context, in *ObjectConvertFromBase64Request, out *ObjectConvertFromBase64Response) error {
	return h.ObjectHandler.ConvertFromBase64(ctx, in, out)
}

func (h *objectHandler) ConvertFromUrl(ctx context.Context, in *ObjectConvertFromUrlRequest, out *ObjectConvertFromUrlResponse) error {
	return h.ObjectHandler.ConvertFromUrl(ctx, in, out)
}
