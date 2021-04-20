package msgresolution

import (
	"fmt"
	"strings"

	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/registry"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
)

type fieldMessageDependency struct {
	fieldName string
	message   registry.Message
	repeated  bool
}
type adaptorUnit interface {
	isAdaptorUnit()
	printProtoDefinitions(p printer.Printer)
	printAsProtoField(p printer.Printer, idx int)
	printAsAdaptorCode(p printer.Printer, referenceName string, parents []string, repeatedStringIdxRef []string) error
	dependencies() [][]fieldMessageDependency
	getRepeatedSizeString() ([]string, error)
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
	repeated   bool
}

func (n *nestedAdaptorUnit) getRepeatedSizeString() ([]string, error) {
	var repeatedSizeStr []string
	done := map[string]struct{}{}
	for _, au := range n.nestedUnit {
		rss, er := au.getRepeatedSizeString()
		if er != nil {
			return nil, er
		}
		if len(rss) == 0 {
			continue
		}
		if len(repeatedSizeStr) == 0 {
			repeatedSizeStr = rss
			for _, _rss := range rss {
				done[_rss] = struct{}{}
			}
		} else {
			for _, _rss := range rss {
				if _, ok := done[_rss]; ok {
					continue
				}
				repeatedSizeStr = append(repeatedSizeStr, _rss)
				done[_rss] = struct{}{}
			}
		}
	}
	if len(repeatedSizeStr) == 0 {
		return nil, errors.Errorf("No repeated field found")
	}
	return repeatedSizeStr, nil
}

func (n *nestedAdaptorUnit) dependencies() [][]fieldMessageDependency {
	var deps [][]fieldMessageDependency
	for _, nu := range n.nestedUnit {
		deps = append(deps, nu.dependencies()...)
	}
	return deps
}

func (n *nestedAdaptorUnit) printAsAdaptorCode(p printer.Printer, referenceName string, parents []string, repeatedStringIdxRef []string) error {
	fieldName := strcase.ToCamel(n.fieldName)
	parentName := strings.Join(parents, "_")
	if n.repeated {
		rss, err := n.getRepeatedSizeString()
		if err != nil {
			return err
		}
		for i := 0; i < len(rss)-1; i += 1 {
			for j := i + 1; j < len(rss); j += 1 {
				p.P("if ", rss[i], " != ", rss[j], " {")
				p.Tab()
				p.P("return nil, errors.Errorf(\"assertion failed %s != %s\", ", rss[i],", ", rss[j],")" )
				p.UnTab()
				p.P("}")
			}
		}
		p.P(referenceName, ".", fieldName, " = make([]*", parentName, "_", fieldName, "Gen, ", rss[0], ")")
		p.P("for i := 0; i < ", rss[0], "; i++ {")
		p.Tab()
		p.P(referenceName, ".", fieldName, "[i] = &", parentName, "_", fieldName, "Gen{}")
		for _, au := range n.nestedUnit {
			err := au.printAsAdaptorCode(p, fmt.Sprintf("%s.%s[i]", referenceName, fieldName), append(parents, fmt.Sprintf("%sGen", fieldName)), append(repeatedStringIdxRef, "i"))
			if err != nil {
				return err
			}
		}
		p.UnTab()
		p.P("}")
	} else {
		p.P(referenceName, ".", fieldName, " = &", parentName, "_", fieldName, "Gen{}")
		for _, au := range n.nestedUnit {
			err := au.printAsAdaptorCode(p, fmt.Sprintf("%s.%s", referenceName, fieldName), append(parents, fmt.Sprintf("%sGen", fieldName)), repeatedStringIdxRef)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (n *nestedAdaptorUnit) printAsProtoField(p printer.Printer, idx int) {
	if n.repeated {
		p.P("repeated ", fmt.Sprintf("%sGen ", strcase.ToCamel(n.fieldName)), n.fieldName, " = ", idx, ";")
	} else {
		p.P(fmt.Sprintf("%sGen ", strcase.ToCamel(n.fieldName)), n.fieldName, " = ", idx, ";")
	}
}

func (n *nestedAdaptorUnit) printProtoDefinitions(p printer.Printer) {
	p.P("message ", fmt.Sprintf("%sGen", strcase.ToCamel(n.fieldName)), " {")
	p.Tab()
	for _, au := range n.nestedUnit {
		au.printProtoDefinitions(p)
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
