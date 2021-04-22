// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: complex_v1/complex_usecase.proto

package complex_v1

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

type ComplexUsecaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetListings *ComplexUsecaseRequest_GetListingsGen `protobuf:"bytes,1,opt,name=get_listings,json=getListings,proto3" json:"get_listings,omitempty"`
	UserInfo    *ComplexUsecaseRequest_UserInfoGen    `protobuf:"bytes,2,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	Meta        *ComplexUsecaseRequest_MetaGen        `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ComplexUsecaseRequest) Reset() {
	*x = ComplexUsecaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_v1_complex_usecase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexUsecaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexUsecaseRequest) ProtoMessage() {}

func (x *ComplexUsecaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_complex_v1_complex_usecase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexUsecaseRequest.ProtoReflect.Descriptor instead.
func (*ComplexUsecaseRequest) Descriptor() ([]byte, []int) {
	return file_complex_v1_complex_usecase_proto_rawDescGZIP(), []int{0}
}

func (x *ComplexUsecaseRequest) GetGetListings() *ComplexUsecaseRequest_GetListingsGen {
	if x != nil {
		return x.GetListings
	}
	return nil
}

func (x *ComplexUsecaseRequest) GetUserInfo() *ComplexUsecaseRequest_UserInfoGen {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *ComplexUsecaseRequest) GetMeta() *ComplexUsecaseRequest_MetaGen {
	if x != nil {
		return x.Meta
	}
	return nil
}

type ComplexUsecaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listings   []*ComplexUsecaseResponse_ListingsGen `protobuf:"bytes,1,rep,name=listings,proto3" json:"listings,omitempty"`
	UserWallet *ComplexUsecaseResponse_UserWalletGen `protobuf:"bytes,2,opt,name=user_wallet,json=userWallet,proto3" json:"user_wallet,omitempty"`
	ApiVersion int64                                 `protobuf:"varint,3,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	CommitId   string                                `protobuf:"bytes,4,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
}

func (x *ComplexUsecaseResponse) Reset() {
	*x = ComplexUsecaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_v1_complex_usecase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexUsecaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexUsecaseResponse) ProtoMessage() {}

func (x *ComplexUsecaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_complex_v1_complex_usecase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexUsecaseResponse.ProtoReflect.Descriptor instead.
func (*ComplexUsecaseResponse) Descriptor() ([]byte, []int) {
	return file_complex_v1_complex_usecase_proto_rawDescGZIP(), []int{1}
}

func (x *ComplexUsecaseResponse) GetListings() []*ComplexUsecaseResponse_ListingsGen {
	if x != nil {
		return x.Listings
	}
	return nil
}

func (x *ComplexUsecaseResponse) GetUserWallet() *ComplexUsecaseResponse_UserWalletGen {
	if x != nil {
		return x.UserWallet
	}
	return nil
}

func (x *ComplexUsecaseResponse) GetApiVersion() int64 {
	if x != nil {
		return x.ApiVersion
	}
	return 0
}

func (x *ComplexUsecaseResponse) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

type ComplexUsecaseRequest_GetListingsGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListingIds []*ComplexUsecaseRequest_GetListingsGen_ListingIdsGen `protobuf:"bytes,1,rep,name=listing_ids,json=listingIds,proto3" json:"listing_ids,omitempty"`
}

func (x *ComplexUsecaseRequest_GetListingsGen) Reset() {
	*x = ComplexUsecaseRequest_GetListingsGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_v1_complex_usecase_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexUsecaseRequest_GetListingsGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexUsecaseRequest_GetListingsGen) ProtoMessage() {}

func (x *ComplexUsecaseRequest_GetListingsGen) ProtoReflect() protoreflect.Message {
	mi := &file_complex_v1_complex_usecase_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexUsecaseRequest_GetListingsGen.ProtoReflect.Descriptor instead.
func (*ComplexUsecaseRequest_GetListingsGen) Descriptor() ([]byte, []int) {
	return file_complex_v1_complex_usecase_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ComplexUsecaseRequest_GetListingsGen) GetListingIds() []*ComplexUsecaseRequest_GetListingsGen_ListingIdsGen {
	if x != nil {
		return x.ListingIds
	}
	return nil
}

type ComplexUsecaseRequest_UserInfoGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ComplexUsecaseRequest_UserInfoGen) Reset() {
	*x = ComplexUsecaseRequest_UserInfoGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_v1_complex_usecase_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexUsecaseRequest_UserInfoGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexUsecaseRequest_UserInfoGen) ProtoMessage() {}

func (x *ComplexUsecaseRequest_UserInfoGen) ProtoReflect() protoreflect.Message {
	mi := &file_complex_v1_complex_usecase_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexUsecaseRequest_UserInfoGen.ProtoReflect.Descriptor instead.
func (*ComplexUsecaseRequest_UserInfoGen) Descriptor() ([]byte, []int) {
	return file_complex_v1_complex_usecase_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ComplexUsecaseRequest_UserInfoGen) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ComplexUsecaseRequest_MetaGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaType []string `protobuf:"bytes,1,rep,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
}

func (x *ComplexUsecaseRequest_MetaGen) Reset() {
	*x = ComplexUsecaseRequest_MetaGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_v1_complex_usecase_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexUsecaseRequest_MetaGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexUsecaseRequest_MetaGen) ProtoMessage() {}

func (x *ComplexUsecaseRequest_MetaGen) ProtoReflect() protoreflect.Message {
	mi := &file_complex_v1_complex_usecase_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexUsecaseRequest_MetaGen.ProtoReflect.Descriptor instead.
func (*ComplexUsecaseRequest_MetaGen) Descriptor() ([]byte, []int) {
	return file_complex_v1_complex_usecase_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ComplexUsecaseRequest_MetaGen) GetMediaType() []string {
	if x != nil {
		return x.MediaType
	}
	return nil
}

type ComplexUsecaseRequest_GetListingsGen_ListingIdsGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids string `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ComplexUsecaseRequest_GetListingsGen_ListingIdsGen) Reset() {
	*x = ComplexUsecaseRequest_GetListingsGen_ListingIdsGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_v1_complex_usecase_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexUsecaseRequest_GetListingsGen_ListingIdsGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexUsecaseRequest_GetListingsGen_ListingIdsGen) ProtoMessage() {}

func (x *ComplexUsecaseRequest_GetListingsGen_ListingIdsGen) ProtoReflect() protoreflect.Message {
	mi := &file_complex_v1_complex_usecase_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexUsecaseRequest_GetListingsGen_ListingIdsGen.ProtoReflect.Descriptor instead.
func (*ComplexUsecaseRequest_GetListingsGen_ListingIdsGen) Descriptor() ([]byte, []int) {
	return file_complex_v1_complex_usecase_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ComplexUsecaseRequest_GetListingsGen_ListingIdsGen) GetIds() string {
	if x != nil {
		return x.Ids
	}
	return ""
}

type ComplexUsecaseResponse_ListingsGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info  *ComplexUsecaseResponse_ListingsGen_InfoGen  `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Media *ComplexUsecaseResponse_ListingsGen_MediaGen `protobuf:"bytes,2,opt,name=media,proto3" json:"media,omitempty"`
}

func (x *ComplexUsecaseResponse_ListingsGen) Reset() {
	*x = ComplexUsecaseResponse_ListingsGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_v1_complex_usecase_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexUsecaseResponse_ListingsGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexUsecaseResponse_ListingsGen) ProtoMessage() {}

func (x *ComplexUsecaseResponse_ListingsGen) ProtoReflect() protoreflect.Message {
	mi := &file_complex_v1_complex_usecase_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexUsecaseResponse_ListingsGen.ProtoReflect.Descriptor instead.
func (*ComplexUsecaseResponse_ListingsGen) Descriptor() ([]byte, []int) {
	return file_complex_v1_complex_usecase_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ComplexUsecaseResponse_ListingsGen) GetInfo() *ComplexUsecaseResponse_ListingsGen_InfoGen {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ComplexUsecaseResponse_ListingsGen) GetMedia() *ComplexUsecaseResponse_ListingsGen_MediaGen {
	if x != nil {
		return x.Media
	}
	return nil
}

type ComplexUsecaseResponse_UserWalletGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance int64 `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *ComplexUsecaseResponse_UserWalletGen) Reset() {
	*x = ComplexUsecaseResponse_UserWalletGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_v1_complex_usecase_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexUsecaseResponse_UserWalletGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexUsecaseResponse_UserWalletGen) ProtoMessage() {}

func (x *ComplexUsecaseResponse_UserWalletGen) ProtoReflect() protoreflect.Message {
	mi := &file_complex_v1_complex_usecase_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexUsecaseResponse_UserWalletGen.ProtoReflect.Descriptor instead.
func (*ComplexUsecaseResponse_UserWalletGen) Descriptor() ([]byte, []int) {
	return file_complex_v1_complex_usecase_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ComplexUsecaseResponse_UserWalletGen) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type ComplexUsecaseResponse_ListingsGen_InfoGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ComplexUsecaseResponse_ListingsGen_InfoGen) Reset() {
	*x = ComplexUsecaseResponse_ListingsGen_InfoGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_v1_complex_usecase_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexUsecaseResponse_ListingsGen_InfoGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexUsecaseResponse_ListingsGen_InfoGen) ProtoMessage() {}

func (x *ComplexUsecaseResponse_ListingsGen_InfoGen) ProtoReflect() protoreflect.Message {
	mi := &file_complex_v1_complex_usecase_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexUsecaseResponse_ListingsGen_InfoGen.ProtoReflect.Descriptor instead.
func (*ComplexUsecaseResponse_ListingsGen_InfoGen) Descriptor() ([]byte, []int) {
	return file_complex_v1_complex_usecase_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *ComplexUsecaseResponse_ListingsGen_InfoGen) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ComplexUsecaseResponse_ListingsGen_InfoGen) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ComplexUsecaseResponse_ListingsGen_MediaGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhotoUrl string `protobuf:"bytes,1,opt,name=photo_url,json=photoUrl,proto3" json:"photo_url,omitempty"`
}

func (x *ComplexUsecaseResponse_ListingsGen_MediaGen) Reset() {
	*x = ComplexUsecaseResponse_ListingsGen_MediaGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_v1_complex_usecase_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexUsecaseResponse_ListingsGen_MediaGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexUsecaseResponse_ListingsGen_MediaGen) ProtoMessage() {}

func (x *ComplexUsecaseResponse_ListingsGen_MediaGen) ProtoReflect() protoreflect.Message {
	mi := &file_complex_v1_complex_usecase_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexUsecaseResponse_ListingsGen_MediaGen.ProtoReflect.Descriptor instead.
func (*ComplexUsecaseResponse_ListingsGen_MediaGen) Descriptor() ([]byte, []int) {
	return file_complex_v1_complex_usecase_proto_rawDescGZIP(), []int{1, 0, 1}
}

func (x *ComplexUsecaseResponse_ListingsGen_MediaGen) GetPhotoUrl() string {
	if x != nil {
		return x.PhotoUrl
	}
	return ""
}

var File_complex_v1_complex_usecase_proto protoreflect.FileDescriptor

var file_complex_v1_complex_usecase_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x76, 0x31, 0x22, 0xd7,
	0x03, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x55, 0x73, 0x65, 0x63, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x78, 0x55, 0x73, 0x65, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x47, 0x65, 0x6e,
	0x52, 0x0b, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4a, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x55, 0x73, 0x65, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x47, 0x65, 0x6e, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x78, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x55, 0x73, 0x65, 0x63,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x47,
	0x65, 0x6e, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x94, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x47, 0x65, 0x6e, 0x12, 0x5f, 0x0a, 0x0b, 0x6c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x55, 0x73, 0x65, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x47,
	0x65, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x73, 0x47, 0x65, 0x6e,
	0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x73, 0x1a, 0x21, 0x0a, 0x0d,
	0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x73, 0x47, 0x65, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x1a,
	0x1d, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x47, 0x65, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x28,
	0x0a, 0x07, 0x4d, 0x65, 0x74, 0x61, 0x47, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x22, 0xb7, 0x04, 0x0a, 0x16, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x55, 0x73, 0x65, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x55, 0x73, 0x65, 0x63, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x47, 0x65, 0x6e, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x51, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x55, 0x73, 0x65, 0x63, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64,
	0x1a, 0x94, 0x02, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x47, 0x65, 0x6e,
	0x12, 0x4a, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x78, 0x55, 0x73, 0x65, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x47, 0x65, 0x6e, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x47, 0x65, 0x6e, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x4d, 0x0a, 0x05,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x55, 0x73, 0x65, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x47, 0x65, 0x6e, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x47, 0x65, 0x6e, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x1a, 0x41, 0x0a, 0x07, 0x49,
	0x6e, 0x66, 0x6f, 0x47, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x27,
	0x0a, 0x08, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x47, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x1a, 0x29, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x32, 0x76, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x55, 0x73, 0x65,
	0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x14, 0x49,
	0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x55, 0x73, 0x65, 0x63,
	0x61, 0x73, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x55, 0x73, 0x65, 0x63, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x5f, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x55, 0x73, 0x65, 0x63, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65,
	0x6c, 0x6c, 0x2f, 0x61, 0x67, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x4f, 0x75, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_complex_v1_complex_usecase_proto_rawDescOnce sync.Once
	file_complex_v1_complex_usecase_proto_rawDescData = file_complex_v1_complex_usecase_proto_rawDesc
)

func file_complex_v1_complex_usecase_proto_rawDescGZIP() []byte {
	file_complex_v1_complex_usecase_proto_rawDescOnce.Do(func() {
		file_complex_v1_complex_usecase_proto_rawDescData = protoimpl.X.CompressGZIP(file_complex_v1_complex_usecase_proto_rawDescData)
	})
	return file_complex_v1_complex_usecase_proto_rawDescData
}

var file_complex_v1_complex_usecase_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_complex_v1_complex_usecase_proto_goTypes = []interface{}{
	(*ComplexUsecaseRequest)(nil),                              // 0: complex_v1.ComplexUsecaseRequest
	(*ComplexUsecaseResponse)(nil),                             // 1: complex_v1.ComplexUsecaseResponse
	(*ComplexUsecaseRequest_GetListingsGen)(nil),               // 2: complex_v1.ComplexUsecaseRequest.GetListingsGen
	(*ComplexUsecaseRequest_UserInfoGen)(nil),                  // 3: complex_v1.ComplexUsecaseRequest.UserInfoGen
	(*ComplexUsecaseRequest_MetaGen)(nil),                      // 4: complex_v1.ComplexUsecaseRequest.MetaGen
	(*ComplexUsecaseRequest_GetListingsGen_ListingIdsGen)(nil), // 5: complex_v1.ComplexUsecaseRequest.GetListingsGen.ListingIdsGen
	(*ComplexUsecaseResponse_ListingsGen)(nil),                 // 6: complex_v1.ComplexUsecaseResponse.ListingsGen
	(*ComplexUsecaseResponse_UserWalletGen)(nil),               // 7: complex_v1.ComplexUsecaseResponse.UserWalletGen
	(*ComplexUsecaseResponse_ListingsGen_InfoGen)(nil),         // 8: complex_v1.ComplexUsecaseResponse.ListingsGen.InfoGen
	(*ComplexUsecaseResponse_ListingsGen_MediaGen)(nil),        // 9: complex_v1.ComplexUsecaseResponse.ListingsGen.MediaGen
}
var file_complex_v1_complex_usecase_proto_depIdxs = []int32{
	2, // 0: complex_v1.ComplexUsecaseRequest.get_listings:type_name -> complex_v1.ComplexUsecaseRequest.GetListingsGen
	3, // 1: complex_v1.ComplexUsecaseRequest.user_info:type_name -> complex_v1.ComplexUsecaseRequest.UserInfoGen
	4, // 2: complex_v1.ComplexUsecaseRequest.meta:type_name -> complex_v1.ComplexUsecaseRequest.MetaGen
	6, // 3: complex_v1.ComplexUsecaseResponse.listings:type_name -> complex_v1.ComplexUsecaseResponse.ListingsGen
	7, // 4: complex_v1.ComplexUsecaseResponse.user_wallet:type_name -> complex_v1.ComplexUsecaseResponse.UserWalletGen
	5, // 5: complex_v1.ComplexUsecaseRequest.GetListingsGen.listing_ids:type_name -> complex_v1.ComplexUsecaseRequest.GetListingsGen.ListingIdsGen
	8, // 6: complex_v1.ComplexUsecaseResponse.ListingsGen.info:type_name -> complex_v1.ComplexUsecaseResponse.ListingsGen.InfoGen
	9, // 7: complex_v1.ComplexUsecaseResponse.ListingsGen.media:type_name -> complex_v1.ComplexUsecaseResponse.ListingsGen.MediaGen
	0, // 8: complex_v1.ComplexUsecaseService.InvokeComplexUsecase:input_type -> complex_v1.ComplexUsecaseRequest
	1, // 9: complex_v1.ComplexUsecaseService.InvokeComplexUsecase:output_type -> complex_v1.ComplexUsecaseResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_complex_v1_complex_usecase_proto_init() }
func file_complex_v1_complex_usecase_proto_init() {
	if File_complex_v1_complex_usecase_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_complex_v1_complex_usecase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexUsecaseRequest); i {
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
		file_complex_v1_complex_usecase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexUsecaseResponse); i {
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
		file_complex_v1_complex_usecase_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexUsecaseRequest_GetListingsGen); i {
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
		file_complex_v1_complex_usecase_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexUsecaseRequest_UserInfoGen); i {
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
		file_complex_v1_complex_usecase_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexUsecaseRequest_MetaGen); i {
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
		file_complex_v1_complex_usecase_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexUsecaseRequest_GetListingsGen_ListingIdsGen); i {
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
		file_complex_v1_complex_usecase_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexUsecaseResponse_ListingsGen); i {
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
		file_complex_v1_complex_usecase_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexUsecaseResponse_UserWalletGen); i {
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
		file_complex_v1_complex_usecase_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexUsecaseResponse_ListingsGen_InfoGen); i {
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
		file_complex_v1_complex_usecase_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexUsecaseResponse_ListingsGen_MediaGen); i {
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
			RawDescriptor: file_complex_v1_complex_usecase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_complex_v1_complex_usecase_proto_goTypes,
		DependencyIndexes: file_complex_v1_complex_usecase_proto_depIdxs,
		MessageInfos:      file_complex_v1_complex_usecase_proto_msgTypes,
	}.Build()
	File_complex_v1_complex_usecase_proto = out.File
	file_complex_v1_complex_usecase_proto_rawDesc = nil
	file_complex_v1_complex_usecase_proto_goTypes = nil
	file_complex_v1_complex_usecase_proto_depIdxs = nil
}
