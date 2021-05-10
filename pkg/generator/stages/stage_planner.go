//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
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
	Plan(ctx *dsl.Context) (stage.Stage, error)
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

func (s *stagePlanner) Plan(ctx *dsl.Context) (stage.Stage, error) {
	o := orchestrator.New(ctx.Api, ctx.Meta)
	adaptorContext, er := s.msgResolver.Resolve(ctx.Api, ctx.Meta, ctx.Output)
	if er != nil {
		return nil, er
	}
	operationContexts := s.opResolver.Resolve(ctx.Api, ctx.Meta, adaptorContext, ctx.Operation)
	_, er = s.inputResolver.Resolve(ctx.Api, ctx.Meta, ctx.Input, operationContexts, adaptorContext)
	if er != nil {
		return nil, er
	}
	o.AddStages(resolveStageOrder(adaptorContext))
	return o, nil
}

func resolveStageOrder(finalStage stage.Stage) []stage.Stage {
	var ret []stage.Stage
	depChan := make(chan stage.Stage)
	doneChan := make(chan struct{})
	go func() {
		for nextStage := range depChan {
			ret = append(ret, nextStage)
		}
		doneChan <- struct{}{}
	}()
	doneStages := map[stage.Stage]struct{}{}
	traversePostOrder(finalStage, depChan, doneStages)
	close(depChan)
	<-doneChan
	return ret
}

func traversePostOrder(currentStage stage.Stage, depChan chan stage.Stage, doneStages map[stage.Stage]struct{}) {
	for _, s := range currentStage.GetStageDependencies() {
		if _, ok := doneStages[s]; ok {
			continue
		}
		doneStages[s] = struct{}{}
		traversePostOrder(s, depChan, doneStages)
	}
	depChan <- currentStage
}
