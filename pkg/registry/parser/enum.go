package parser

import (
	"fmt"
	"github.com/carousell/aggproto/pkg/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type EnumValueContainer struct {
	FieldName string
	Value     int32
}

type EnumContainer struct {
	EnumName    string
	PackageName string
	EnumParent  *MessageContainer
	EnumFields  []*EnumValueContainer
}

func (e *EnumContainer) FullName() string {
	return fmt.Sprintf("%s.%s", e.PackageName, e.EnumName)
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
