package inresolution

import "github.com/carousell/aggproto/pkg/generator/stages/opresolution"

type InputResolver interface {
	Resolve(opCtxs []opresolution.OperationContext) *InputContext
}
