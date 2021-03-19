package parser

import (
	"fmt"
	"strings"

	"github.com/carousell/aggproto/pkg/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type messageField struct {
	fieldType registry.FieldType
	name      string
	message   registry.Message
	context   registry.Message
	repeated  bool
}

func (m *messageField) Type() registry.FieldType {
	return m.fieldType
}

func (m *messageField) Name() string {
	return m.name
}

func (m *messageField) Message() registry.Message {
	return m.message
}

func (m *messageField) Context() registry.Message {
	return m.context
}

func (m *messageField) Repeated() bool {
	return m.repeated
}

func parseField(r registry.Registry, msgContext *messageContainer, fieldDescriptorProto *descriptorpb.FieldDescriptorProto) registry.Field {
	if fieldDescriptorProto.OneofIndex != nil {
		// TODO
		panic("unhandled oneof")
	}
	field := &messageField{
		context: msgContext,
		name:    fieldDescriptorProto.GetName(),
	}
	switch fieldDescriptorProto.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		field.fieldType = registry.FieldTypeMessage
		tn := fieldDescriptorProto.GetTypeName()
		if strings.HasPrefix(tn, ".") {
			tn = strings.Replace(tn, ".", "", 1)
		}
		msgs := r.ListMessages(registry.LMOWithFullName(tn))
		if len(msgs) != 1 {
			panic(fmt.Sprintf("message not found %s", tn))
		}
		field.message = msgs[0]
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		field.fieldType = registry.FieldTypeString
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		field.fieldType = registry.FieldTypeBool
	default:
		// TODO
		panic("unhandled type")
	}
	switch fieldDescriptorProto.GetLabel() {
	case descriptorpb.FieldDescriptorProto_LABEL_REPEATED:
		field.repeated = true
	default:
	}

	return field
}
