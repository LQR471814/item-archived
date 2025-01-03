// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: v1/api.proto

package v1

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

type ImageFormat int32

const (
	ImageFormat_JPG ImageFormat = 0
	ImageFormat_PNG ImageFormat = 1
	ImageFormat_GIF ImageFormat = 2
	ImageFormat_SVG ImageFormat = 3
)

// Enum value maps for ImageFormat.
var (
	ImageFormat_name = map[int32]string{
		0: "JPG",
		1: "PNG",
		2: "GIF",
		3: "SVG",
	}
	ImageFormat_value = map[string]int32{
		"JPG": 0,
		"PNG": 1,
		"GIF": 2,
		"SVG": 3,
	}
)

func (x ImageFormat) Enum() *ImageFormat {
	p := new(ImageFormat)
	*p = x
	return p
}

func (x ImageFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_api_proto_enumTypes[0].Descriptor()
}

func (ImageFormat) Type() protoreflect.EnumType {
	return &file_v1_api_proto_enumTypes[0]
}

func (x ImageFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageFormat.Descriptor instead.
func (ImageFormat) EnumDescriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{0}
}

// EntryMetadata describes the metadata present in both items and containers
type EntryMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tags        []string     `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	Description *string      `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Image       []byte       `protobuf:"bytes,4,opt,name=image,proto3,oneof" json:"image,omitempty"`
	ImageFormat *ImageFormat `protobuf:"varint,5,opt,name=image_format,json=imageFormat,proto3,enum=v1.ImageFormat,oneof" json:"image_format,omitempty"`
}

func (x *EntryMetadata) Reset() {
	*x = EntryMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryMetadata) ProtoMessage() {}

func (x *EntryMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryMetadata.ProtoReflect.Descriptor instead.
func (*EntryMetadata) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *EntryMetadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EntryMetadata) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *EntryMetadata) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *EntryMetadata) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *EntryMetadata) GetImageFormat() ImageFormat {
	if x != nil && x.ImageFormat != nil {
		return *x.ImageFormat
	}
	return ImageFormat_JPG
}

// Read
type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// path should be an array of strings in the following format:
	//
	// [id.tag_1.tag_2.container, id_2.tag_3.container, id_3.item]
	//
	// both the tags and the entry type should be added to the path segment separated by dots
	Path []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *ReadRequest) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *EntryMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// children will be defined if the entry is a container, otherwise it will be null
	Children *ReadResponse_Children `protobuf:"bytes,2,opt,name=children,proto3,oneof" json:"children,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *ReadResponse) GetMetadata() *EntryMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ReadResponse) GetChildren() *ReadResponse_Children {
	if x != nil {
		return x.Children
	}
	return nil
}

// Create creates a container or item
type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *EntryMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// this should follow the same convention as the path in ReadRequest
	Path []string `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	// create_container will create a container instead of an item if true
	CreateContainer bool `protobuf:"varint,3,opt,name=create_container,json=createContainer,proto3" json:"create_container,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRequest) GetMetadata() *EntryMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CreateRequest) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *CreateRequest) GetCreateContainer() bool {
	if x != nil {
		return x.CreateContainer
	}
	return false
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{4}
}

// Move can move a container or an item
type MoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// this should follow the same convention as the path in ReadRequest
	Src []string `protobuf:"bytes,1,rep,name=src,proto3" json:"src,omitempty"`
	// this should follow the same convention as the path in ReadRequest
	Dest []string `protobuf:"bytes,2,rep,name=dest,proto3" json:"dest,omitempty"`
}

func (x *MoveRequest) Reset() {
	*x = MoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveRequest) ProtoMessage() {}

func (x *MoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveRequest.ProtoReflect.Descriptor instead.
func (*MoveRequest) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{5}
}

func (x *MoveRequest) GetSrc() []string {
	if x != nil {
		return x.Src
	}
	return nil
}

func (x *MoveRequest) GetDest() []string {
	if x != nil {
		return x.Dest
	}
	return nil
}

type MoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MoveResponse) Reset() {
	*x = MoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveResponse) ProtoMessage() {}

func (x *MoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveResponse.ProtoReflect.Descriptor instead.
func (*MoveResponse) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{6}
}

// Delete can delete a container or an item
type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// this should follow the same convention as the path in ReadRequest
	Path []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRequest) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{8}
}

// Search
type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{9}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*SearchResponse_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{10}
}

func (x *SearchResponse) GetEntries() []*SearchResponse_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type ReadResponse_Children struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// note: this is a filename formatted as "id.tag_1.tag_2", the .item ext should be appended as necessary
	ItemNames []string `protobuf:"bytes,1,rep,name=item_names,json=itemNames,proto3" json:"item_names,omitempty"`
	// note: this is a filename formatted as "id.tag_1.tag_2", the .container ext should be appended as necessary
	ContainerNames []string `protobuf:"bytes,2,rep,name=container_names,json=containerNames,proto3" json:"container_names,omitempty"`
}

func (x *ReadResponse_Children) Reset() {
	*x = ReadResponse_Children{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse_Children) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse_Children) ProtoMessage() {}

func (x *ReadResponse_Children) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse_Children.ProtoReflect.Descriptor instead.
func (*ReadResponse_Children) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ReadResponse_Children) GetItemNames() []string {
	if x != nil {
		return x.ItemNames
	}
	return nil
}

func (x *ReadResponse_Children) GetContainerNames() []string {
	if x != nil {
		return x.ContainerNames
	}
	return nil
}

type SearchResponse_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path []string       `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	Meta *EntryMetadata `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *SearchResponse_Entry) Reset() {
	*x = SearchResponse_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse_Entry) ProtoMessage() {}

func (x *SearchResponse_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse_Entry.ProtoReflect.Descriptor instead.
func (*SearchResponse_Entry) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{10, 0}
}

func (x *SearchResponse_Entry) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *SearchResponse_Entry) GetMeta() *EntryMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_v1_api_proto protoreflect.FileDescriptor

var file_v1_api_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x76, 0x31, 0x22, 0xd9, 0x01, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x0c, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x48, 0x02, 0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x0f, 0x0a,
	0x0d, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x21,
	0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x22, 0xda, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x3a, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x48, 0x00,
	0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x1a, 0x52, 0x0a,
	0x08, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x7d,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0x10, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x33, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x42, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x2a, 0x31, 0x0a,
	0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x07, 0x0a, 0x03,
	0x4a, 0x50, 0x47, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x47, 0x49, 0x46, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x56, 0x47, 0x10, 0x03,
	0x32, 0xf9, 0x01, 0x0a, 0x0e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0f, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x04, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x50, 0x0a, 0x06,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x14, 0x69, 0x74, 0x65, 0x6d, 0x2d, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02,
	0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_api_proto_rawDescOnce sync.Once
	file_v1_api_proto_rawDescData = file_v1_api_proto_rawDesc
)

func file_v1_api_proto_rawDescGZIP() []byte {
	file_v1_api_proto_rawDescOnce.Do(func() {
		file_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_api_proto_rawDescData)
	})
	return file_v1_api_proto_rawDescData
}

var file_v1_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_v1_api_proto_goTypes = []any{
	(ImageFormat)(0),              // 0: v1.ImageFormat
	(*EntryMetadata)(nil),         // 1: v1.EntryMetadata
	(*ReadRequest)(nil),           // 2: v1.ReadRequest
	(*ReadResponse)(nil),          // 3: v1.ReadResponse
	(*CreateRequest)(nil),         // 4: v1.CreateRequest
	(*CreateResponse)(nil),        // 5: v1.CreateResponse
	(*MoveRequest)(nil),           // 6: v1.MoveRequest
	(*MoveResponse)(nil),          // 7: v1.MoveResponse
	(*DeleteRequest)(nil),         // 8: v1.DeleteRequest
	(*DeleteResponse)(nil),        // 9: v1.DeleteResponse
	(*SearchRequest)(nil),         // 10: v1.SearchRequest
	(*SearchResponse)(nil),        // 11: v1.SearchResponse
	(*ReadResponse_Children)(nil), // 12: v1.ReadResponse.Children
	(*SearchResponse_Entry)(nil),  // 13: v1.SearchResponse.Entry
}
var file_v1_api_proto_depIdxs = []int32{
	0,  // 0: v1.EntryMetadata.image_format:type_name -> v1.ImageFormat
	1,  // 1: v1.ReadResponse.metadata:type_name -> v1.EntryMetadata
	12, // 2: v1.ReadResponse.children:type_name -> v1.ReadResponse.Children
	1,  // 3: v1.CreateRequest.metadata:type_name -> v1.EntryMetadata
	13, // 4: v1.SearchResponse.entries:type_name -> v1.SearchResponse.Entry
	1,  // 5: v1.SearchResponse.Entry.meta:type_name -> v1.EntryMetadata
	2,  // 6: v1.ArchiveService.Read:input_type -> v1.ReadRequest
	4,  // 7: v1.ArchiveService.Create:input_type -> v1.CreateRequest
	6,  // 8: v1.ArchiveService.Move:input_type -> v1.MoveRequest
	8,  // 9: v1.ArchiveService.Delete:input_type -> v1.DeleteRequest
	10, // 10: v1.ArchiveService.Search:input_type -> v1.SearchRequest
	3,  // 11: v1.ArchiveService.Read:output_type -> v1.ReadResponse
	5,  // 12: v1.ArchiveService.Create:output_type -> v1.CreateResponse
	7,  // 13: v1.ArchiveService.Move:output_type -> v1.MoveResponse
	9,  // 14: v1.ArchiveService.Delete:output_type -> v1.DeleteResponse
	11, // 15: v1.ArchiveService.Search:output_type -> v1.SearchResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_v1_api_proto_init() }
func file_v1_api_proto_init() {
	if File_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_api_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EntryMetadata); i {
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
		file_v1_api_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ReadRequest); i {
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
		file_v1_api_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ReadResponse); i {
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
		file_v1_api_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRequest); i {
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
		file_v1_api_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreateResponse); i {
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
		file_v1_api_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*MoveRequest); i {
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
		file_v1_api_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*MoveResponse); i {
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
		file_v1_api_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRequest); i {
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
		file_v1_api_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteResponse); i {
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
		file_v1_api_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*SearchRequest); i {
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
		file_v1_api_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*SearchResponse); i {
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
		file_v1_api_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ReadResponse_Children); i {
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
		file_v1_api_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*SearchResponse_Entry); i {
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
	file_v1_api_proto_msgTypes[0].OneofWrappers = []any{}
	file_v1_api_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_api_proto_goTypes,
		DependencyIndexes: file_v1_api_proto_depIdxs,
		EnumInfos:         file_v1_api_proto_enumTypes,
		MessageInfos:      file_v1_api_proto_msgTypes,
	}.Build()
	File_v1_api_proto = out.File
	file_v1_api_proto_rawDesc = nil
	file_v1_api_proto_goTypes = nil
	file_v1_api_proto_depIdxs = nil
}
