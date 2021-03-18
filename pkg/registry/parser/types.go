package parser

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

type Parser interface {
	AddProtoContainer(file *descriptorpb.FileDescriptorProto) error
}
