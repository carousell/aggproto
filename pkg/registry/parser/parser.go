package parser

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

type parser struct {
	r persistentRegistry
}

func (p *parser) SaveRegistry() error {
	return p.r.save()
}

func (p *parser) AddProtoContainer(file *descriptorpb.FileDescriptorProto) error {
	p.r.addFileDescriptor(file)
	return nil
}

func LoadParser(registryDirName string) Parser {
	return &parser{r: Load(registryDirName)}
}
