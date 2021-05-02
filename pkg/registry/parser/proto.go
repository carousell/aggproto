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

func (pc *protoContainer) enums() []*EnumContainer {
	var enums []*EnumContainer
	for _, enumType := range pc.pb.GetEnumType() {
		enums = append(enums, parseEnum(pc.r, pc.pb.GetPackage(), enumType))
	}
	return enums
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
