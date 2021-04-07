package orchestrator

import (
	"fmt"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/generator/stage"
	"github.com/iancoleman/strcase"
)

type Orchestrator interface {
	stage.Stage
}
type orchestrator struct {
	api  dsl.ApiDescriptor
	meta dsl.Meta
}

func New(api dsl.ApiDescriptor, meta dsl.Meta) Orchestrator {
	return &orchestrator{api: api, meta: meta}
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
}

func (o *orchestrator) PrintCode(factory printer.Factory) {
	n := o.api.Name()
	higherN := strcase.ToCamel(n)
	lowerN := strcase.ToLowerCamel(n)
	p := factory.Get(fmt.Sprintf("%s_orchestrator.go", o.api.FileName()))
	p.P("package ", o.api.Group(), "_v", o.api.Version())
	p.P()
	p.P("import \"context\"")
	p.P()
	p.P("type ", lowerN, "Svc struct {")
	p.Tab()
	p.P("Unimplemented", higherN, "ServiceServer")
	p.UnTab()
	p.P("}")
	p.P()
	p.P("func New() ", higherN, "ServiceServer {")
	p.Tab()
	p.P("return &", lowerN, "Svc{}")
	p.UnTab()
	p.P("}")
	p.P()
	p.P("func (s *", lowerN, "Svc) Invoke", higherN, "(ctx context.Context, ", "req *", higherN, "Request) (*", higherN, "Response, error){")
	p.P("}")
}
