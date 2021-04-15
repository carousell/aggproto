package inresolution

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/generator/stage"
	"github.com/carousell/aggproto/pkg/registry"
	"github.com/iancoleman/strcase"
)

type InputContext struct {
	apiDescriptor dsl.ApiDescriptor
	required      []registry.Message
	argUnits      []argUnit
	meta          dsl.Meta
}

func (i *InputContext) AddStageDependency(stage stage.Stage) {
	panic("input stage cannot have a dependency")
}

func (i *InputContext) GetStageDependencies() []stage.Stage {
	return nil
}

func (i *InputContext) PrintProto(factory printer.Factory) {
	p := factory.Get(fmt.Sprintf("%s.proto", i.apiDescriptor.FileName()))
	p.P("message ", fmt.Sprintf("%sRequest", strcase.ToCamel(i.apiDescriptor.Name())), " {")
	p.Tab()
	for idx, au := range i.argUnits {
		au.printProtoDependency(p, idx)
	}
	for idx, au := range i.argUnits {
		au.printProtoUsage(p, idx)
	}
	p.UnTab()
	p.P("}")
}

func (i *InputContext) PrintCodeUsage(p printer.Printer) {
	var retReferences []string
	done := map[registry.Message]struct{}{}
	for _, au := range i.argUnits {
		for _, msg := range au.produces() {
			if _, ok := done[msg]; ok {
				continue
			}
			done[msg]= struct{}{}
			retReferences = append(retReferences, strcase.ToLowerCamel(msg.Name()))

		}
	}
	if len(retReferences) == 0 {
		return
	}
	p.P(strings.Join(retReferences, ", "), " := transform", strcase.ToCamel(i.apiDescriptor.Name()), "Request(req)")
}
func (i *InputContext) PrintCode(printerFactory printer.Factory) {
	p := printerFactory.Get(fmt.Sprintf("%s_input.go", i.apiDescriptor.FileName()))
	p.P("package ", i.apiDescriptor.Group(), "_v", i.apiDescriptor.Version())
	reqClassName := fmt.Sprintf("%sRequest", strcase.ToCamel(i.apiDescriptor.Name()))
	p.P()
	prepareImports(p, i.meta, i.argUnits)
	p.P()
	p.P("func ", "transform", reqClassName, "(req *", reqClassName, ") ", prepareRequiredReturns(i.argUnits), " {")
	p.Tab()
	var retReferences []string
	done := map[registry.Message]struct{}{}
	for _, au := range i.argUnits {
		retReferences = append(retReferences, au.prepareRequired(p, "req", done)...)
		//retReferences = append(retReferences, prepareRequired(p, r, fmt.Sprintf("req.%s", r.Name())))
	}
	p.P("return ", strings.Join(retReferences, ", "))
	p.UnTab()
	p.P("}")
}

func prepareImports(p printer.Printer, meta dsl.Meta, argUnits []argUnit) {
	packages := map[string]struct{}{}
	for _, au := range argUnits {
		for _, msg := range au.produces() {
			packages[msg.Package()] = struct{}{}
		}
	}
	if len(packages) > 0 {
		p.P("import (")
		p.Tab()
		for pkg, _ := range packages {
			p.P("\"", meta.GoPackage, "/", pkg, "\"")
		}
		p.UnTab()
		p.P(")")
	}
}

func prepareRequiredReturns(argUnits []argUnit) string {
	if len(argUnits) == 0 {
		return ""
	}
	var ret []string
	done := map[registry.Message]struct{}{}
	for _, au := range argUnits {
		for _, msg := range au.produces() {
			if _, ok := done[msg]; ok {
				continue
			}
			done[msg] = struct{}{}
			ret = append(ret, fmt.Sprintf("*%s.%s", msg.Package(), msg.Name()))
		}

	}
	buf := new(bytes.Buffer)
	if len(ret) > 1 {
		buf.WriteString("(")
	}
	buf.WriteString(strings.Join(ret, ", "))
	if len(ret) > 1 {
		buf.WriteString(")")
	}
	return buf.String()
}
