package opresolution

import (
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/registry"
)

type opContext struct {
	operation registry.UnaryOperation
}

func (o *opContext) PrintProto(p printer.Factory) {
}

func (o *opContext) PrintCode(p printer.Factory) {
}

func (o *opContext) InputDependency() registry.Message {
	return o.operation.Input()
}
