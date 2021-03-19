package parser

import (
	"fmt"

	"github.com/carousell/aggproto/pkg/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type messageContainer struct {
	name        string
	packageName string
	parent      *messageContainer
	definitions []registry.Message
	fields      []registry.Field
}

func (m *messageContainer) FullName() string {
	return fmt.Sprintf("%s.%s", m.packageName, m.name)
}

func (m *messageContainer) Parent() registry.Message {
	if m.parent != nil {
		return m.parent
	}
	return nil
}

func (m *messageContainer) Fields() []registry.Field {
	return m.fields
}

func (m *messageContainer) Definitions() []registry.Message {
	return m.definitions
}

func (m *messageContainer) Package() string {
	return m.packageName
}

func (m *messageContainer) Name() string {
	return m.name
}

func parseMessage(r registry.Registry, packageName string, msgType *descriptorpb.DescriptorProto) *messageContainer {
	if msgType == nil {
		return nil
	}
	msg := &messageContainer{
		packageName: packageName,
		name:        msgType.GetName(),
	}
	var definitions []registry.Message
	for _, dp := range msgType.NestedType {
		subMsg := parseMessage(r, fmt.Sprintf("%s.%s", packageName, msgType.GetName()), dp)
		subMsg.parent = msg
		if subMsg != nil {
			definitions = append(definitions, subMsg)
		}
	}
	msg.definitions = definitions
	return msg
}
func populateMessageField(r registry.Registry, packageName string, msgType *descriptorpb.DescriptorProto) {
	msgName := fmt.Sprintf("%s.%s", packageName, msgType.GetName())
	msgs := r.ListMessages(registry.LMOWithFullName(msgName))
	if len(msgs) != 1 {
		panic(fmt.Sprintf("message not found %s", msgName))
	}
	msg := msgs[0].(*messageContainer)
	var fields []registry.Field
	for _, ft := range msgType.Field {
		fields = append(fields, parseField(r, msg, ft))
	}
	msg.fields = fields

	for _, dp := range msgType.NestedType {
		populateMessageField(r, msgName, dp)
	}
}
