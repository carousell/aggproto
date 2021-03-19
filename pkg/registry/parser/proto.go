package parser

import (
	"github.com/carousell/aggproto/pkg/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type protoContainer struct {
	pb *descriptorpb.FileDescriptorProto
	r  registry.Registry
}

func (pc *protoContainer) messages() []registry.Message {
	var msgs []registry.Message
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
