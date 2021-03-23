package msgresolution

import (
	"fmt"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/registry"
	"github.com/iancoleman/strcase"
)

type adaptorContext struct {
	apiDescriptor dsl.ApiDescriptor
	adaptorUnits  []adaptorUnit
}

func (a *adaptorContext) PrintCode(p printer.Printer) {
	p.P("package ", a.apiDescriptor.Group(), "_v", a.apiDescriptor.Version())
	respClassName := fmt.Sprintf("%sResponse", strcase.ToCamel(a.apiDescriptor.Name()))
	p.P("func ", "Adapt", respClassName, "()*", respClassName, "{")
	p.Tab()
	p.P("resp := &", respClassName, "{}")
	for _, au := range a.adaptorUnits {
		au.printAsAdaptorCode(p, "resp")
	}
	p.P("return resp")
	p.UnTab()
	p.P("}")
}

func (a *adaptorContext) PrintProto(p printer.Printer) {
	p.P("message ", fmt.Sprintf("%sResponse", strcase.ToCamel(a.apiDescriptor.Name())), "{")
	p.Tab()
	for idx, au := range a.adaptorUnits {
		au.printProtoDefinitions(p, idx+1)
	}
	for idx, au := range a.adaptorUnits {
		au.printAsProtoField(p, idx+1)
	}
	p.UnTab()
	p.P("}")
}

func (a *adaptorContext) Dependencies() []registry.Message {
	return nil
}
