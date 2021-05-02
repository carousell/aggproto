package parser

import (
	"github.com/carousell/aggproto/pkg/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type EnumContainer struct {
	EnumName    string
	PackageName string
	EnumParent  *MessageContainer
	EnumFields  []*EnumField
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

	var fields []*EnumField
	for _, dp := range enumType.GetValue() {
		enumF := &EnumField{
			FieldName: dp.GetName(),
			Value:     dp.GetNumber(),
		}

		fields = append(fields, enumF)
	}
	enum.EnumFields = fields
	return enum
}
