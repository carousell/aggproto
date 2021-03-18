package parser

import (
	"github.com/carousell/aggproto/pkg/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type parser struct {
	r registry.Registry
}

func (p *parser) AddProtoContainer(file *descriptorpb.FileDescriptorProto) error {
	pc := &protoContainer{file, p.r}
	p.r.AddMessages(pc.messages()...)
	return nil
}
