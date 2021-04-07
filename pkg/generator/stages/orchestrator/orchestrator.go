package orchestrator

import (
	"fmt"
	"strings"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/generator/stage"
	"github.com/carousell/aggproto/pkg/generator/stages/opresolution"
	"github.com/iancoleman/strcase"
)

type Orchestrator interface {
	stage.Stage
	AddStages(stages []stage.Stage)
}
type orchestrator struct {
	api    dsl.ApiDescriptor
	meta   dsl.Meta
	stages []stage.Stage
}

func (o *orchestrator) AddStageDependency(stage stage.Stage) {
	panic("should never be called")
}

func (o *orchestrator) GetStageDependencies() []stage.Stage {
	panic("should never be called")
}

func (o *orchestrator) PrintCodeUsage(p printer.Printer) {
	panic("should never be called")
}

func New(api dsl.ApiDescriptor, meta dsl.Meta) Orchestrator {
	return &orchestrator{api: api, meta: meta}
}

func (o *orchestrator) AddStages(stages []stage.Stage) {
	o.stages = stages
}

func (o *orchestrator) PrintProto(factory printer.Factory) {
	p := factory.Get(fmt.Sprintf("%s.proto", o.api.FileName()))
	p.P("syntax = \"proto3\";")
	p.P("package ", o.api.Group(), "_v", o.api.Version(), ";")
	p.P("option go_package = \"", o.meta.GoPackage, "/", o.api.Group(), "_v", o.api.Version(), "\";")
	p.P("service ", strcase.ToCamel(o.api.Name()), "Service {")
	p.Tab()
	p.P("rpc Invoke", strcase.ToCamel(o.api.Name()), " (", strcase.ToCamel(o.api.Name()), "Request) returns (", strcase.ToCamel(o.api.Name()), "Response);")
	p.UnTab()
	p.P("}")
	for _, s := range o.stages {
		s.PrintProto(factory)
	}
}

func (o *orchestrator) PrintCode(factory printer.Factory) {
	n := o.api.Name()
	higherN := strcase.ToCamel(n)
	lowerN := strcase.ToLowerCamel(n)
	p := factory.Get(fmt.Sprintf("%s_orchestrator.go", o.api.FileName()))
	p.P("package ", o.api.Group(), "_v", o.api.Version())
	p.P()
	printImports(p, o.stages)
	p.P()
	p.P("type ", lowerN, "Svc struct {")
	p.Tab()
	p.P("Unimplemented", higherN, "ServiceServer")
	for _, s := range o.stages {
		if opCtx, ok := s.(opresolution.OperationContext); ok {
			p.P(opCtx.ClientReferenceName())
		}
	}
	p.UnTab()
	p.P("}")
	p.P()
	deps := getClientDependencies(o.stages)
	p.P("func New(", strings.Join(deps, ", "), ") ", higherN, "ServiceServer {")
	p.Tab()
	p.P("return &", lowerN, "Svc{")
	p.Tab()
	for _, s := range o.stages {
		if opCtx, ok := s.(opresolution.OperationContext); ok {
			opCtx.PrintConstructorUsage(p)
		}
	}
	p.UnTab()
	p.P("}")
	p.UnTab()
	p.P("}")
	p.P()
	p.P("func (s *", lowerN, "Svc) Invoke", higherN, "(ctx context.Context, ", "req *", higherN, "Request) (*", higherN, "Response, error){")
	p.Tab()
	for _, s := range o.stages {
		s.PrintCodeUsage(p)
	}
	p.P("return resp, nil")
	p.UnTab()
	p.P("}")
	for _, s := range o.stages {
		s.PrintCode(factory)
	}
}

func getClientDependencies(stages []stage.Stage) []string {
	deps := map[string]struct{}{}
	for _, s := range stages {
		if opCtx, ok := s.(opresolution.OperationContext); ok {
			deps[opCtx.ClientDependency()] = struct{}{}
		}
	}
	var ret []string
	for dep, _ := range deps {
		ret = append(ret, dep)
	}
	return ret
}

func printImports(p printer.Printer, stages []stage.Stage) {
	uniqueImports := map[string]struct{}{}
	for _, s := range stages {
		if opCtx, ok := s.(opresolution.OperationContext); ok {
			uniqueImports[opCtx.RequiredImport()] = struct{}{}
		}
	}
	p.P("import (")
	p.Tab()
	p.P("\"context\"")
	p.P()
	for imp, _ := range uniqueImports {
		p.P("\"", imp, "\"")
	}
	p.UnTab()
	p.P(")")

}
