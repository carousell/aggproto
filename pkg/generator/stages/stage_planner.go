package stages

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/stage"
	"github.com/carousell/aggproto/pkg/generator/stages/inresolution"
	"github.com/carousell/aggproto/pkg/generator/stages/msgresolution"
	"github.com/carousell/aggproto/pkg/generator/stages/opresolution"
	"github.com/carousell/aggproto/pkg/generator/stages/orchestrator"
	"github.com/carousell/aggproto/pkg/registry"
)

type StagePlanner interface {
	Plan(ctx *dsl.Context) stage.Stage
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

func (s *stagePlanner) Plan(ctx *dsl.Context) stage.Stage {
	o := orchestrator.New(ctx.Api, ctx.Meta)
	adaptorContext := s.msgResolver.Resolve(ctx.Api, ctx.Meta, ctx.Output)
	operationContexts := s.opResolver.Resolve(ctx.Api, ctx.Meta, adaptorContext, ctx.Operation)
	_ = s.inputResolver.Resolve(ctx.Api, ctx.Meta, operationContexts)
	o.AddStages(resolveStageOrder(adaptorContext))
	return o
}

func resolveStageOrder(finalStage stage.Stage) []stage.Stage {
	var ret []stage.Stage
	depChan := make(chan stage.Stage)
	go func() {
		for nextStage := range depChan {
			ret = append(ret, nextStage)
		}
	}()
	traversePostOrder(finalStage, depChan)
	close(depChan)
	return ret
}

func traversePostOrder(currentStage stage.Stage, depChan chan stage.Stage) {
	for _, s := range currentStage.GetStageDependencies() {
		traversePostOrder(s, depChan)
	}
	depChan <- currentStage
}
