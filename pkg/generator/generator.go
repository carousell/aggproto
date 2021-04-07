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

func (g *generator) PlanAndGenerate(ctx *dsl.Context) map[string]string {
	planner := stages.Planner(g.r)
	plan := planner.Plan(ctx)

	printerFactory := printer.New()
	for _, genCtx := range plan {
		genCtx.PrintProto(printerFactory)
		genCtx.PrintCode(printerFactory)
	}
	return printerFactory.Out()

}
