// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/file/shared.proto

package file

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 存储引擎
type Engine int32

const (
	Engine_ENGINE_INVALID Engine = 0   // 无效
	Engine_ENGINE_LOCAL   Engine = 1   // 本地
	Engine_ENGINE_MINIO   Engine = 2   // MinIO
	Engine_ENGINE_QINIU   Engine = 3   // 七牛
	Engine_ENGINE_CUSTOM  Engine = 100 // 自定义
)

// Enum value maps for Engine.
var (
	Engine_name = map[int32]string{
		0:   "ENGINE_INVALID",
		1:   "ENGINE_LOCAL",
		2:   "ENGINE_MINIO",
		3:   "ENGINE_QINIU",
		100: "ENGINE_CUSTOM",
	}
	Engine_value = map[string]int32{
		"ENGINE_INVALID": 0,
		"ENGINE_LOCAL":   1,
		"ENGINE_MINIO":   2,
		"ENGINE_QINIU":   3,
		"ENGINE_CUSTOM":  100,
	}
)

func (x Engine) Enum() *Engine {
	p := new(Engine)
	*p = x
	return p
}

func (x Engine) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Engine) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_file_shared_proto_enumTypes[0].Descriptor()
}

func (Engine) Type() protoreflect.EnumType {
	return &file_proto_file_shared_proto_enumTypes[0]
}

func (x Engine) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Engine.Descriptor instead.
func (Engine) EnumDescriptor() ([]byte, []int) {
	return file_proto_file_shared_proto_rawDescGZIP(), []int{0}
}

// 状态
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 状态信息
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proto_file_shared_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// 空白请求
type BlankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BlankRequest) Reset() {
	*x = BlankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlankRequest) ProtoMessage() {}

func (x *BlankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlankRequest.ProtoReflect.Descriptor instead.
func (*BlankRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_shared_proto_rawDescGZIP(), []int{1}
}

// 空白回复
type BlankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
}

func (x *BlankResponse) Reset() {
	*x = BlankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_shared_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlankResponse) ProtoMessage() {}

func (x *BlankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_shared_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlankResponse.ProtoReflect.Descriptor instead.
func (*BlankResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_shared_proto_rawDescGZIP(), []int{2}
}

func (x *BlankResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// Uuid回复
type UuidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Uuid   string  `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`     // uuid
}

func (x *UuidResponse) Reset() {
	*x = UuidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_shared_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UuidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UuidResponse) ProtoMessage() {}

func (x *UuidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_shared_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UuidResponse.ProtoReflect.Descriptor instead.
func (*UuidResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_shared_proto_rawDescGZIP(), []int{3}
}

func (x *UuidResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *UuidResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 存储桶实体
type BucketEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`                  // uuid
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                  // 名称
	Token        string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`                // 访问令牌
	TotalSize    uint64 `protobuf:"varint,4,opt,name=totalSize,proto3" json:"totalSize,omitempty"`       // 空间总容量, 单位byte
	UsedSize     uint64 `protobuf:"varint,5,opt,name=usedSize,proto3" json:"usedSize,omitempty"`         // 空间已用容量, 单位byte
	Engine       Engine `protobuf:"varint,6,opt,name=engine,proto3,enum=Engine" json:"engine,omitempty"` // 存储引擎
	Address      string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`            // 存储引擎的地址
	Scope        string `protobuf:"bytes,8,opt,name=scope,proto3" json:"scope,omitempty"`                // 作用范围
	AccessKey    string `protobuf:"bytes,9,opt,name=accessKey,proto3" json:"accessKey,omitempty"`        // 存储引擎的访问Key
	AccessSecret string `protobuf:"bytes,10,opt,name=accessSecret,proto3" json:"accessSecret,omitempty"` // 存储引擎的访问Secret
	Url          string `protobuf:"bytes,21,opt,name=url,proto3" json:"url,omitempty"`                   // 存储引擎的外部访问地址
}

func (x *BucketEntity) Reset() {
	*x = BucketEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_shared_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketEntity) ProtoMessage() {}

func (x *BucketEntity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_shared_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketEntity.ProtoReflect.Descriptor instead.
func (*BucketEntity) Descriptor() ([]byte, []int) {
	return file_proto_file_shared_proto_rawDescGZIP(), []int{4}
}

func (x *BucketEntity) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *BucketEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BucketEntity) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *BucketEntity) GetTotalSize() uint64 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *BucketEntity) GetUsedSize() uint64 {
	if x != nil {
		return x.UsedSize
	}
	return 0
}

func (x *BucketEntity) GetEngine() Engine {
	if x != nil {
		return x.Engine
	}
	return Engine_ENGINE_INVALID
}

func (x *BucketEntity) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BucketEntity) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *BucketEntity) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *BucketEntity) GetAccessSecret() string {
	if x != nil {
		return x.AccessSecret
	}
	return ""
}

func (x *BucketEntity) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// 对象实体
type ObjectEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`         // 唯一ID
	Filepath string `protobuf:"bytes,2,opt,name=filepath,proto3" json:"filepath,omitempty"` // 文件路径
	Url      string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`           // 地址
	Size     uint64 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`        // 文件大小
	Md5      string `protobuf:"bytes,5,opt,name=md5,proto3" json:"md5,omitempty"`           // MD5值
	Uname    string `protobuf:"bytes,6,opt,name=uname,proto3" json:"uname,omitempty"`       // 存储名
}

func (x *ObjectEntity) Reset() {
	*x = ObjectEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_shared_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectEntity) ProtoMessage() {}

func (x *ObjectEntity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_shared_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectEntity.ProtoReflect.Descriptor instead.
func (*ObjectEntity) Descriptor() ([]byte, []int) {
	return file_proto_file_shared_proto_rawDescGZIP(), []int{5}
}

func (x *ObjectEntity) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ObjectEntity) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *ObjectEntity) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ObjectEntity) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ObjectEntity) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *ObjectEntity) GetUname() string {
	if x != nil {
		return x.Uname
	}
	return ""
}

// Base64格式的源
type Base64Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filepath string `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"` // 文件路径
	Uname    string `protobuf:"bytes,2,opt,name=uname,proto3" json:"uname,omitempty"`       // 存储名
	Content  string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`   // base64编码的内容
}

func (x *Base64Source) Reset() {
	*x = Base64Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_shared_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base64Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base64Source) ProtoMessage() {}

func (x *Base64Source) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_shared_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base64Source.ProtoReflect.Descriptor instead.
func (*Base64Source) Descriptor() ([]byte, []int) {
	return file_proto_file_shared_proto_rawDescGZIP(), []int{6}
}

func (x *Base64Source) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *Base64Source) GetUname() string {
	if x != nil {
		return x.Uname
	}
	return ""
}

func (x *Base64Source) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_proto_file_shared_proto protoreflect.FileDescriptor

var file_proto_file_shared_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x30, 0x0a, 0x0d, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x43, 0x0a, 0x0c, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0xab, 0x02, 0x0a, 0x0c, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x64, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x75, 0x73, 0x65, 0x64, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1f, 0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x07, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x64, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x12,
	0x14, 0x0a, 0x05, 0x75, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x75, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x75, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2a, 0x65, 0x0a, 0x06, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x45,
	0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x10, 0x0a, 0x0c, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x49,
	0x4f, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x51, 0x49,
	0x4e, 0x49, 0x55, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f,
	0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x64, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x3b, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_file_shared_proto_rawDescOnce sync.Once
	file_proto_file_shared_proto_rawDescData = file_proto_file_shared_proto_rawDesc
)

func file_proto_file_shared_proto_rawDescGZIP() []byte {
	file_proto_file_shared_proto_rawDescOnce.Do(func() {
		file_proto_file_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_file_shared_proto_rawDescData)
	})
	return file_proto_file_shared_proto_rawDescData
}

var file_proto_file_shared_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_file_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_file_shared_proto_goTypes = []interface{}{
	(Engine)(0),           // 0: Engine
	(*Status)(nil),        // 1: Status
	(*BlankRequest)(nil),  // 2: BlankRequest
	(*BlankResponse)(nil), // 3: BlankResponse
	(*UuidResponse)(nil),  // 4: UuidResponse
	(*BucketEntity)(nil),  // 5: BucketEntity
	(*ObjectEntity)(nil),  // 6: ObjectEntity
	(*Base64Source)(nil),  // 7: Base64Source
}
var file_proto_file_shared_proto_depIdxs = []int32{
	1, // 0: BlankResponse.status:type_name -> Status
	1, // 1: UuidResponse.status:type_name -> Status
	0, // 2: BucketEntity.engine:type_name -> Engine
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_file_shared_proto_init() }
func file_proto_file_shared_proto_init() {
	if File_proto_file_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_file_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_file_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlankRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_file_shared_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlankResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_file_shared_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UuidResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_file_shared_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketEntity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_file_shared_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectEntity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_file_shared_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base64Source); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_file_shared_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_file_shared_proto_goTypes,
		DependencyIndexes: file_proto_file_shared_proto_depIdxs,
		EnumInfos:         file_proto_file_shared_proto_enumTypes,
		MessageInfos:      file_proto_file_shared_proto_msgTypes,
	}.Build()
	File_proto_file_shared_proto = out.File
	file_proto_file_shared_proto_rawDesc = nil
	file_proto_file_shared_proto_goTypes = nil
	file_proto_file_shared_proto_depIdxs = nil
}
