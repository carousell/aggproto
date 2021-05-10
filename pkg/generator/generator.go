//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package generator

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/generator/stages"
	"github.com/carousell/aggproto/pkg/registry"
)

type generator struct {
	r registry.Registry
}

func Generator(r registry.Registry) *generator {
	return &generator{r: r}
}

func (g *generator) PlanAndGenerate(ctx *dsl.Context) (map[string]string, error) {
	planner := stages.Planner(g.r)
	plan, er := planner.Plan(ctx)
	if er != nil {
		return nil, er
	}
	printerFactory := printer.New()
	plan.PrintProto(printerFactory)
	er = plan.PrintCode(printerFactory)
	if er != nil {
		return nil, er
	}
	return printerFactory.Out(), nil
}
