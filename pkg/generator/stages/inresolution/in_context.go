package inresolution

import (
	"fmt"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/registry"
	"github.com/iancoleman/strcase"
)

type InputContext struct {
	apiDescriptor dsl.ApiDescriptor
	required      []registry.Message
}

func printMessage(p printer.Printer, m registry.Message) {
	p.P("message ", fmt.Sprintf("%sGen {", m.Name()))
	p.Tab()
	for _, subDef := range m.Definitions() {
		printMessage(p, subDef)
	}
	for idx, f := range m.Fields() {
		switch f.Type() {
		case registry.FieldTypeString:
			p.P("string ", f.Name(), " = ", idx+1, ";")
		case registry.FieldTypeInt64:
			p.P("int64 ", f.Name(), " = ", idx+1, ";")
		case registry.FieldTypeDouble:
			p.P("float64 ", f.Name(), " = ", idx+1, ";")
		case registry.FieldTypeBool:
			p.P("bool ", f.Name(), " = ", idx+1, ";")
		case registry.FieldTypeMessage:
			p.P(strcase.ToCamel(f.Name()), "Gen ", strcase.ToSnake(f.Name()), " = ", idx+1, ";")
		default:
			panic("unsupported type " + f.Name())
		}
	}
	p.UnTab()
	p.P("}")
}

func (i *InputContext) PrintProto(factory printer.Factory) {
	p := factory.Get(fmt.Sprintf("%s.proto", i.apiDescriptor.FileName()))
	p.P("message ", fmt.Sprintf("%sRequest", strcase.ToCamel(i.apiDescriptor.Name())), " {")
	p.Tab()
	for _, r := range i.required {
		printMessage(p, r)
	}
	for idx, r := range i.required {
		p.P(r.Name(), "Gen ", strcase.ToSnake(r.Name()), " = ", idx+1, ";")
	}
	p.UnTab()
	p.P("}")
}

func (i *InputContext) PrintCode(p printer.Factory) {
}
