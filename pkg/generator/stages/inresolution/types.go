package inresolution

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/stages/opresolution"
)

type InputResolver interface {
	Resolve(api dsl.ApiDescriptor, opCtxs []opresolution.OperationContext) *InputContext
}
