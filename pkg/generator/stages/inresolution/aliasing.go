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
	"strings"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/pkg/errors"
)

func applyAliasing(aliasArg *dsl.AliasArgDescriptor, argUnits []argUnit) (argUnit, error) {
	rawArgPath := aliasArg.Output()
	for _, au := range argUnits {
		argPaths := strings.Split(rawArgPath, ".")
		if ok, fau, er := matchAndRemoveArgPath(au, argPaths, 0); ok {
			return createFieldArgUnit(aliasArg.Input(), fau), nil
		} else if er != nil {
			return nil, er
		}
	}
	return nil, nil
}
func createFieldArgUnit(input string, fau *fieldArgUnit) argUnit {
	splits := strings.Split(input, ".")
	var au argUnit
	for i := len(splits) - 1; i >= 0; i-- {
		fieldName := splits[i]
		repeated := false
		if strings.HasSuffix(splits[i], "[]") {
			repeated = true
			fieldName = strings.Trim(fieldName, "[]")
		}
		if au == nil {
			au = &fieldArgUnit{fieldName, fau.fieldType, fau.producerStack, repeated}
		} else {
			au = &nestedArgUnit{fieldName: fieldName, nestedArgs: []argUnit{au}, repeated: repeated}
		}
	}
	return au
}

func matchAndRemoveArgPath(au argUnit, argPaths []string, depth int) (bool, *fieldArgUnit, error) {
	if len(argPaths) == 0 {
		return false, nil, nil
	}
	nextArgPath := argPaths[0]
	repeated := false
	if strings.HasSuffix(nextArgPath, "[]") {
		nextArgPath = strings.Trim(nextArgPath, "[]")
		repeated = true
	}
	if nau, ok := au.(*nestedArgUnit); ok {
		if depth == 0 && nextArgPath == nau.ctx.Package() {
			return matchAndRemoveArgPath(au, argPaths[1:], depth+1)
		}
		if nextArgPath != nau.fieldName {
			return false, nil, nil
		}
		if repeated || nau.repeated && !(nau.repeated && repeated) {
			return false, nil, errors.Errorf("mismatched array indicators in input")

		}
		for idx, cau := range nau.nestedArgs {
			matched, fau, er := matchAndRemoveArgPath(cau, argPaths[1:], depth+1)
			if er != nil {
				return false, nil, er
			}
			if matched {
				if cau == fau {
					nau.nestedArgs = deleteFromAUSlice(nau.nestedArgs, idx)
				} else {
					if cnau, ok := cau.(*nestedArgUnit); ok {
						if len(cnau.nestedArgs) == 0 {
							deleteFromAUSlice(nau.nestedArgs, idx)
						}
					}
				}
				return true, fau, nil
			}
		}
	}
	if fau, ok := au.(*fieldArgUnit); ok {
		if fau.fieldName == nextArgPath && len(argPaths) == 1 {
			return true, fau, nil
		}
	}
	return false, nil, nil
}

func deleteFromAUSlice(args []argUnit, idx int) []argUnit {
	if len(args) == 1 && idx == 0 {
		return nil
	}
	if idx == 0 {
		return args[1:]
	}
	tAU := args[0]
	args[0] = args[idx]
	args[idx] = tAU
	return args[1:]
}
