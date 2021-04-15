package opresolution

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/generator/stage"
	"github.com/carousell/aggproto/pkg/generator/stages/msgresolution"
	"github.com/carousell/aggproto/pkg/registry"
)

type OperationResolver interface {
	Resolve(api dsl.ApiDescriptor, meta dsl.Meta, ctx msgresolution.AdaptorContext, descriptor dsl.OpDescriptor) []OperationContext
}

type OperationContext interface {
	stage.Stage
	InputDependency() registry.Message
	ClientReferenceName() string
	RequiredImport() string
	ClientDependency() string
	PrintConstructorUsage(p printer.Printer)
	Produces() registry.Message
	Consumes() registry.Message
}
