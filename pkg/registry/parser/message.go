package parser

import (
	"fmt"

	"github.com/carousell/aggproto/pkg/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type MessageContainer struct {
	MessageName        string
	PackageName        string
	MessageParent      *MessageContainer
	MessageDefinitions []registry.Message
	MessageFields      []registry.Field
}

func (m *MessageContainer) FullName() string {
	return fmt.Sprintf("%s.%s", m.PackageName, m.MessageName)
}

func (m *MessageContainer) Parent() registry.Message {
	if m.MessageParent != nil {
		return m.MessageParent
	}
	return nil
}

func (m *MessageContainer) Fields() []registry.Field {
	return m.MessageFields
}

func (m *MessageContainer) Definitions() []registry.Message {
	return m.MessageDefinitions
}

func (m *MessageContainer) Package() string {
	return m.PackageName
}

func (m *MessageContainer) Name() string {
	return m.MessageName
}

func parseMessage(r registry.Registry, packageName string, msgType *descriptorpb.DescriptorProto) *MessageContainer {
	if msgType == nil {
		return nil
	}
	msg := &MessageContainer{
		PackageName: packageName,
		MessageName: msgType.GetName(),
	}
	var definitions []registry.Message
	for _, dp := range msgType.NestedType {
		subMsg := parseMessage(r, fmt.Sprintf("%s.%s", packageName, msgType.GetName()), dp)
		subMsg.MessageParent = msg
		if subMsg != nil {
			definitions = append(definitions, subMsg)
		}
	}
	msg.MessageDefinitions = definitions
	return msg
}
func populateMessageField(r registry.Registry, packageName string, msgType *descriptorpb.DescriptorProto) {
	msgName := fmt.Sprintf("%s.%s", packageName, msgType.GetName())
	msgs := r.ListMessages(registry.LMOWithFullName(msgName))
	if len(msgs) != 1 {
		panic(fmt.Sprintf("message not found %s", msgName))
	}
	msg := msgs[0].(*MessageContainer)
	var fields []registry.Field
	for _, ft := range msgType.Field {
		fields = append(fields, parseField(r, msg, ft))
	}
	msg.MessageFields = fields

	for _, dp := range msgType.NestedType {
		populateMessageField(r, msgName, dp)
	}
}
