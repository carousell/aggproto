package stages

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/stage"
	"github.com/carousell/aggproto/pkg/generator/stages/inresolution"
	"github.com/carousell/aggproto/pkg/generator/stages/msgresolution"
	"github.com/carousell/aggproto/pkg/generator/stages/opresolution"
	"github.com/carousell/aggproto/pkg/registry"
)

type StagePlanner interface {
	Plan(ctx dsl.Context) []stage.Stage
}

func Planner(r registry.Registry) StagePlanner {
	return &stagePlanner{
		r:             r,
		msgResolver:   msgresolution.New(r),
		opResolver:    opresolution.New(r),
		inputResolver: inresolution.New(r),
	}
}

type stagePlanner struct {
	r             registry.Registry
	msgResolver   msgresolution.MessageResolver
	opResolver    opresolution.OperationResolver
	inputResolver inresolution.InputResolver
}

func (s *stagePlanner) Plan(ctx dsl.Context) []stage.Stage {
	var res []stage.Stage
	adaptorContext := s.msgResolver.Resolve(ctx.Api, ctx.Output)
	res = append(res, adaptorContext)

	operationContexts := s.opResolver.Resolve(adaptorContext, ctx.Operation)

	for _, oc := range operationContexts {
		res = append(res, oc)
	}

	inputContext := s.inputResolver.Resolve(operationContexts)
	res = append(res, inputContext)
	return res
}
