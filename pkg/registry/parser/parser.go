//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
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
