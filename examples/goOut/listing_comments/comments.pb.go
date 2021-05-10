// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: listing_comments/comments.proto

package listing_comments

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetListingCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListingId string `protobuf:"bytes,1,opt,name=listing_id,json=listingId,proto3" json:"listing_id,omitempty"`
}

func (x *GetListingCommentsRequest) Reset() {
	*x = GetListingCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_comments_comments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListingCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListingCommentsRequest) ProtoMessage() {}

func (x *GetListingCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listing_comments_comments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListingCommentsRequest.ProtoReflect.Descriptor instead.
func (*GetListingCommentsRequest) Descriptor() ([]byte, []int) {
	return file_listing_comments_comments_proto_rawDescGZIP(), []int{0}
}

func (x *GetListingCommentsRequest) GetListingId() string {
	if x != nil {
		return x.ListingId
	}
	return ""
}

type ListingComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*ListingComment_Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *ListingComment) Reset() {
	*x = ListingComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_comments_comments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListingComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListingComment) ProtoMessage() {}

func (x *ListingComment) ProtoReflect() protoreflect.Message {
	mi := &file_listing_comments_comments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListingComment.ProtoReflect.Descriptor instead.
func (*ListingComment) Descriptor() ([]byte, []int) {
	return file_listing_comments_comments_proto_rawDescGZIP(), []int{1}
}

func (x *ListingComment) GetComments() []*ListingComment_Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type GetListingCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListingComment *ListingComment `protobuf:"bytes,1,opt,name=listing_comment,json=listingComment,proto3" json:"listing_comment,omitempty"`
}

func (x *GetListingCommentsResponse) Reset() {
	*x = GetListingCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_comments_comments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListingCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListingCommentsResponse) ProtoMessage() {}

func (x *GetListingCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listing_comments_comments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListingCommentsResponse.ProtoReflect.Descriptor instead.
func (*GetListingCommentsResponse) Descriptor() ([]byte, []int) {
	return file_listing_comments_comments_proto_rawDescGZIP(), []int{2}
}

func (x *GetListingCommentsResponse) GetListingComment() *ListingComment {
	if x != nil {
		return x.ListingComment
	}
	return nil
}

type BulkGetListingCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListingId []string `protobuf:"bytes,1,rep,name=listing_id,json=listingId,proto3" json:"listing_id,omitempty"`
}

func (x *BulkGetListingCommentsRequest) Reset() {
	*x = BulkGetListingCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_comments_comments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkGetListingCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkGetListingCommentsRequest) ProtoMessage() {}

func (x *BulkGetListingCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listing_comments_comments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkGetListingCommentsRequest.ProtoReflect.Descriptor instead.
func (*BulkGetListingCommentsRequest) Descriptor() ([]byte, []int) {
	return file_listing_comments_comments_proto_rawDescGZIP(), []int{3}
}

func (x *BulkGetListingCommentsRequest) GetListingId() []string {
	if x != nil {
		return x.ListingId
	}
	return nil
}

type BulkGetListingCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListingComments []*ListingComment `protobuf:"bytes,1,rep,name=listing_comments,json=listingComments,proto3" json:"listing_comments,omitempty"`
}

func (x *BulkGetListingCommentsResponse) Reset() {
	*x = BulkGetListingCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_comments_comments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkGetListingCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkGetListingCommentsResponse) ProtoMessage() {}

func (x *BulkGetListingCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listing_comments_comments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkGetListingCommentsResponse.ProtoReflect.Descriptor instead.
func (*BulkGetListingCommentsResponse) Descriptor() ([]byte, []int) {
	return file_listing_comments_comments_proto_rawDescGZIP(), []int{4}
}

func (x *BulkGetListingCommentsResponse) GetListingComments() []*ListingComment {
	if x != nil {
		return x.ListingComments
	}
	return nil
}

type ListingComment_Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body  string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ListingComment_Comment) Reset() {
	*x = ListingComment_Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_comments_comments_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListingComment_Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListingComment_Comment) ProtoMessage() {}

func (x *ListingComment_Comment) ProtoReflect() protoreflect.Message {
	mi := &file_listing_comments_comments_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListingComment_Comment.ProtoReflect.Descriptor instead.
func (*ListingComment_Comment) Descriptor() ([]byte, []int) {
	return file_listing_comments_comments_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListingComment_Comment) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListingComment_Comment) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_listing_comments_comments_proto protoreflect.FileDescriptor

var file_listing_comments_comments_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x3a, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22,
	0x8b, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x44, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x33, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x67, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0f, 0x6c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x1d, 0x42, 0x75, 0x6c, 0x6b, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x6d, 0x0a, 0x1e, 0x42, 0x75, 0x6c, 0x6b, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x10, 0x6c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x83, 0x02, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x71, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x2b, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7d, 0x0a, 0x16,
	0x42, 0x75, 0x6c, 0x6b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2f, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3f, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73,
	0x65, 0x6c, 0x6c, 0x2f, 0x61, 0x67, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x4f, 0x75, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_listing_comments_comments_proto_rawDescOnce sync.Once
	file_listing_comments_comments_proto_rawDescData = file_listing_comments_comments_proto_rawDesc
)

func file_listing_comments_comments_proto_rawDescGZIP() []byte {
	file_listing_comments_comments_proto_rawDescOnce.Do(func() {
		file_listing_comments_comments_proto_rawDescData = protoimpl.X.CompressGZIP(file_listing_comments_comments_proto_rawDescData)
	})
	return file_listing_comments_comments_proto_rawDescData
}

var file_listing_comments_comments_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_listing_comments_comments_proto_goTypes = []interface{}{
	(*GetListingCommentsRequest)(nil),      // 0: listing_comments.GetListingCommentsRequest
	(*ListingComment)(nil),                 // 1: listing_comments.ListingComment
	(*GetListingCommentsResponse)(nil),     // 2: listing_comments.GetListingCommentsResponse
	(*BulkGetListingCommentsRequest)(nil),  // 3: listing_comments.BulkGetListingCommentsRequest
	(*BulkGetListingCommentsResponse)(nil), // 4: listing_comments.BulkGetListingCommentsResponse
	(*ListingComment_Comment)(nil),         // 5: listing_comments.ListingComment.Comment
}
var file_listing_comments_comments_proto_depIdxs = []int32{
	5, // 0: listing_comments.ListingComment.comments:type_name -> listing_comments.ListingComment.Comment
	1, // 1: listing_comments.GetListingCommentsResponse.listing_comment:type_name -> listing_comments.ListingComment
	1, // 2: listing_comments.BulkGetListingCommentsResponse.listing_comments:type_name -> listing_comments.ListingComment
	0, // 3: listing_comments.ListingComments.GetListingComments:input_type -> listing_comments.GetListingCommentsRequest
	3, // 4: listing_comments.ListingComments.BulkGetListingComments:input_type -> listing_comments.BulkGetListingCommentsRequest
	2, // 5: listing_comments.ListingComments.GetListingComments:output_type -> listing_comments.GetListingCommentsResponse
	4, // 6: listing_comments.ListingComments.BulkGetListingComments:output_type -> listing_comments.BulkGetListingCommentsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_listing_comments_comments_proto_init() }
func file_listing_comments_comments_proto_init() {
	if File_listing_comments_comments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_listing_comments_comments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListingCommentsRequest); i {
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
		file_listing_comments_comments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListingComment); i {
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
		file_listing_comments_comments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListingCommentsResponse); i {
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
		file_listing_comments_comments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkGetListingCommentsRequest); i {
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
		file_listing_comments_comments_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkGetListingCommentsResponse); i {
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
		file_listing_comments_comments_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListingComment_Comment); i {
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
			RawDescriptor: file_listing_comments_comments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_listing_comments_comments_proto_goTypes,
		DependencyIndexes: file_listing_comments_comments_proto_depIdxs,
		MessageInfos:      file_listing_comments_comments_proto_msgTypes,
	}.Build()
	File_listing_comments_comments_proto = out.File
	file_listing_comments_comments_proto_rawDesc = nil
	file_listing_comments_comments_proto_goTypes = nil
	file_listing_comments_comments_proto_depIdxs = nil
}
