package parser

import (
	"github.com/carousell/aggproto/pkg/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type EnumValueContainer struct {
	FieldName string
	Value     int32
}

func (r *EnumValueContainer) Type() registry.FieldType {
	return registry.FieldTypeEnum
}

func (r *EnumValueContainer) Name() string {
	return r.FieldName
}

func (r *EnumValueContainer) Message() registry.Message {
	return nil
}

func (r *EnumValueContainer) Context() registry.Message {
	return nil
}

func (r *EnumValueContainer) Repeated() bool {
	return false
}

type EnumContainer struct {
	EnumName    string
	PackageName string
	EnumParent  *MessageContainer
	EnumFields  []*EnumValueContainer
}

func (e *EnumContainer) Name() string {
	return e.EnumName
}

func parseEnum(r registry.Registry, packageName string, enumType *descriptorpb.EnumDescriptorProto) *EnumContainer {
	if enumType == nil {
		return nil
	}
	enum := &EnumContainer{
		PackageName: packageName,
		EnumName:    enumType.GetName(),
	}

	var fields []*EnumValueContainer
	for _, dp := range enumType.GetValue() {
		enumF := &EnumValueContainer{
			FieldName: dp.GetName(),
			Value:     dp.GetNumber(),
		}

		fields = append(fields, enumF)
	}
	enum.EnumFields = fields
	return enum
}
