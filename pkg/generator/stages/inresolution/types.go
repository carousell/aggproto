package inresolution

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/stages/opresolution"
)

type InputResolver interface {
	Resolve(api dsl.ApiDescriptor, meta dsl.Meta, opCtxs []opresolution.OperationContext) *InputContext
}
