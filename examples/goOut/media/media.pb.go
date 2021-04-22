// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: media/media.proto

package media

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

type Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhotoUrl string `protobuf:"bytes,1,opt,name=photo_url,json=photoUrl,proto3" json:"photo_url,omitempty"`
}

func (x *Media) Reset() {
	*x = Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_media_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_media_media_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_media_media_proto_rawDescGZIP(), []int{0}
}

func (x *Media) GetPhotoUrl() string {
	if x != nil {
		return x.PhotoUrl
	}
	return ""
}

type GetMediaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaId string `protobuf:"bytes,1,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
}

func (x *GetMediaRequest) Reset() {
	*x = GetMediaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_media_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMediaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMediaRequest) ProtoMessage() {}

func (x *GetMediaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_media_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMediaRequest.ProtoReflect.Descriptor instead.
func (*GetMediaRequest) Descriptor() ([]byte, []int) {
	return file_media_media_proto_rawDescGZIP(), []int{1}
}

func (x *GetMediaRequest) GetMediaId() string {
	if x != nil {
		return x.MediaId
	}
	return ""
}

type GetMediaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Media *Media `protobuf:"bytes,1,opt,name=media,proto3" json:"media,omitempty"`
}

func (x *GetMediaResponse) Reset() {
	*x = GetMediaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_media_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMediaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMediaResponse) ProtoMessage() {}

func (x *GetMediaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_media_media_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMediaResponse.ProtoReflect.Descriptor instead.
func (*GetMediaResponse) Descriptor() ([]byte, []int) {
	return file_media_media_proto_rawDescGZIP(), []int{2}
}

func (x *GetMediaResponse) GetMedia() *Media {
	if x != nil {
		return x.Media
	}
	return nil
}

type BulkGetMediaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaIds  []string `protobuf:"bytes,1,rep,name=media_ids,json=mediaIds,proto3" json:"media_ids,omitempty"`
	UserId    string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MediaType []string `protobuf:"bytes,3,rep,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
}

func (x *BulkGetMediaRequest) Reset() {
	*x = BulkGetMediaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_media_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkGetMediaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkGetMediaRequest) ProtoMessage() {}

func (x *BulkGetMediaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_media_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkGetMediaRequest.ProtoReflect.Descriptor instead.
func (*BulkGetMediaRequest) Descriptor() ([]byte, []int) {
	return file_media_media_proto_rawDescGZIP(), []int{3}
}

func (x *BulkGetMediaRequest) GetMediaIds() []string {
	if x != nil {
		return x.MediaIds
	}
	return nil
}

func (x *BulkGetMediaRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BulkGetMediaRequest) GetMediaType() []string {
	if x != nil {
		return x.MediaType
	}
	return nil
}

type BulkGetMediaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Medias []*Media `protobuf:"bytes,1,rep,name=medias,proto3" json:"medias,omitempty"`
}

func (x *BulkGetMediaResponse) Reset() {
	*x = BulkGetMediaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_media_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkGetMediaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkGetMediaResponse) ProtoMessage() {}

func (x *BulkGetMediaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_media_media_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkGetMediaResponse.ProtoReflect.Descriptor instead.
func (*BulkGetMediaResponse) Descriptor() ([]byte, []int) {
	return file_media_media_proto_rawDescGZIP(), []int{4}
}

func (x *BulkGetMediaResponse) GetMedias() []*Media {
	if x != nil {
		return x.Medias
	}
	return nil
}

var File_media_media_proto protoreflect.FileDescriptor

var file_media_media_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x22, 0x24, 0x0a, 0x05, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c,
	0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x22, 0x36,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52,
	0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x22, 0x6a, 0x0a, 0x13, 0x42, 0x75, 0x6c, 0x6b, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x3c, 0x0a, 0x14, 0x42, 0x75, 0x6c, 0x6b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73,
	0x32, 0x98, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x16, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x49, 0x0a, 0x0c, 0x42, 0x75, 0x6c, 0x6b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x12, 0x1a, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73,
	0x65, 0x6c, 0x6c, 0x2f, 0x61, 0x67, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x4f, 0x75, 0x74, 0x2f, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_media_media_proto_rawDescOnce sync.Once
	file_media_media_proto_rawDescData = file_media_media_proto_rawDesc
)

func file_media_media_proto_rawDescGZIP() []byte {
	file_media_media_proto_rawDescOnce.Do(func() {
		file_media_media_proto_rawDescData = protoimpl.X.CompressGZIP(file_media_media_proto_rawDescData)
	})
	return file_media_media_proto_rawDescData
}

var file_media_media_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_media_media_proto_goTypes = []interface{}{
	(*Media)(nil),                // 0: media.Media
	(*GetMediaRequest)(nil),      // 1: media.GetMediaRequest
	(*GetMediaResponse)(nil),     // 2: media.GetMediaResponse
	(*BulkGetMediaRequest)(nil),  // 3: media.BulkGetMediaRequest
	(*BulkGetMediaResponse)(nil), // 4: media.BulkGetMediaResponse
}
var file_media_media_proto_depIdxs = []int32{
	0, // 0: media.GetMediaResponse.media:type_name -> media.Media
	0, // 1: media.BulkGetMediaResponse.medias:type_name -> media.Media
	1, // 2: media.MediaService.GetMedia:input_type -> media.GetMediaRequest
	3, // 3: media.MediaService.BulkGetMedia:input_type -> media.BulkGetMediaRequest
	2, // 4: media.MediaService.GetMedia:output_type -> media.GetMediaResponse
	4, // 5: media.MediaService.BulkGetMedia:output_type -> media.BulkGetMediaResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_media_media_proto_init() }
func file_media_media_proto_init() {
	if File_media_media_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_media_media_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Media); i {
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
		file_media_media_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMediaRequest); i {
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
		file_media_media_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMediaResponse); i {
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
		file_media_media_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkGetMediaRequest); i {
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
		file_media_media_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkGetMediaResponse); i {
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
			RawDescriptor: file_media_media_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_media_media_proto_goTypes,
		DependencyIndexes: file_media_media_proto_depIdxs,
		MessageInfos:      file_media_media_proto_msgTypes,
	}.Build()
	File_media_media_proto = out.File
	file_media_media_proto_rawDesc = nil
	file_media_media_proto_goTypes = nil
	file_media_media_proto_depIdxs = nil
}
