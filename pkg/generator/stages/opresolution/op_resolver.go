package opresolution

import (
	"log"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/stages/msgresolution"
	"github.com/carousell/aggproto/pkg/registry"
)

type opResolver struct {
	r registry.Registry
}

func New(r registry.Registry) OperationResolver {
	return &opResolver{r: r}
}

func (o *opResolver) Resolve(api dsl.ApiDescriptor, meta dsl.Meta, adaptorContext msgresolution.AdaptorContext, descriptor dsl.OpDescriptor) []OperationContext {
	allowedOpsMap := map[string]registry.UnaryOperation{}
	if len(descriptor.AllowedOperations) > 0 {
		allowedOps := o.r.ListOperations(registry.LSOWithOperationNameIn(descriptor.AllowedOperations))
		if len(allowedOps) != len(descriptor.AllowedOperations) {
			log.Fatalf("all operations not found required %d found %d", len(descriptor.AllowedOperations), len(allowedOps))
		}
		for _, op := range allowedOps {
			allowedOpsMap[op.FullName()] = op
		}
	}

	var resolvedOps []registry.UnaryOperation
	dependencies := adaptorContext.Dependencies()
depLoop:
	for _, dep := range dependencies {
		operations := o.r.ListOperations(registry.LSOWithOutputMessage(dep))
		if len(operations) == 0 {
			log.Fatalf("no operation found for message %s", dep.FullName())
		}
		if len(operations) > 1 {
			if len(allowedOpsMap) == 0 {
				log.Fatalf("multiple operations found for out %s specify allowedOperations", dep.FullName())
			}

			for _, op := range operations {
				if _, ok := allowedOpsMap[op.FullName()]; ok {
					resolvedOps = append(resolvedOps, op)
					continue depLoop
				}
			}
			log.Fatalf("Could not identify intersection between allowed ops and resolved ops")
		} else {
			resolvedOps = append(resolvedOps, operations[0])
		}
	}
	var ret []OperationContext
	for _, op := range resolvedOps {
		opCtx := &opContext{operation: op, api: api, meta: meta}
		ret = append(ret, opCtx)
		adaptorContext.AddStageDependency(opCtx)
	}
	return ret
}
