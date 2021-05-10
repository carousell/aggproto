//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
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

func pipedGetNextDependencyRepeatedSizeString(depStack []fieldMessageDependency, depRef []string) (int, string, error, []string) {
	for idx, fmd := range depStack {
		if fmd.repeated {
			return idx, fmt.Sprintf("len(%s.%s)", strings.Join(depRef, "."), strcase.ToCamel(fmd.fieldName)), nil, append(depRef, strcase.ToCamel(fmd.fieldName))
		} else {
			if len(depRef) == 0 {
				depRef = append(depRef, strcase.ToLowerCamel(fmd.fieldName))
			} else {
				depRef = append(depRef, strcase.ToCamel(fmd.fieldName))
			}
		}
	}
	return -1, "", errors.Errorf("no repeated found in dep stack"), depRef
}

func pipedCodePrinter(p printer.Printer, producerStack []fieldMessageDependency, depStack []fieldMessageDependency, done map[string]struct{}, ref []string, depRef []string) error {
	// initialize messages and slices
	for idx, fmd := range producerStack {
		if fmd.repeated {
			depRepIdx, depRepLen, er, depRef := pipedGetNextDependencyRepeatedSizeString(depStack, depRef)
			if er != nil {
				return er
			}
			ref = append(ref, strcase.ToCamel(fmd.fieldName))
			if fmd.msg != nil {
				p.P(strings.Join(ref, "."), " = make([]*", fmd.msg.Package(), ".", strcase.ToCamel(fmd.msg.Name()), ", ", depRepLen, ")")
				p.P("for i:=0; i < ", depRepLen, "; i+=1 {")
				p.Tab()
				p.P(strings.Join(ref, "."), "[i] = &", fmd.msg.Package(), ".", strcase.ToCamel(fmd.msg.Name()), "{}")
				p.UnTab()
				p.P("}")
				return pipedCodePrinter(p, producerStack[idx+1:], depStack[depRepIdx+1:], done, ref, depRef)
			} else {
				p.P(strings.Join(ref, "."), " = make([]", fmd.fieldType.GoTypeString(), ", ", depRepLen, ")")
				p.P("for i:=0; i < ", depRepLen, "; i+=1 {")
				p.Tab()
				var remDepRepRef []string
				for i := depRepIdx + 1; i < len(depStack); i += 1 {
					remDepRepRef = append(remDepRepRef, strcase.ToCamel(depStack[i].fieldName))
				}
				if len(remDepRepRef) > 0 {
					p.P(strings.Join(ref, "."), "[i] = ", strings.Join(depRef, "."), "[i].", strings.Join(remDepRepRef, "."))
				} else {
					p.P(strings.Join(ref, "."), "[i] = ", strings.Join(depRef, "."), "[i]")
				}
				p.UnTab()
				p.P("}")
				return nil
			}
		} else if fmd.msg != nil {
			p.P(strings.Join(ref, "."), ".", strcase.ToCamel(fmd.fieldName), " = &", fmd.msg.Package(), ".", strcase.ToCamel(fmd.msg.Name()), ")")
		}
	}
	return nil
}
func (u *pipedFieldUnit) printCode(p printer.Printer, done map[string]struct{}, own bool) (string, error) {
	retFmd := u.producerStack[0]
	ret := strcase.ToLowerCamel(retFmd.msg.Name())
	if _, ok := done[ret]; !ok && !own {
		p.P(ret, " := &", retFmd.msg.Package(), ".", strcase.ToCamel(retFmd.msg.Name()), "{}")
	}
	err := pipedCodePrinter(p, u.producerStack[1:], u.depStack, done, []string{ret}, nil)
	if err != nil {
		return "", err
	}
	if own {
		return "", nil
	}
	return ret, nil
	//for idx, fmd := range u.producerStack {
	//	if ret == "" {
	//		ret = strcase.ToLowerCamel(fmd.msg.Name())
	//		ref = append(ref, ret)
	//	} else {
	//		ref = append(ref, strcase.ToCamel(fmd.fieldName))
	//	}
	//	cacheKey = fmt.Sprintf("%s.%s", cacheKey, fmd.fieldName)
	//	if _, ok := done[cacheKey]; !ok {
	//		if fmd.repeated {
	//			if fmd.msg != nil {
	//				// todo missing length
	//				p.P(strings.Join(ref, "."), " := make(*", fmd.msg.Package(), ".", strcase.ToCamel(fmd.msg.Name()), ")")
	//			} else {
	//				p.P(strings.Join(ref, "."))
	//			}
	//			p.P()
	//		} else {
	//			if fmd.msg == nil {
	//				continue
	//			}
	//			p.P(strcase.ToLowerCamel(fmd.msg.Name()), " := &", fmd.msg.Package(), ".", strcase.ToCamel(fmd.msg.Name()), "{}")
	//		}
	//		done[cacheKey] = struct{}{}
	//	}
	//}
	//lhs := ret
	//for _, fmd := range u.producerStack {
	//	lhs = fmt.Sprintf("%s.%s", lhs, strcase.ToCamel(fmd.fieldName))
	//}
	//rhs := strcase.ToLowerCamel(u.depStack[0].msg.Name())
	//for idx, fmd := range u.depStack {
	//	if idx == 0 {
	//		continue
	//	}
	//	rhs = fmt.Sprintf("%s.%s", rhs, strcase.ToCamel(fmd.fieldName))
	//}
	//p.P(lhs, " = ", rhs)
	//return ret

}

type pipedContext struct {
	units       []*pipedFieldUnit
	stages      []stage.Stage
	producesMsg registry.Message
	api         dsl.ApiDescriptor
	meta        dsl.Meta
	partialOwn  bool
}

func (pc *pipedContext) PrintProto(printer.Factory) {
}

func (pc *pipedContext) PrintCode(printerFactory printer.Factory) error {
	p := printerFactory.Get(fmt.Sprintf("%s_transformer.go", pc.api.FileName()))
	p.P("package ", pc.api.Group(), "_v", pc.api.Version())
	preparePipedImports(p, pc.meta, pc.producesMsg, pc.dependencies())
	var dep []string
	for _, d := range pc.dependencies() {
		dep = append(dep, fmt.Sprintf("%s *%s.%s", strcase.ToLowerCamel(d.Name()), d.Package(), strcase.ToCamel(d.Name())))
	}
	if pc.partialOwn {
		p.P("func transformTo", strcase.ToCamel(pc.producesMsg.Name()), "(",
			strcase.ToLowerCamel(pc.producesMsg.Name()), " *", pc.producesMsg.Package(), ".", strcase.ToCamel(pc.producesMsg.Name()), ", ",
			strings.Join(dep, ", "), ") ", "{")
	} else {
		p.P("func transformTo", strcase.ToCamel(pc.producesMsg.Name()), "(", strings.Join(dep, ", "), ") *", pc.producesMsg.Package(), ".", strcase.ToCamel(pc.producesMsg.Name()), "{")
	}

	p.Tab()
	done := map[string]struct{}{}
	var ret []string
	for _, pfu := range pc.units {
		retRef, er := pfu.printCode(p, done, pc.partialOwn)
		if er != nil {
			return er
		}
		if retRef != "" {
			ret = append(ret, retRef)
		}
	}
	p.P("return ", strings.Join(ret, ", "))
	p.UnTab()
	p.P("}")

	return nil
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
	if !pc.partialOwn {
		p.P(strcase.ToLowerCamel(pc.producesMsg.Name()), " := transformTo", strcase.ToCamel(pc.producesMsg.Name()), "(", strings.Join(dep, ", "), ")")
	} else {
		// todo maybe a more suitable prefix
		p.P("transformTo", strcase.ToCamel(pc.producesMsg.Name()), "(", strcase.ToLowerCamel(pc.producesMsg.Name()), ", ", strings.Join(dep, ", "), ")")
	}
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

func groupPipedUnitsByProduces(api dsl.ApiDescriptor, meta dsl.Meta, units []*pipedFieldUnit, argUnits []argUnit) ([]*pipedContext, error) {
	producerMap := map[registry.Message][]*pipedFieldUnit{}
	for _, u := range units {
		if _, ok := producerMap[u.producerStack[0].msg]; !ok {
			producerMap[u.producerStack[0].msg] = nil
		}
		producerMap[u.producerStack[0].msg] = append(producerMap[u.producerStack[0].msg], u)
	}
	auProducerMap := map[registry.Message]struct{}{}
	for _, au := range argUnits {
		produces := au.produces()
		if len(produces) > 1 {
			panic("impossible assertion")
		}
		msg := produces[0]
		if _, ok := producerMap[msg]; ok {
			auProducerMap[msg] = struct{}{}
		}
	}

	var ret []*pipedContext
	for msg, v := range producerMap {
		partialOwnership := false
		if _, ok := auProducerMap[msg]; ok {
			partialOwnership = true
		}
		ret = append(ret, &pipedContext{units: v, producesMsg: msg, api: api, meta: meta, partialOwn: partialOwnership})
	}
	return ret, nil
}

func applyPiping(pipeArg *dsl.PipedArgDescriptor, units []argUnit, r registry.Registry) (*pipedFieldUnit, error) {
	rawArgPath := pipeArg.Output()
	for _, au := range units {
		argPaths := strings.Split(rawArgPath, ".")
		if ok, fau, er := matchAndRemoveArgPath(au, argPaths, 0); er != nil {
			return nil, er
		} else if ok {
			msgs := r.ListMessages(registry.LMOWithPrefixMatch(pipeArg.Input()))
			if len(msgs) != 1 {
				return nil, errors.Errorf("piped return must match exactly one message found %d for %s", len(msgs), pipeArg.Input())
			}
			depStack, err := resolveField(msgs[0], pipeArg.Input())
			if err != nil {
				return nil, err
			}
			// todo assert equality in num of repeated
			return &pipedFieldUnit{producerStack: fau.producerStack[0], depStack: depStack}, nil
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
		repeated := false
		if strings.HasSuffix(fd, "[]") {
			repeated = true
			fd = strings.Trim(fd, "[]")
		}
		if idx == 0 && fd == resolvedMsg.Package() {
			continue
		}
		if idx <= 1 && fd == resolvedMsg.Name() {
			continue
		}
		for _, field := range resolvedMsg.Fields() {
			if field.Name() == fd {
				if (repeated || field.Repeated()) && !(repeated && field.Repeated()) {
					return nil, errors.Errorf("mismatched array types ")
				}
				found = true
				if field.Type() != registry.FieldTypeMessage {
					if idx != len(splits)-1 {
						return nil, errors.Errorf("partially resolved to field %s", input)
					}
					resolvedField = field
					dependencyStack = append(dependencyStack, fieldMessageDependency{fieldName: field.Name(), repeated: field.Repeated(), fieldType: field.Type()})
				} else {
					resolvedMsg = field.Message()
					dependencyStack = append(dependencyStack, fieldMessageDependency{fieldName: field.Name(), msg: resolvedMsg, repeated: field.Repeated(), fieldType: field.Type()})
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
