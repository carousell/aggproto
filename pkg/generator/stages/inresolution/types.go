package inresolution

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/stages/msgresolution"
	"github.com/carousell/aggproto/pkg/generator/stages/opresolution"
)

type InputResolver interface {
	Resolve(api dsl.ApiDescriptor, meta dsl.Meta, input []dsl.ArgDescriptor, opCtxs []opresolution.OperationContext, context msgresolution.AdaptorContext) (*InputContext, error)
}
