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
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/stages/msgresolution"
	"github.com/carousell/aggproto/pkg/generator/stages/opresolution"
	"github.com/carousell/aggproto/pkg/registry"
)

type inResolver struct {
	r registry.Registry
}

func New(r registry.Registry) InputResolver {
	return &inResolver{r: r}
}

func (i *inResolver) Resolve(api dsl.ApiDescriptor, meta dsl.Meta, input []dsl.ArgDescriptor, opCtxs []opresolution.OperationContext, adaptorContext msgresolution.AdaptorContext) (*InputContext, error) {
	var requiredInput []registry.Message
	requiredInputsMap := map[string]struct{}{}
	for _, op := range opCtxs {
		msg := op.InputDependency()
		if _, ok := requiredInputsMap[msg.FullName()]; !ok {
			requiredInputsMap[msg.FullName()] = struct{}{}
			requiredInput = append(requiredInput, msg)
		}
	}
	argUnits := explode(requiredInput)
	var createdAU []argUnit
	var pipedUnits []*pipedFieldUnit
	for _, in := range input {
		if aliasArg, ok := in.(*dsl.AliasArgDescriptor); ok {
			au, er := applyAliasing(aliasArg, argUnits)
			if er != nil {
				return nil, er
			}
			if au != nil {
				createdAU = append(createdAU, au)
			}
		}
		if pipeArg, ok := in.(*dsl.PipedArgDescriptor); ok {
			pfu, er := applyPiping(pipeArg, argUnits, i.r)
			if er != nil {
				return nil, er
			}
			pipedUnits = append(pipedUnits, pfu)
		}
	}
	allAU := filterEmpty(append(argUnits, createdAU...))
	pipedContexts, _ := groupPipedUnitsByProduces(api, meta, pipedUnits, allAU)
	err := addPipedDependencies(pipedContexts, opCtxs)
	if err != nil {
		return nil, err
	}
	argUnits, err = mergeArgUnits(filterEmpty(allAU))
	if err != nil {
		return nil, err
	}
	inCtx := &InputContext{apiDescriptor: api, required: requiredInput, meta: meta, argUnits: argUnits}
	if len(opCtxs) > 0 {
		for _, opCtx := range opCtxs {
			opCtx.AddStageDependency(inCtx)
		}
	} else {
		adaptorContext.AddStageDependency(inCtx)
	}
	return inCtx, nil
}

func filterEmpty(units []argUnit) []argUnit {
	var ret []argUnit
	for _, au := range units {
		if au == nil {
			continue
		}
		if _, ok := au.(*fieldArgUnit); ok {
			ret = append(ret, au)
		} else if nau, ok := au.(*nestedArgUnit); ok {
			if len(nau.nestedArgs) != 0 {
				ret = append(ret, au)
			}
		}
	}
	return ret
}

func mergeArgUnits(units []argUnit) ([]argUnit, error) {
	var ret []argUnit
	naus := map[string]*nestedArgUnit{}
	for _, au := range units {
		if nau, ok := au.(*nestedArgUnit); ok {
			if existingNau, ok := naus[nau.fieldName]; ok {
				err := existingNau.tryMerge(nau)
				if err != nil {
					return nil, err
				}
			} else {
				naus[nau.fieldName] = nau
				ret = append(ret, nau)
			}
		} else {
			//todo
		}
	}
	return ret, nil
}

func explode(input []registry.Message) []argUnit {
	var ret []argUnit
	for _, msg := range input {
		ret = append(ret, explodeMessage(msg.Name(), msg, []fieldMessageDependency{{msg: msg, fieldName: msg.Name(), repeated: false, fieldType: registry.FieldTypeMessage}}))
	}
	return ret
}

type fieldMessageDependency struct {
	msg       registry.Message
	fieldName string
	repeated  bool
	fieldType registry.FieldType
}

//stack
func explodeMessage(name string, msg registry.Message, producesStack []fieldMessageDependency) *nestedArgUnit {
	au := &nestedArgUnit{fieldName: name, ctx: msg}

	for _, f := range msg.Fields() {
		switch f.Type() {
		case registry.FieldTypeMessage:
			nau := explodeMessage(f.Name(), f.Message(), append(producesStack, fieldMessageDependency{f.Message(), f.Name(), f.Repeated(), f.Type()}))
			nau.repeated = f.Repeated()
			au.nestedArgs = append(au.nestedArgs, nau)
		default:
			producesStack := append(producesStack, fieldMessageDependency{nil, f.Name(), f.Repeated(), f.Type()})
			au.nestedArgs = append(au.nestedArgs, &fieldArgUnit{fieldName: f.Name(), fieldType: f.Type(), producerStack: [][]fieldMessageDependency{producesStack}, repeated: f.Repeated()})
		}
	}
	return au
}
