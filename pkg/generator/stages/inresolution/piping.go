package inresolution

import (
	"fmt"
	"strings"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/generator/stage"
	"github.com/carousell/aggproto/pkg/generator/stages/opresolution"
	"github.com/carousell/aggproto/pkg/registry"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
)

type pipedFieldUnit struct {
	depStack      []fieldMessageDependency
	producerStack []fieldMessageDependency
}

func (u *pipedFieldUnit) printCode(p printer.Printer, done map[registry.Message]struct{}) string {
	var ret string
	for _, fmd := range u.producerStack {
		if ret == "" {
			ret = strcase.ToLowerCamel(fmd.msg.Name())
		}
		if fmd.msg == nil {
			continue
		}
		if _, ok := done[fmd.msg]; !ok {
			p.P(strcase.ToLowerCamel(fmd.msg.Name()), " := &", fmd.msg.Package(), ".", strcase.ToCamel(fmd.msg.Name()), "{}")
			done[fmd.msg] = struct{}{}
		}
	}
	lhs := ret
	for _, fmd := range u.producerStack {
		lhs = fmt.Sprintf("%s.%s", lhs, strcase.ToCamel(fmd.fieldName))
	}
	rhs := strcase.ToLowerCamel(u.depStack[0].msg.Name())
	for idx, fmd := range u.depStack {
		if idx == 0 {
			continue
		}
		rhs = fmt.Sprintf("%s.%s", rhs, strcase.ToCamel(fmd.fieldName))
	}
	p.P(lhs, " = ", rhs)
	return ret

}

type pipedContext struct {
	units       []*pipedFieldUnit
	stages      []stage.Stage
	producesMsg registry.Message
	api         dsl.ApiDescriptor
	meta        dsl.Meta
}

func (pc *pipedContext) PrintProto(printer.Factory) {
}

func (pc *pipedContext) PrintCode(printerFactory printer.Factory) {
	p := printerFactory.Get(fmt.Sprintf("%s_transformer.go", pc.api.FileName()))
	p.P("package ", pc.api.Group(), "_v", pc.api.Version())
	preparePipedImports(p, pc.meta, pc.producesMsg, pc.dependencies())
	var dep []string
	for _, d := range pc.dependencies() {
		dep = append(dep, fmt.Sprintf("%s *%s.%s", strcase.ToLowerCamel(d.Name()), d.Package(), strcase.ToCamel(d.Name())))
	}

	p.P("func transformTo", strcase.ToCamel(pc.producesMsg.Name()), "(", strings.Join(dep, ", "), ") *", pc.producesMsg.Package(), ".", strcase.ToCamel(pc.producesMsg.Name()), "{")
	p.Tab()
	done := map[registry.Message]struct{}{}
	var ret []string
	for _, pfu := range pc.units {
		ret = append(ret, pfu.printCode(p, done))
	}
	p.P("return ", strings.Join(ret, ", "))
	p.UnTab()
	p.P("}")

}

func preparePipedImports(p printer.Printer, meta dsl.Meta, msg registry.Message, dependencies []registry.Message) {
	p.P("import (")
	p.Tab()
	for _, msg := range append(dependencies, msg) {
		p.P("\"", meta.GoPackage, "/", msg.Package(), "\"")
	}
	p.UnTab()
	p.P(")")
}

func (pc *pipedContext) PrintCodeUsage(p printer.Printer) {
	var dep []string
	for _, d := range pc.dependencies() {
		dep = append(dep, fmt.Sprintf("%s", strcase.ToLowerCamel(d.Name())))
	}
	p.P(strcase.ToLowerCamel(pc.producesMsg.Name()), " := transformTo", strcase.ToCamel(pc.producesMsg.Name()), "(", strings.Join(dep, ", "), ")")
}

func (pc *pipedContext) AddStageDependency(stage stage.Stage) {
	pc.stages = append(pc.stages, stage)
}

func (pc *pipedContext) GetStageDependencies() []stage.Stage {
	return pc.stages
}
func (pc *pipedContext) dependencies() []registry.Message {
	done := map[registry.Message]struct{}{}
	var ret []registry.Message
	for _, pfu := range pc.units {
		if _, ok := done[pfu.depStack[0].msg]; ok {
			continue
		}
		done[pfu.depStack[0].msg] = struct{}{}
		ret = append(ret, pfu.depStack[0].msg)
	}
	return ret
}
func (pc *pipedContext) produces() registry.Message {
	return pc.producesMsg
}

func addPipedDependencies(pipedContexts []*pipedContext, opCtxs []opresolution.OperationContext) error {
	for _, pc := range pipedContexts {
		for _, d := range pc.dependencies() {
			found := false
			for _, op := range opCtxs {
				if d == op.Produces() {
					found = true
					pc.AddStageDependency(op)
					break
				}
			}
			if !found {
				return errors.Errorf("no operation produces dependent message %s", d.FullName())
			}
		}
		producesMsg := pc.produces()
		found := false
		for _, op := range opCtxs {
			if producesMsg == op.Consumes() {
				found = true
				op.AddStageDependency(pc)
			}
		}
		if !found {
			return errors.Errorf("no operation consumes produced message %s", producesMsg.FullName())
		}

	}
	return nil
}

func groupPipedUnitsByProduces(api dsl.ApiDescriptor, meta dsl.Meta, units []*pipedFieldUnit) []*pipedContext {
	producerMap := map[registry.Message][]*pipedFieldUnit{}
	for _, u := range units {
		if _, ok := producerMap[u.producerStack[0].msg]; !ok {
			producerMap[u.producerStack[0].msg] = nil
		}
		producerMap[u.producerStack[0].msg] = append(producerMap[u.producerStack[0].msg], u)
	}
	var ret []*pipedContext
	for msg, v := range producerMap {
		ret = append(ret, &pipedContext{units: v, producesMsg: msg, api: api, meta: meta})
	}
	return ret
}

func applyPiping(pipeArg *dsl.PipedArgDescriptor, units []argUnit, r registry.Registry) (*pipedFieldUnit, error) {
	rawArgPath := pipeArg.Output()
	for _, au := range units {
		argPaths := strings.Split(rawArgPath, ".")
		if ok, fau := matchAndRemoveArgPath(au, argPaths, 0); ok {
			msgs := r.ListMessages(registry.LMOWithPrefixMatch(pipeArg.Input()))
			if len(msgs) != 1 {
				return nil, errors.Errorf("piped return must match exactly one message found %d for %s", len(msgs), pipeArg.Input())
			}
			depStack, err := resolveField(msgs[0], pipeArg.Input())
			if err != nil {
				return nil, err
			}
			return &pipedFieldUnit{producerStack: fau.producerStack, depStack: depStack}, nil
		}
	}

	return nil, errors.Errorf("piped arg did not match any required input %s ", rawArgPath)
}

func resolveField(message registry.Message, input string) ([]fieldMessageDependency, error) {
	splits := strings.Split(input, ".")
	resolvedMsg := message
	dependencyStack := []fieldMessageDependency{{msg: message, fieldName: message.Name()}}
	var resolvedField registry.Field
	for idx, fd := range splits {
		found := false
		if idx == 0 && fd == resolvedMsg.Package() {
			continue
		}
		if idx <= 1 && fd == resolvedMsg.Name() {
			continue
		}
		for _, field := range resolvedMsg.Fields() {
			if field.Name() == fd {
				found = true
				if field.Type() != registry.FieldTypeMessage {
					if idx != len(splits)-1 {
						return nil, errors.Errorf("partially resolved to field %s", input)
					}
					resolvedField = field
					dependencyStack = append(dependencyStack, fieldMessageDependency{fieldName: field.Name()})
				} else {
					resolvedMsg = field.Message()
					dependencyStack = append(dependencyStack, fieldMessageDependency{fieldName: field.Name(), msg: resolvedMsg})
				}
				break
			}
		}
		if !found {
			return nil, errors.Errorf("Did not resolve to a known field %s", input)
		}
	}
	if resolvedField == nil {
		return nil, errors.Errorf("Did not resolve to a field %s", input)
	}
	return dependencyStack, nil
}
