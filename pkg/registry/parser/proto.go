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
	"github.com/carousell/aggproto/pkg/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type protoContainer struct {
	pb *descriptorpb.FileDescriptorProto
	r  registry.Registry
}

func (pc *protoContainer) messages() []*MessageContainer {
	var msgs []*MessageContainer
	for _, msgType := range pc.pb.GetMessageType() {
		msgs = append(msgs, parseMessage(pc.r, pc.pb.GetPackage(), msgType))
	}
	return msgs
}

func (pc *protoContainer) populateMessageFields() {
	for _, msgType := range pc.pb.GetMessageType() {
		populateMessageField(pc.r, pc.pb.GetPackage(), msgType)
	}
}

func (pc *protoContainer) operations() []*UnaryOperationContainer {
	var svcs []*ServiceContainer
	for _, svcType := range pc.pb.Service {
		svcs = append(svcs, parseService(pc.r, svcType, pc.pb.GetPackage()))
	}
	var ops []*UnaryOperationContainer
	for _, svc := range svcs {
		ops = append(ops, svc.UnaryOps...)
	}
	return ops
}

//func (pc *protoContainer) services() []*ServiceContainer {
//	var svcs []*ServiceContainer
//	for _, svcType := range pc.pb.Service {
//		svcs = append(svcs, &ServiceContainer{
//			PackageName: pc.pb.GetPackage(),
//			ServiceName: svcType.GetName(),
//			scvType:     svcType,
//			parent:      pc,
//		})
//	}
//	return svcs
//}
