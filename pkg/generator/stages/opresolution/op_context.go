package opresolution

import (
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/registry"
)

type opContext struct {
	operation registry.UnaryOperation
}

func (o *opContext) PrintProto(p printer.Printer) {
}

func (o *opContext) PrintCode(p printer.Printer) {
	panic("implement me")
}

func (o *opContext) InputDependency() registry.Message {
	return o.operation.Input()
}
