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

func (g *generator) PlanAndGenerate(ctx dsl.Context) {
	planner := stages.Planner(g.r)
	plan := planner.Plan(ctx)

	p := printer.New()
	for _, genCtx := range plan {
		genCtx.PrintProto(p)
	}


}
