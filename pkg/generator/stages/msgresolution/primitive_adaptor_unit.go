package msgresolution

import (
	"strings"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/iancoleman/strcase"
)

type staticPrimitiveAdaptorUnit struct {
	fieldName   string
	primitiveFD dsl.PrimitiveFieldDescriptor
}

func (s *staticPrimitiveAdaptorUnit) getRepeatedSizeString([]string) ([]string, error) {
	return nil, nil
}

func (s *staticPrimitiveAdaptorUnit) dependencies() [][]fieldMessageDependency {
	return nil
}

func (s *staticPrimitiveAdaptorUnit) printAsAdaptorCode(p printer.Printer, referenceName string, parents []string, repeatedStringIdxRef []string) error {
	fieldName := strcase.ToCamel(s.fieldName)
	switch fd := s.primitiveFD.(type) {
	case *dsl.StringValueFieldDescriptor:
		p.P(referenceName, ".", fieldName, " = \"", fd.Value, "\"")
	case *dsl.BoolValueFieldDescriptor:
		p.P(referenceName, ".", fieldName, " = ", fd.Value)
	case *dsl.FloatValueFieldDescriptor:
		p.P(referenceName, ".", fieldName, " = ", fd.Value)
	case *dsl.IntValueFieldDescriptor:
		p.P(referenceName, ".", fieldName, " = ", fd.Value)
	}

	return nil
}

func (s *staticPrimitiveAdaptorUnit) printAsProtoField(p printer.Printer, fieldIdx int) {
	switch s.primitiveFD.(type) {
	case *dsl.StringValueFieldDescriptor:
		p.P("string ", s.fieldName, " = ", fieldIdx, ";")
	case *dsl.BoolValueFieldDescriptor:
		p.P("bool ", s.fieldName, " = ", fieldIdx, ";")
	case *dsl.IntValueFieldDescriptor:
		p.P("int64 ", s.fieldName, " = ", fieldIdx, ";")
	case *dsl.FloatValueFieldDescriptor:
		p.P("float64 ", s.fieldName, " = ", fieldIdx, ";")
	}
}

func (s *staticPrimitiveAdaptorUnit) printProtoDefinitions(p printer.Printer) {
}

func (s *staticPrimitiveAdaptorUnit) isAdaptorUnit() {
}

func makeStaticPrimitiveAdaptorUnit(ofd *dsl.NamespacedMessageFieldDescriptor, ifd dsl.PrimitiveFieldDescriptor) adaptorUnit {
	splits := strings.Split(ofd.NamespacedField, ".")
	var unit adaptorUnit
	for i := len(splits) - 1; i >= 0; i -= 1 {
		if unit == nil {
			unit = &staticPrimitiveAdaptorUnit{fieldName: splits[i], primitiveFD: ifd}
		} else {
			unit = &nestedAdaptorUnit{fieldName: splits[i], nestedUnit: []adaptorUnit{unit}}
		}
	}
	return unit
}

func (s *staticPrimitiveAdaptorUnit) printProtoSnippet() {

}
