package opresolution

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/stage"
	"github.com/carousell/aggproto/pkg/generator/stages/msgresolution"
	"github.com/carousell/aggproto/pkg/registry"
)

type OperationResolver interface {
	Resolve(ctx msgresolution.AdaptorContext, descriptor dsl.OpDescriptor) []OperationContext
}

type OperationContext interface {
	stage.Stage
	InputDependency() registry.Message
}
