package msgresolution

import (
	"strings"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
)

type staticPrimitiveAdaptorUnit struct {
	fieldName   string
	primitiveFD dsl.PrimitiveFieldDescriptor
}

func (s *staticPrimitiveAdaptorUnit) printAsAdaptorCode(p printer.Printer, referenceName string) {
	switch fd := s.primitiveFD.(type) {
	case *dsl.StringValueFieldDescriptor:
		p.P(referenceName, ".", s.fieldName, " = \"", fd.Value, "\"")
	case *dsl.BoolValueFieldDescriptor:
		p.P(referenceName, ".", s.fieldName, " = ", fd.Value)
	case *dsl.FloatValueFieldDescriptor:
		p.P(referenceName, ".", s.fieldName, " = ", fd.Value)
	case *dsl.IntValueFieldDescriptor:
		p.P(referenceName, ".", s.fieldName, " = ", fd.Value)
	}

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

func (s *staticPrimitiveAdaptorUnit) printProtoDefinitions(p printer.Printer, fieldIdx int) {
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
