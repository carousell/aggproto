package msgresolution

import (
	"fmt"

	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/registry"
	"github.com/iancoleman/strcase"
)

type fieldMessageDependency struct {
	fieldName string
	message   registry.Message
}
type adaptorUnit interface {
	isAdaptorUnit()
	printProtoDefinitions(p printer.Printer, fieldIdx int)
	printAsProtoField(p printer.Printer, idx int)
	printAsAdaptorCode(p printer.Printer, referenceName string)
	dependencies() [][]fieldMessageDependency
}

type adaptorUnits []adaptorUnit

func (au adaptorUnits) tryMerge() adaptorUnits {
	var mergedUnits []adaptorUnit
	for i := 0; i < len(au); i += 1 {
		if au[i] == nil {
			continue
		}
		if nau, ok := au[i].(*nestedAdaptorUnit); ok {
			for j := i + 1; j < len(au); j += 1 {
				if onau, ok := au[j].(*nestedAdaptorUnit); ok {
					if nau.merge(onau) {
						au[j] = nil
					}
				}
			}
		}
		mergedUnits = append(mergedUnits, au[i])
	}
	return mergedUnits
}

type nestedAdaptorUnit struct {
	fieldName  string
	nestedUnit []adaptorUnit
}

func (n *nestedAdaptorUnit) dependencies() [][]fieldMessageDependency {
	var deps [][]fieldMessageDependency
	for _, nu := range n.nestedUnit {
		deps = append(deps, nu.dependencies()...)
	}
	return deps
}

func (n *nestedAdaptorUnit) printAsAdaptorCode(p printer.Printer, referenceName string) {
	fieldName := strcase.ToCamel(n.fieldName)
	p.P(referenceName, ".", fieldName, " = &", fieldName, "Gen{}")
	for _, au := range n.nestedUnit {
		au.printAsAdaptorCode(p, fmt.Sprintf("%s.%s", referenceName, fieldName))
	}
}

func (n *nestedAdaptorUnit) printAsProtoField(p printer.Printer, idx int) {
	p.P(fmt.Sprintf("%sGen ", strcase.ToCamel(n.fieldName)), n.fieldName, " = ", idx, ";")
}

func (n *nestedAdaptorUnit) printProtoDefinitions(p printer.Printer, fieldIdx int) {
	p.P("message ", fmt.Sprintf("%sGen", strcase.ToCamel(n.fieldName)), "{")
	p.Tab()
	for idx, au := range n.nestedUnit {
		au.printProtoDefinitions(p, idx+1)
	}
	for idx, au := range n.nestedUnit {
		au.printAsProtoField(p, idx+1)
	}
	p.UnTab()
	p.P("}")
}

func (n *nestedAdaptorUnit) merge(other *nestedAdaptorUnit) bool {
	if n.fieldName != other.fieldName {
		return false
	}
onuLoop:
	for _, onu := range other.nestedUnit {
		if onu, ok := onu.(*nestedAdaptorUnit); ok {
			for _, nu := range n.nestedUnit {
				if nu, ok := nu.(*nestedAdaptorUnit); ok {
					if nu.merge(onu) {
						continue onuLoop
					}
				}
			}
		}
		n.nestedUnit = append(n.nestedUnit, onu)
	}
	return true
}
func (n *nestedAdaptorUnit) isAdaptorUnit() {
	panic("Should never be called")
}
