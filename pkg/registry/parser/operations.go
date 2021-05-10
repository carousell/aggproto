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
	"fmt"
	"log"
	"strings"

	"github.com/carousell/aggproto/pkg/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type ServiceContainer struct {
	UnaryOps    []*UnaryOperationContainer
	PackageName string
	ServiceName string
}

func (s *ServiceContainer) UnaryOperations() []registry.UnaryOperation {
	ops := make([]registry.UnaryOperation, len(s.UnaryOps))
	for i, op := range s.UnaryOps {
		ops[i] = op
	}
	return ops
}

func (s *ServiceContainer) Package() string {
	return s.PackageName
}

func (s *ServiceContainer) Name() string {
	return s.ServiceName
}

func parseService(r registry.Registry, svcType *descriptorpb.ServiceDescriptorProto, packageName string) *ServiceContainer {
	sc := &ServiceContainer{
		PackageName: packageName,
		ServiceName: svcType.GetName(),
	}
	var ops []*UnaryOperationContainer
	for _, m := range svcType.Method {
		op := parseOperations(r, m, sc)
		if op != nil {
			ops = append(ops, op)
		}
	}
	sc.UnaryOps = append(sc.UnaryOps, ops...)
	return sc
}

type UnaryOperationContainer struct {
	InputMsg  registry.Message
	OutputMsg registry.Message
	OpName    string
	OpCtx     *ServiceContainer
}

func (o *UnaryOperationContainer) Input() registry.Message {
	return o.InputMsg
}

func (o *UnaryOperationContainer) Output() registry.Message {
	return o.OutputMsg
}

func (o *UnaryOperationContainer) Name() string {
	return o.OpName
}

func (o *UnaryOperationContainer) Context() registry.Service {
	return o.OpCtx
}

func (o *UnaryOperationContainer) FullName() string {
	return fmt.Sprintf("%s.%s/%s", o.OpCtx.PackageName, o.OpCtx.ServiceName, o.OpName)
}

func parseOperations(r registry.Registry, method *descriptorpb.MethodDescriptorProto, sc *ServiceContainer) *UnaryOperationContainer {
	if method.GetClientStreaming() || method.GetServerStreaming() {
		return nil
	}
	inputMessages := r.ListMessages(registry.LMOWithFullName(strings.TrimLeft(method.GetInputType(), ".")))
	if len(inputMessages) != 1 {
		log.Fatalf("expected an input message for name %s found %d", method.GetInputType(), len(inputMessages))
	}
	outputMessages := r.ListMessages(registry.LMOWithFullName(strings.TrimLeft(method.GetOutputType(), ".")))
	if len(outputMessages) != 1 {
		log.Fatalf("expected an output message for name %s found %d", method.GetOutputType(), len(outputMessages))
	}
	return &UnaryOperationContainer{
		InputMsg:  inputMessages[0],
		OutputMsg: outputMessages[0],
		OpName:    method.GetName(),
		OpCtx:     sc,
	}

}
