package msgresolution

import (
	"fmt"
	"strings"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/registry"
	"github.com/iancoleman/strcase"
)

type adaptorContext struct {
	apiDescriptor dsl.ApiDescriptor
	adaptorUnits  []adaptorUnit
}

func topLevelDeps(deps [][]fieldMessageDependency) []fieldMessageDependency {
	tld := map[string]fieldMessageDependency{}
	for _, dep := range deps {
		if len(dep) > 0 {
			if _, ok := tld[dep[0].fieldName]; !ok {
				tld[dep[0].fieldName] = dep[0]
			}
		}
	}
	var ret []fieldMessageDependency
	for _, v := range tld {
		ret = append(ret, v)
	}
	return ret
}
func printTopLevelDependencies(deps [][]fieldMessageDependency) string {
	tld := topLevelDeps(deps)
	var params []string
	for _, v := range tld {
		params = append(params, fmt.Sprintf("%s *%s.%s", strcase.ToLowerCamel(v.fieldName), v.message.Package(), strcase.ToCamel(v.message.Name())))
	}
	return strings.Join(params, ", ")
}

func prepareDependencies(p printer.Printer, deps [][]fieldMessageDependency) {
	doneDeps := map[string]struct{}{}
	for _, dep := range deps {
		for i := 0; i < len(dep)-1; i += 1 {
			key := fmt.Sprintf("%s_%s", dep[i+1].fieldName, dep[i].fieldName)
			if _, ok := doneDeps[key]; ok {
				continue
			}
			p.P(strcase.ToLowerCamel(dep[i+1].fieldName), " := ", strcase.ToLowerCamel(dep[i].fieldName), ".", strcase.ToLowerCamel(dep[i+1].fieldName))
			doneDeps[key] = struct{}{}
		}
	}
}

func (a *adaptorContext) PrintCode(printerFactory printer.Factory) {
	p := printerFactory.Get(fmt.Sprintf("%s.go", a.apiDescriptor.FileName()))
	p.P("package ", a.apiDescriptor.Group(), "_v", a.apiDescriptor.Version())
	respClassName := fmt.Sprintf("%sResponse", strcase.ToCamel(a.apiDescriptor.Name()))
	var deps [][]fieldMessageDependency
	for _, au := range a.adaptorUnits {
		deps = append(deps, au.dependencies()...)
	}
	p.P("func ", "Adapt", respClassName, "(", printTopLevelDependencies(deps), ") *", respClassName, "{")
	p.Tab()
	prepareDependencies(p, deps)
	p.P("resp := &", respClassName, "{}")
	for _, au := range a.adaptorUnits {
		au.printAsAdaptorCode(p, "resp")
	}
	p.P("return resp")
	p.UnTab()
	p.P("}")
}

func (a *adaptorContext) PrintProto(printerFactory printer.Factory) {
	p := printerFactory.Get(fmt.Sprintf("%s.proto", a.apiDescriptor.FileName()))
	p.P("message ", fmt.Sprintf("%sResponse", strcase.ToCamel(a.apiDescriptor.Name())), " {")
	p.Tab()
	for _, au := range a.adaptorUnits {
		au.printProtoDefinitions(p)
	}
	for idx, au := range a.adaptorUnits {
		au.printAsProtoField(p, idx+1)
	}
	p.UnTab()
	p.P("}")
}

func (a *adaptorContext) Dependencies() []registry.Message {
	var deps [][]fieldMessageDependency
	for _, au := range a.adaptorUnits {
		deps = append(deps, au.dependencies()...)
	}
	tld := topLevelDeps(deps)
	var ret []registry.Message
	for _, fmd := range tld {
		ret = append(ret, fmd.message)
	}
	return ret
}
