package inresolution

import (
	"strings"

	"github.com/carousell/aggproto/pkg/dsl"
)

func applyAliasing(aliasArg *dsl.AliasArgDescriptor, argUnits []argUnit) argUnit {
	rawArgPath := aliasArg.Output()
	for _, au := range argUnits {
		argPaths := strings.Split(rawArgPath, ".")
		if ok, fau := matchAndRemoveArgPath(au, argPaths, 0); ok {
			return createFieldArgUnit(aliasArg.Input(), fau)
		}
	}
	return nil
}
func createFieldArgUnit(input string, fau *fieldArgUnit) argUnit {
	splits := strings.Split(input, ".")
	var au argUnit
	for i := len(splits) - 1; i >= 0; i-- {
		if au == nil {
			au = &fieldArgUnit{splits[i], fau.fieldMsg, fau.producerStack}
		} else {
			au = &nestedArgUnit{fieldName: splits[i], nestedArgs: []argUnit{au}}
		}
	}
	return au
}

func matchAndRemoveArgPath(au argUnit, argPaths []string, depth int) (bool, *fieldArgUnit) {
	if len(argPaths) == 0 {
		return false, nil
	}
	nextArgPath := argPaths[0]
	if nau, ok := au.(*nestedArgUnit); ok {
		if depth == 0 && nextArgPath == nau.ctx.Package() {
			return matchAndRemoveArgPath(au, argPaths[1:], depth+1)
		}
		if nextArgPath != nau.fieldName {
			return false, nil
		}
		for idx, cau := range nau.nestedArgs {
			matched, fau := matchAndRemoveArgPath(cau, argPaths[1:], depth+1)
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
				return true, fau
			}
		}
	}
	if fau, ok := au.(*fieldArgUnit); ok {
		if fau.fieldName == nextArgPath && len(argPaths) == 1 {
			return true, fau
		}
	}
	return false, nil
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
