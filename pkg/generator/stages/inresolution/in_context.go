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

func (i *InputContext) PrintProto(p printer.Printer) {
	p.P("message ", fmt.Sprintf("%sRequest", strcase.ToCamel(i.apiDescriptor.Name())), "{")
	p.Tab()
	for idx, r := range i.required {
		p.P(r.Name(), " ", strcase.ToSnake(r.Name()), " = ", idx, ";")
	}
	p.UnTab()
	p.P("}")
}

func (i *InputContext) PrintCode(p printer.Printer) {
	panic("implement me")
}
