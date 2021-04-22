// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: union_mask_v1/masked_listing_w_wallet.proto

package union_mask_v1

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

type MaskedListingWWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetUserWalletRequest *MaskedListingWWalletRequest_GetUserWalletRequestGen `protobuf:"bytes,1,opt,name=get_user_wallet_request,json=getUserWalletRequest,proto3" json:"get_user_wallet_request,omitempty"`
	GetListingRequest    *MaskedListingWWalletRequest_GetListingRequestGen    `protobuf:"bytes,2,opt,name=get_listing_request,json=getListingRequest,proto3" json:"get_listing_request,omitempty"`
}

func (x *MaskedListingWWalletRequest) Reset() {
	*x = MaskedListingWWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskedListingWWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskedListingWWalletRequest) ProtoMessage() {}

func (x *MaskedListingWWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskedListingWWalletRequest.ProtoReflect.Descriptor instead.
func (*MaskedListingWWalletRequest) Descriptor() ([]byte, []int) {
	return file_union_mask_v1_masked_listing_w_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *MaskedListingWWalletRequest) GetGetUserWalletRequest() *MaskedListingWWalletRequest_GetUserWalletRequestGen {
	if x != nil {
		return x.GetUserWalletRequest
	}
	return nil
}

func (x *MaskedListingWWalletRequest) GetGetListingRequest() *MaskedListingWWalletRequest_GetListingRequestGen {
	if x != nil {
		return x.GetListingRequest
	}
	return nil
}

type MaskedListingWWalletResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listing *MaskedListingWWalletResponse_ListingGen `protobuf:"bytes,1,opt,name=listing,proto3" json:"listing,omitempty"`
	Wallet  *MaskedListingWWalletResponse_WalletGen  `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
}

func (x *MaskedListingWWalletResponse) Reset() {
	*x = MaskedListingWWalletResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskedListingWWalletResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskedListingWWalletResponse) ProtoMessage() {}

func (x *MaskedListingWWalletResponse) ProtoReflect() protoreflect.Message {
	mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskedListingWWalletResponse.ProtoReflect.Descriptor instead.
func (*MaskedListingWWalletResponse) Descriptor() ([]byte, []int) {
	return file_union_mask_v1_masked_listing_w_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *MaskedListingWWalletResponse) GetListing() *MaskedListingWWalletResponse_ListingGen {
	if x != nil {
		return x.Listing
	}
	return nil
}

func (x *MaskedListingWWalletResponse) GetWallet() *MaskedListingWWalletResponse_WalletGen {
	if x != nil {
		return x.Wallet
	}
	return nil
}

type MaskedListingWWalletRequest_GetUserWalletRequestGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *MaskedListingWWalletRequest_GetUserWalletRequestGen) Reset() {
	*x = MaskedListingWWalletRequest_GetUserWalletRequestGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskedListingWWalletRequest_GetUserWalletRequestGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskedListingWWalletRequest_GetUserWalletRequestGen) ProtoMessage() {}

func (x *MaskedListingWWalletRequest_GetUserWalletRequestGen) ProtoReflect() protoreflect.Message {
	mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskedListingWWalletRequest_GetUserWalletRequestGen.ProtoReflect.Descriptor instead.
func (*MaskedListingWWalletRequest_GetUserWalletRequestGen) Descriptor() ([]byte, []int) {
	return file_union_mask_v1_masked_listing_w_wallet_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MaskedListingWWalletRequest_GetUserWalletRequestGen) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type MaskedListingWWalletRequest_GetListingRequestGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListingId string `protobuf:"bytes,1,opt,name=listing_id,json=listingId,proto3" json:"listing_id,omitempty"`
}

func (x *MaskedListingWWalletRequest_GetListingRequestGen) Reset() {
	*x = MaskedListingWWalletRequest_GetListingRequestGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskedListingWWalletRequest_GetListingRequestGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskedListingWWalletRequest_GetListingRequestGen) ProtoMessage() {}

func (x *MaskedListingWWalletRequest_GetListingRequestGen) ProtoReflect() protoreflect.Message {
	mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskedListingWWalletRequest_GetListingRequestGen.ProtoReflect.Descriptor instead.
func (*MaskedListingWWalletRequest_GetListingRequestGen) Descriptor() ([]byte, []int) {
	return file_union_mask_v1_masked_listing_w_wallet_proto_rawDescGZIP(), []int{0, 1}
}

func (x *MaskedListingWWalletRequest_GetListingRequestGen) GetListingId() string {
	if x != nil {
		return x.ListingId
	}
	return ""
}

type MaskedListingWWalletResponse_ListingGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetListingResponse *MaskedListingWWalletResponse_ListingGen_GetListingResponseGen `protobuf:"bytes,1,opt,name=GetListingResponse,proto3" json:"GetListingResponse,omitempty"`
}

func (x *MaskedListingWWalletResponse_ListingGen) Reset() {
	*x = MaskedListingWWalletResponse_ListingGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskedListingWWalletResponse_ListingGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskedListingWWalletResponse_ListingGen) ProtoMessage() {}

func (x *MaskedListingWWalletResponse_ListingGen) ProtoReflect() protoreflect.Message {
	mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskedListingWWalletResponse_ListingGen.ProtoReflect.Descriptor instead.
func (*MaskedListingWWalletResponse_ListingGen) Descriptor() ([]byte, []int) {
	return file_union_mask_v1_masked_listing_w_wallet_proto_rawDescGZIP(), []int{1, 0}
}

func (x *MaskedListingWWalletResponse_ListingGen) GetGetListingResponse() *MaskedListingWWalletResponse_ListingGen_GetListingResponseGen {
	if x != nil {
		return x.GetListingResponse
	}
	return nil
}

type MaskedListingWWalletResponse_WalletGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetUserWalletResponse *MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen `protobuf:"bytes,1,opt,name=GetUserWalletResponse,proto3" json:"GetUserWalletResponse,omitempty"`
}

func (x *MaskedListingWWalletResponse_WalletGen) Reset() {
	*x = MaskedListingWWalletResponse_WalletGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskedListingWWalletResponse_WalletGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskedListingWWalletResponse_WalletGen) ProtoMessage() {}

func (x *MaskedListingWWalletResponse_WalletGen) ProtoReflect() protoreflect.Message {
	mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskedListingWWalletResponse_WalletGen.ProtoReflect.Descriptor instead.
func (*MaskedListingWWalletResponse_WalletGen) Descriptor() ([]byte, []int) {
	return file_union_mask_v1_masked_listing_w_wallet_proto_rawDescGZIP(), []int{1, 1}
}

func (x *MaskedListingWWalletResponse_WalletGen) GetGetUserWalletResponse() *MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen {
	if x != nil {
		return x.GetUserWalletResponse
	}
	return nil
}

type MaskedListingWWalletResponse_ListingGen_GetListingResponseGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listing *MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen `protobuf:"bytes,1,opt,name=listing,proto3" json:"listing,omitempty"`
}

func (x *MaskedListingWWalletResponse_ListingGen_GetListingResponseGen) Reset() {
	*x = MaskedListingWWalletResponse_ListingGen_GetListingResponseGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskedListingWWalletResponse_ListingGen_GetListingResponseGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskedListingWWalletResponse_ListingGen_GetListingResponseGen) ProtoMessage() {}

func (x *MaskedListingWWalletResponse_ListingGen_GetListingResponseGen) ProtoReflect() protoreflect.Message {
	mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskedListingWWalletResponse_ListingGen_GetListingResponseGen.ProtoReflect.Descriptor instead.
func (*MaskedListingWWalletResponse_ListingGen_GetListingResponseGen) Descriptor() ([]byte, []int) {
	return file_union_mask_v1_masked_listing_w_wallet_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *MaskedListingWWalletResponse_ListingGen_GetListingResponseGen) GetListing() *MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen {
	if x != nil {
		return x.Listing
	}
	return nil
}

type MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen) Reset() {
	*x = MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen) ProtoMessage() {}

func (x *MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen) ProtoReflect() protoreflect.Message {
	mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen.ProtoReflect.Descriptor instead.
func (*MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen) Descriptor() ([]byte, []int) {
	return file_union_mask_v1_masked_listing_w_wallet_proto_rawDescGZIP(), []int{1, 0, 0, 0}
}

func (x *MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserWallet *MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen_UserWalletGen `protobuf:"bytes,1,opt,name=user_wallet,json=userWallet,proto3" json:"user_wallet,omitempty"`
}

func (x *MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen) Reset() {
	*x = MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen) ProtoMessage() {}

func (x *MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen) ProtoReflect() protoreflect.Message {
	mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen.ProtoReflect.Descriptor instead.
func (*MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen) Descriptor() ([]byte, []int) {
	return file_union_mask_v1_masked_listing_w_wallet_proto_rawDescGZIP(), []int{1, 1, 0}
}

func (x *MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen) GetUserWallet() *MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen_UserWalletGen {
	if x != nil {
		return x.UserWallet
	}
	return nil
}

type MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen_UserWalletGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinBalance int64 `protobuf:"varint,1,opt,name=coin_balance,json=coinBalance,proto3" json:"coin_balance,omitempty"`
}

func (x *MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen_UserWalletGen) Reset() {
	*x = MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen_UserWalletGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen_UserWalletGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen_UserWalletGen) ProtoMessage() {
}

func (x *MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen_UserWalletGen) ProtoReflect() protoreflect.Message {
	mi := &file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen_UserWalletGen.ProtoReflect.Descriptor instead.
func (*MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen_UserWalletGen) Descriptor() ([]byte, []int) {
	return file_union_mask_v1_masked_listing_w_wallet_proto_rawDescGZIP(), []int{1, 1, 0, 0}
}

func (x *MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen_UserWalletGen) GetCoinBalance() int64 {
	if x != nil {
		return x.CoinBalance
	}
	return 0
}

var File_union_mask_v1_masked_listing_w_wallet_proto protoreflect.FileDescriptor

var file_union_mask_v1_masked_listing_w_wallet_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x31, 0x2f,
	0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x77,
	0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75,
	0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x31, 0x22, 0xf4, 0x02, 0x0a,
	0x1b, 0x4d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x57, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x79, 0x0a, 0x17,
	0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e,
	0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x73, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x57, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65,
	0x6e, 0x52, 0x14, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x6f, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x57, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x47, 0x65, 0x6e, 0x52, 0x11, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x47, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x35, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x47, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x22, 0x84, 0x07, 0x0a, 0x1c, 0x4d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x57, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x57, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x65, 0x6e, 0x52, 0x07, 0x6c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x4d, 0x0a, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x57, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x52, 0x06, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x1a, 0xdd, 0x02, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x47, 0x65, 0x6e, 0x12, 0x7c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x4c, 0x2e, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x31,
	0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x57, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x6e, 0x52, 0x12,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x1a, 0xd0, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x6e, 0x12, 0x71, 0x0a, 0x07,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x57, 0x2e,
	0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x73, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x57, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x47, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x47, 0x65, 0x6e, 0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x1a,
	0x44, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x65, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xe2, 0x02, 0x0a, 0x09, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x47, 0x65, 0x6e, 0x12, 0x84, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x57, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x47, 0x65, 0x6e, 0x52, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0xcd, 0x01, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x6e, 0x12, 0x7d, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5c, 0x2e, 0x75,
	0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x73,
	0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x57, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x47, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x6e, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x1a, 0x32, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x69, 0x6e, 0x5f,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63,
	0x6f, 0x69, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x32, 0x94, 0x01, 0x0a, 0x1b, 0x4d,
	0x61, 0x73, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x57, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x75, 0x0a, 0x1a, 0x49, 0x6e,
	0x76, 0x6f, 0x6b, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x57, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x75, 0x6e, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x57, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x57, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x6c, 0x2f, 0x61, 0x67, 0x67, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x4f, 0x75,
	0x74, 0x2f, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_union_mask_v1_masked_listing_w_wallet_proto_rawDescOnce sync.Once
	file_union_mask_v1_masked_listing_w_wallet_proto_rawDescData = file_union_mask_v1_masked_listing_w_wallet_proto_rawDesc
)

func file_union_mask_v1_masked_listing_w_wallet_proto_rawDescGZIP() []byte {
	file_union_mask_v1_masked_listing_w_wallet_proto_rawDescOnce.Do(func() {
		file_union_mask_v1_masked_listing_w_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_union_mask_v1_masked_listing_w_wallet_proto_rawDescData)
	})
	return file_union_mask_v1_masked_listing_w_wallet_proto_rawDescData
}

var file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_union_mask_v1_masked_listing_w_wallet_proto_goTypes = []interface{}{
	(*MaskedListingWWalletRequest)(nil),                                                   // 0: union_mask_v1.MaskedListingWWalletRequest
	(*MaskedListingWWalletResponse)(nil),                                                  // 1: union_mask_v1.MaskedListingWWalletResponse
	(*MaskedListingWWalletRequest_GetUserWalletRequestGen)(nil),                           // 2: union_mask_v1.MaskedListingWWalletRequest.GetUserWalletRequestGen
	(*MaskedListingWWalletRequest_GetListingRequestGen)(nil),                              // 3: union_mask_v1.MaskedListingWWalletRequest.GetListingRequestGen
	(*MaskedListingWWalletResponse_ListingGen)(nil),                                       // 4: union_mask_v1.MaskedListingWWalletResponse.ListingGen
	(*MaskedListingWWalletResponse_WalletGen)(nil),                                        // 5: union_mask_v1.MaskedListingWWalletResponse.WalletGen
	(*MaskedListingWWalletResponse_ListingGen_GetListingResponseGen)(nil),                 // 6: union_mask_v1.MaskedListingWWalletResponse.ListingGen.GetListingResponseGen
	(*MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen)(nil),      // 7: union_mask_v1.MaskedListingWWalletResponse.ListingGen.GetListingResponseGen.ListingGen
	(*MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen)(nil),               // 8: union_mask_v1.MaskedListingWWalletResponse.WalletGen.GetUserWalletResponseGen
	(*MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen_UserWalletGen)(nil), // 9: union_mask_v1.MaskedListingWWalletResponse.WalletGen.GetUserWalletResponseGen.UserWalletGen
}
var file_union_mask_v1_masked_listing_w_wallet_proto_depIdxs = []int32{
	2, // 0: union_mask_v1.MaskedListingWWalletRequest.get_user_wallet_request:type_name -> union_mask_v1.MaskedListingWWalletRequest.GetUserWalletRequestGen
	3, // 1: union_mask_v1.MaskedListingWWalletRequest.get_listing_request:type_name -> union_mask_v1.MaskedListingWWalletRequest.GetListingRequestGen
	4, // 2: union_mask_v1.MaskedListingWWalletResponse.listing:type_name -> union_mask_v1.MaskedListingWWalletResponse.ListingGen
	5, // 3: union_mask_v1.MaskedListingWWalletResponse.wallet:type_name -> union_mask_v1.MaskedListingWWalletResponse.WalletGen
	6, // 4: union_mask_v1.MaskedListingWWalletResponse.ListingGen.GetListingResponse:type_name -> union_mask_v1.MaskedListingWWalletResponse.ListingGen.GetListingResponseGen
	8, // 5: union_mask_v1.MaskedListingWWalletResponse.WalletGen.GetUserWalletResponse:type_name -> union_mask_v1.MaskedListingWWalletResponse.WalletGen.GetUserWalletResponseGen
	7, // 6: union_mask_v1.MaskedListingWWalletResponse.ListingGen.GetListingResponseGen.listing:type_name -> union_mask_v1.MaskedListingWWalletResponse.ListingGen.GetListingResponseGen.ListingGen
	9, // 7: union_mask_v1.MaskedListingWWalletResponse.WalletGen.GetUserWalletResponseGen.user_wallet:type_name -> union_mask_v1.MaskedListingWWalletResponse.WalletGen.GetUserWalletResponseGen.UserWalletGen
	0, // 8: union_mask_v1.MaskedListingWWalletService.InvokeMaskedListingWWallet:input_type -> union_mask_v1.MaskedListingWWalletRequest
	1, // 9: union_mask_v1.MaskedListingWWalletService.InvokeMaskedListingWWallet:output_type -> union_mask_v1.MaskedListingWWalletResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_union_mask_v1_masked_listing_w_wallet_proto_init() }
func file_union_mask_v1_masked_listing_w_wallet_proto_init() {
	if File_union_mask_v1_masked_listing_w_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskedListingWWalletRequest); i {
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
		file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskedListingWWalletResponse); i {
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
		file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskedListingWWalletRequest_GetUserWalletRequestGen); i {
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
		file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskedListingWWalletRequest_GetListingRequestGen); i {
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
		file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskedListingWWalletResponse_ListingGen); i {
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
		file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskedListingWWalletResponse_WalletGen); i {
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
		file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskedListingWWalletResponse_ListingGen_GetListingResponseGen); i {
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
		file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskedListingWWalletResponse_ListingGen_GetListingResponseGen_ListingGen); i {
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
		file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen); i {
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
		file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskedListingWWalletResponse_WalletGen_GetUserWalletResponseGen_UserWalletGen); i {
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
			RawDescriptor: file_union_mask_v1_masked_listing_w_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_union_mask_v1_masked_listing_w_wallet_proto_goTypes,
		DependencyIndexes: file_union_mask_v1_masked_listing_w_wallet_proto_depIdxs,
		MessageInfos:      file_union_mask_v1_masked_listing_w_wallet_proto_msgTypes,
	}.Build()
	File_union_mask_v1_masked_listing_w_wallet_proto = out.File
	file_union_mask_v1_masked_listing_w_wallet_proto_rawDesc = nil
	file_union_mask_v1_masked_listing_w_wallet_proto_goTypes = nil
	file_union_mask_v1_masked_listing_w_wallet_proto_depIdxs = nil
}
