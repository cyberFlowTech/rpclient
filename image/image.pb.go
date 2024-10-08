// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: define/image.proto

package image

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

// 压缩场景
type ThumbType int32

const (
	ThumbType_avatar   ThumbType = 0
	ThumbType_dynamics ThumbType = 1
)

// Enum value maps for ThumbType.
var (
	ThumbType_name = map[int32]string{
		0: "avatar",
		1: "dynamics",
	}
	ThumbType_value = map[string]int32{
		"avatar":   0,
		"dynamics": 1,
	}
)

func (x ThumbType) Enum() *ThumbType {
	p := new(ThumbType)
	*p = x
	return p
}

func (x ThumbType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ThumbType) Descriptor() protoreflect.EnumDescriptor {
	return file_define_image_proto_enumTypes[0].Descriptor()
}

func (ThumbType) Type() protoreflect.EnumType {
	return &file_define_image_proto_enumTypes[0]
}

func (x ThumbType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ThumbType.Descriptor instead.
func (ThumbType) EnumDescriptor() ([]byte, []int) {
	return file_define_image_proto_rawDescGZIP(), []int{0}
}

// 图片大小
type ImageSize int32

const (
	ImageSize_small  ImageSize = 0 // 小
	ImageSize_medium ImageSize = 1 // 中
	ImageSize_large  ImageSize = 2 // 大
)

// Enum value maps for ImageSize.
var (
	ImageSize_name = map[int32]string{
		0: "small",
		1: "medium",
		2: "large",
	}
	ImageSize_value = map[string]int32{
		"small":  0,
		"medium": 1,
		"large":  2,
	}
)

func (x ImageSize) Enum() *ImageSize {
	p := new(ImageSize)
	*p = x
	return p
}

func (x ImageSize) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageSize) Descriptor() protoreflect.EnumDescriptor {
	return file_define_image_proto_enumTypes[1].Descriptor()
}

func (ImageSize) Type() protoreflect.EnumType {
	return &file_define_image_proto_enumTypes[1]
}

func (x ImageSize) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageSize.Descriptor instead.
func (ImageSize) EnumDescriptor() ([]byte, []int) {
	return file_define_image_proto_rawDescGZIP(), []int{1}
}

type ImageQuality int32

const (
	ImageQuality_high     ImageQuality = 0 // 高
	ImageQuality_q_medium ImageQuality = 1 // 中
	ImageQuality_low      ImageQuality = 2 // 低
)

// Enum value maps for ImageQuality.
var (
	ImageQuality_name = map[int32]string{
		0: "high",
		1: "q_medium",
		2: "low",
	}
	ImageQuality_value = map[string]int32{
		"high":     0,
		"q_medium": 1,
		"low":      2,
	}
)

func (x ImageQuality) Enum() *ImageQuality {
	p := new(ImageQuality)
	*p = x
	return p
}

func (x ImageQuality) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageQuality) Descriptor() protoreflect.EnumDescriptor {
	return file_define_image_proto_enumTypes[2].Descriptor()
}

func (ImageQuality) Type() protoreflect.EnumType {
	return &file_define_image_proto_enumTypes[2]
}

func (x ImageQuality) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageQuality.Descriptor instead.
func (ImageQuality) EnumDescriptor() ([]byte, []int) {
	return file_define_image_proto_rawDescGZIP(), []int{2}
}

type UploadByUrlReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string       `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
	Name    string       `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"` // 文件名 例：image/20240220/test.jpg
	Type    ThumbType    `protobuf:"varint,3,opt,name=Type,proto3,enum=image.ThumbType" json:"Type,omitempty"`
	Size    ImageSize    `protobuf:"varint,4,opt,name=Size,proto3,enum=image.ImageSize" json:"Size,omitempty"`
	Quality ImageQuality `protobuf:"varint,5,opt,name=Quality,proto3,enum=image.ImageQuality" json:"Quality,omitempty"`
}

func (x *UploadByUrlReq) Reset() {
	*x = UploadByUrlReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadByUrlReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadByUrlReq) ProtoMessage() {}

func (x *UploadByUrlReq) ProtoReflect() protoreflect.Message {
	mi := &file_define_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadByUrlReq.ProtoReflect.Descriptor instead.
func (*UploadByUrlReq) Descriptor() ([]byte, []int) {
	return file_define_image_proto_rawDescGZIP(), []int{0}
}

func (x *UploadByUrlReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UploadByUrlReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadByUrlReq) GetType() ThumbType {
	if x != nil {
		return x.Type
	}
	return ThumbType_avatar
}

func (x *UploadByUrlReq) GetSize() ImageSize {
	if x != nil {
		return x.Size
	}
	return ImageSize_small
}

func (x *UploadByUrlReq) GetQuality() ImageQuality {
	if x != nil {
		return x.Quality
	}
	return ImageQuality_high
}

type UploadByFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File    []byte       `protobuf:"bytes,1,opt,name=File,proto3" json:"File,omitempty"`
	Name    string       `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"` // 文件名
	Type    ThumbType    `protobuf:"varint,3,opt,name=Type,proto3,enum=image.ThumbType" json:"Type,omitempty"`
	Size    ImageSize    `protobuf:"varint,4,opt,name=Size,proto3,enum=image.ImageSize" json:"Size,omitempty"`
	Quality ImageQuality `protobuf:"varint,5,opt,name=Quality,proto3,enum=image.ImageQuality" json:"Quality,omitempty"`
}

func (x *UploadByFileReq) Reset() {
	*x = UploadByFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadByFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadByFileReq) ProtoMessage() {}

func (x *UploadByFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_define_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadByFileReq.ProtoReflect.Descriptor instead.
func (*UploadByFileReq) Descriptor() ([]byte, []int) {
	return file_define_image_proto_rawDescGZIP(), []int{1}
}

func (x *UploadByFileReq) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *UploadByFileReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadByFileReq) GetType() ThumbType {
	if x != nil {
		return x.Type
	}
	return ThumbType_avatar
}

func (x *UploadByFileReq) GetSize() ImageSize {
	if x != nil {
		return x.Size
	}
	return ImageSize_small
}

func (x *UploadByFileReq) GetQuality() ImageQuality {
	if x != nil {
		return x.Quality
	}
	return ImageQuality_high
}

type CommonRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *CommonRes) Reset() {
	*x = CommonRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_image_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonRes) ProtoMessage() {}

func (x *CommonRes) ProtoReflect() protoreflect.Message {
	mi := &file_define_image_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonRes.ProtoReflect.Descriptor instead.
func (*CommonRes) Descriptor() ([]byte, []int) {
	return file_define_image_proto_rawDescGZIP(), []int{2}
}

func (x *CommonRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UploadRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   *CommonRes `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	Url      string     `protobuf:"bytes,2,opt,name=Url,proto3" json:"Url,omitempty"`
	ThumbUrl string     `protobuf:"bytes,3,opt,name=ThumbUrl,proto3" json:"ThumbUrl,omitempty"`
}

func (x *UploadRes) Reset() {
	*x = UploadRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_image_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRes) ProtoMessage() {}

func (x *UploadRes) ProtoReflect() protoreflect.Message {
	mi := &file_define_image_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRes.ProtoReflect.Descriptor instead.
func (*UploadRes) Descriptor() ([]byte, []int) {
	return file_define_image_proto_rawDescGZIP(), []int{3}
}

func (x *UploadRes) GetResult() *CommonRes {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *UploadRes) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UploadRes) GetThumbUrl() string {
	if x != nil {
		return x.ThumbUrl
	}
	return ""
}

type SolidColorAvatarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text   string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
	Width  int64  `protobuf:"varint,2,opt,name=Width,proto3" json:"Width,omitempty"`
	Height int64  `protobuf:"varint,3,opt,name=Height,proto3" json:"Height,omitempty"`
}

func (x *SolidColorAvatarReq) Reset() {
	*x = SolidColorAvatarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_image_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolidColorAvatarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolidColorAvatarReq) ProtoMessage() {}

func (x *SolidColorAvatarReq) ProtoReflect() protoreflect.Message {
	mi := &file_define_image_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolidColorAvatarReq.ProtoReflect.Descriptor instead.
func (*SolidColorAvatarReq) Descriptor() ([]byte, []int) {
	return file_define_image_proto_rawDescGZIP(), []int{4}
}

func (x *SolidColorAvatarReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SolidColorAvatarReq) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *SolidColorAvatarReq) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type SolidColorAvatarRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   *CommonRes `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	Url      string     `protobuf:"bytes,2,opt,name=Url,proto3" json:"Url,omitempty"`
	ThumbUrl string     `protobuf:"bytes,3,opt,name=ThumbUrl,proto3" json:"ThumbUrl,omitempty"`
}

func (x *SolidColorAvatarRes) Reset() {
	*x = SolidColorAvatarRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_image_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolidColorAvatarRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolidColorAvatarRes) ProtoMessage() {}

func (x *SolidColorAvatarRes) ProtoReflect() protoreflect.Message {
	mi := &file_define_image_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolidColorAvatarRes.ProtoReflect.Descriptor instead.
func (*SolidColorAvatarRes) Descriptor() ([]byte, []int) {
	return file_define_image_proto_rawDescGZIP(), []int{5}
}

func (x *SolidColorAvatarRes) GetResult() *CommonRes {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *SolidColorAvatarRes) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SolidColorAvatarRes) GetThumbUrl() string {
	if x != nil {
		return x.ThumbUrl
	}
	return ""
}

var File_define_image_proto protoreflect.FileDescriptor

var file_define_image_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x0e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x79, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x2d, 0x0a, 0x07, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x51,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x07, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22,
	0xb4, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x24, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x51, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x07, 0x51,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x31, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x63, 0x0a, 0x09, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55,
	0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x55, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x55, 0x72, 0x6c, 0x22, 0x57,
	0x0a, 0x13, 0x53, 0x6f, 0x6c, 0x69, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x6d, 0x0a, 0x13, 0x53, 0x6f, 0x6c, 0x69, 0x64,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x12, 0x28,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x55, 0x72, 0x6c, 0x2a, 0x25, 0x0a, 0x09, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x73, 0x10, 0x01, 0x2a, 0x2d, 0x0a,
	0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x73, 0x6d,
	0x61, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x10, 0x02, 0x2a, 0x2f, 0x0a, 0x0c,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x08, 0x0a, 0x04,
	0x68, 0x69, 0x67, 0x68, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x71, 0x5f, 0x6d, 0x65, 0x64, 0x69,
	0x75, 0x6d, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x10, 0x02, 0x32, 0xcc, 0x01,
	0x0a, 0x09, 0x4d, 0x69, 0x6d, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x15, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x79, 0x55, 0x72, 0x6c, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x79, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x12, 0x4d, 0x0a,
	0x13, 0x47, 0x65, 0x6e, 0x53, 0x6f, 0x6c, 0x69, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x6f, 0x6c,
	0x69, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x1a, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x6f, 0x6c, 0x69, 0x64, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_define_image_proto_rawDescOnce sync.Once
	file_define_image_proto_rawDescData = file_define_image_proto_rawDesc
)

func file_define_image_proto_rawDescGZIP() []byte {
	file_define_image_proto_rawDescOnce.Do(func() {
		file_define_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_define_image_proto_rawDescData)
	})
	return file_define_image_proto_rawDescData
}

var file_define_image_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_define_image_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_define_image_proto_goTypes = []interface{}{
	(ThumbType)(0),              // 0: image.ThumbType
	(ImageSize)(0),              // 1: image.ImageSize
	(ImageQuality)(0),           // 2: image.ImageQuality
	(*UploadByUrlReq)(nil),      // 3: image.UploadByUrlReq
	(*UploadByFileReq)(nil),     // 4: image.UploadByFileReq
	(*CommonRes)(nil),           // 5: image.CommonRes
	(*UploadRes)(nil),           // 6: image.UploadRes
	(*SolidColorAvatarReq)(nil), // 7: image.SolidColorAvatarReq
	(*SolidColorAvatarRes)(nil), // 8: image.SolidColorAvatarRes
}
var file_define_image_proto_depIdxs = []int32{
	0,  // 0: image.UploadByUrlReq.Type:type_name -> image.ThumbType
	1,  // 1: image.UploadByUrlReq.Size:type_name -> image.ImageSize
	2,  // 2: image.UploadByUrlReq.Quality:type_name -> image.ImageQuality
	0,  // 3: image.UploadByFileReq.Type:type_name -> image.ThumbType
	1,  // 4: image.UploadByFileReq.Size:type_name -> image.ImageSize
	2,  // 5: image.UploadByFileReq.Quality:type_name -> image.ImageQuality
	5,  // 6: image.UploadRes.Result:type_name -> image.CommonRes
	5,  // 7: image.SolidColorAvatarRes.Result:type_name -> image.CommonRes
	3,  // 8: image.MimoImage.UploadByUrl:input_type -> image.UploadByUrlReq
	4,  // 9: image.MimoImage.UploadByFile:input_type -> image.UploadByFileReq
	7,  // 10: image.MimoImage.GenSolidColorAvatar:input_type -> image.SolidColorAvatarReq
	6,  // 11: image.MimoImage.UploadByUrl:output_type -> image.UploadRes
	6,  // 12: image.MimoImage.UploadByFile:output_type -> image.UploadRes
	8,  // 13: image.MimoImage.GenSolidColorAvatar:output_type -> image.SolidColorAvatarRes
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_define_image_proto_init() }
func file_define_image_proto_init() {
	if File_define_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_define_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadByUrlReq); i {
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
		file_define_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadByFileReq); i {
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
		file_define_image_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonRes); i {
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
		file_define_image_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRes); i {
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
		file_define_image_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolidColorAvatarReq); i {
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
		file_define_image_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolidColorAvatarRes); i {
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
			RawDescriptor: file_define_image_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_define_image_proto_goTypes,
		DependencyIndexes: file_define_image_proto_depIdxs,
		EnumInfos:         file_define_image_proto_enumTypes,
		MessageInfos:      file_define_image_proto_msgTypes,
	}.Build()
	File_define_image_proto = out.File
	file_define_image_proto_rawDesc = nil
	file_define_image_proto_goTypes = nil
	file_define_image_proto_depIdxs = nil
}
