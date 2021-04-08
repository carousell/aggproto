// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: static_primitives_v1/mock_listing.proto

package static_primitives_v1

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

type MockListingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MockListingRequest) Reset() {
	*x = MockListingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_static_primitives_v1_mock_listing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MockListingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MockListingRequest) ProtoMessage() {}

func (x *MockListingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_static_primitives_v1_mock_listing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MockListingRequest.ProtoReflect.Descriptor instead.
func (*MockListingRequest) Descriptor() ([]byte, []int) {
	return file_static_primitives_v1_mock_listing_proto_rawDescGZIP(), []int{0}
}

type MockListingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listing *MockListingResponse_ListingGen `protobuf:"bytes,1,opt,name=listing,proto3" json:"listing,omitempty"`
}

func (x *MockListingResponse) Reset() {
	*x = MockListingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_static_primitives_v1_mock_listing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MockListingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MockListingResponse) ProtoMessage() {}

func (x *MockListingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_static_primitives_v1_mock_listing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MockListingResponse.ProtoReflect.Descriptor instead.
func (*MockListingResponse) Descriptor() ([]byte, []int) {
	return file_static_primitives_v1_mock_listing_proto_rawDescGZIP(), []int{1}
}

func (x *MockListingResponse) GetListing() *MockListingResponse_ListingGen {
	if x != nil {
		return x.Listing
	}
	return nil
}

type MockListingResponse_ListingGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *MockListingResponse_ListingGen) Reset() {
	*x = MockListingResponse_ListingGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_static_primitives_v1_mock_listing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MockListingResponse_ListingGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MockListingResponse_ListingGen) ProtoMessage() {}

func (x *MockListingResponse_ListingGen) ProtoReflect() protoreflect.Message {
	mi := &file_static_primitives_v1_mock_listing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MockListingResponse_ListingGen.ProtoReflect.Descriptor instead.
func (*MockListingResponse_ListingGen) Descriptor() ([]byte, []int) {
	return file_static_primitives_v1_mock_listing_proto_rawDescGZIP(), []int{1, 0}
}

func (x *MockListingResponse_ListingGen) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MockListingResponse_ListingGen) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_static_primitives_v1_mock_listing_proto protoreflect.FileDescriptor

var file_static_primitives_v1_mock_listing_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x22,
	0x14, 0x0a, 0x12, 0x4d, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xab, 0x01, 0x0a, 0x13, 0x4d, 0x6f, 0x63, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x47, 0x65, 0x6e, 0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x44, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x32, 0x7e, 0x0a, 0x12, 0x4d, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x11, 0x49, 0x6e, 0x76,
	0x6f, 0x6b, 0x65, 0x4d, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x28,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e,
	0x4d, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x6c, 0x2f, 0x61, 0x67, 0x67, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x6f,
	0x4f, 0x75, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_static_primitives_v1_mock_listing_proto_rawDescOnce sync.Once
	file_static_primitives_v1_mock_listing_proto_rawDescData = file_static_primitives_v1_mock_listing_proto_rawDesc
)

func file_static_primitives_v1_mock_listing_proto_rawDescGZIP() []byte {
	file_static_primitives_v1_mock_listing_proto_rawDescOnce.Do(func() {
		file_static_primitives_v1_mock_listing_proto_rawDescData = protoimpl.X.CompressGZIP(file_static_primitives_v1_mock_listing_proto_rawDescData)
	})
	return file_static_primitives_v1_mock_listing_proto_rawDescData
}

var file_static_primitives_v1_mock_listing_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_static_primitives_v1_mock_listing_proto_goTypes = []interface{}{
	(*MockListingRequest)(nil),             // 0: static_primitives_v1.MockListingRequest
	(*MockListingResponse)(nil),            // 1: static_primitives_v1.MockListingResponse
	(*MockListingResponse_ListingGen)(nil), // 2: static_primitives_v1.MockListingResponse.ListingGen
}
var file_static_primitives_v1_mock_listing_proto_depIdxs = []int32{
	2, // 0: static_primitives_v1.MockListingResponse.listing:type_name -> static_primitives_v1.MockListingResponse.ListingGen
	0, // 1: static_primitives_v1.MockListingService.InvokeMockListing:input_type -> static_primitives_v1.MockListingRequest
	1, // 2: static_primitives_v1.MockListingService.InvokeMockListing:output_type -> static_primitives_v1.MockListingResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_static_primitives_v1_mock_listing_proto_init() }
func file_static_primitives_v1_mock_listing_proto_init() {
	if File_static_primitives_v1_mock_listing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_static_primitives_v1_mock_listing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MockListingRequest); i {
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
		file_static_primitives_v1_mock_listing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MockListingResponse); i {
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
		file_static_primitives_v1_mock_listing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MockListingResponse_ListingGen); i {
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
			RawDescriptor: file_static_primitives_v1_mock_listing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_static_primitives_v1_mock_listing_proto_goTypes,
		DependencyIndexes: file_static_primitives_v1_mock_listing_proto_depIdxs,
		MessageInfos:      file_static_primitives_v1_mock_listing_proto_msgTypes,
	}.Build()
	File_static_primitives_v1_mock_listing_proto = out.File
	file_static_primitives_v1_mock_listing_proto_rawDesc = nil
	file_static_primitives_v1_mock_listing_proto_goTypes = nil
	file_static_primitives_v1_mock_listing_proto_depIdxs = nil
}
