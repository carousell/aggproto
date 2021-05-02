package parser

import (
	"fmt"
	"strings"

	"github.com/carousell/aggproto/pkg/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type MessageField struct {
	FieldType        registry.FieldType
	FieldName        string
	FieldMessageType registry.Message
	FieldContext     registry.Message
	RepeatedField    bool
}

func (m *MessageField) Type() registry.FieldType {
	return m.FieldType
}

func (m *MessageField) Name() string {
	return m.FieldName
}

func (m *MessageField) Message() registry.Message {
	return m.FieldMessageType
}

func (m *MessageField) Context() registry.Message {
	return m.FieldContext
}

func (m *MessageField) Repeated() bool {
	return m.RepeatedField
}

func parseField(r registry.Registry, msgContext *MessageContainer, fieldDescriptorProto *descriptorpb.FieldDescriptorProto) *MessageField {
	if fieldDescriptorProto.OneofIndex != nil {
		// TODO
		panic("unhandled oneof")
	}
	field := &MessageField{
		FieldContext: msgContext,
		FieldName:    fieldDescriptorProto.GetName(),
	}
	switch fieldDescriptorProto.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		field.FieldType = registry.FieldTypeMessage
		tn := fieldDescriptorProto.GetTypeName()
		if strings.HasPrefix(tn, ".") {
			tn = strings.Replace(tn, ".", "", 1)
		}
		msgs := r.ListMessages(registry.LMOWithFullName(tn))
		if len(msgs) != 1 {
			panic(fmt.Sprintf("message not found %s", tn))
		}
		field.FieldMessageType = msgs[0]
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		field.FieldType = registry.FieldTypeString
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		field.FieldType = registry.FieldTypeBool
	case descriptorpb.FieldDescriptorProto_TYPE_INT64:
		field.FieldType = registry.FieldTypeInt64
	case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		field.FieldType = registry.FieldTypeEnum
	default:
		// TODO
		panic("unhandled type " + fieldDescriptorProto.GetType().String())
	}
	switch fieldDescriptorProto.GetLabel() {
	case descriptorpb.FieldDescriptorProto_LABEL_REPEATED:
		field.RepeatedField = true
	default:
	}

	return field
}
