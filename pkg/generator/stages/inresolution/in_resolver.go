package inresolution

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/stages/opresolution"
	"github.com/carousell/aggproto/pkg/registry"
)

type inResolver struct {
	r registry.Registry
}

func New(r registry.Registry) InputResolver {
	return &inResolver{r: r}
}

func (i *inResolver) Resolve(api dsl.ApiDescriptor, meta dsl.Meta, opCtxs []opresolution.OperationContext) *InputContext {
	requiredInputsMap := map[string]registry.Message{}
	for _, op := range opCtxs {
		msg := op.InputDependency()
		if _, ok := requiredInputsMap[msg.FullName()]; !ok {
			requiredInputsMap[msg.FullName()] = msg
		}
	}
	var requiredInput []registry.Message
	for _, v := range requiredInputsMap {
		requiredInput = append(requiredInput, v)
	}
	inCtx := &InputContext{apiDescriptor: api, required: requiredInput, meta: meta}
	for _, opCtx := range opCtxs {
		opCtx.AddStageDependency(inCtx)
	}
	return inCtx
}
