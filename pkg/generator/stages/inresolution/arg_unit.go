package inresolution

import (
	"errors"
	"fmt"
	"strings"

	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/registry"
	"github.com/iancoleman/strcase"
)

type argUnit interface {
	printProtoDependency(p printer.Printer, idx int)
	getFieldName() string
	printProtoUsage(p printer.Printer, idx int)
	produces() []registry.Message
	prepareRequired(p printer.Printer, refName string, done map[registry.Message]struct{}) []string
}

type nestedArgUnit struct {
	fieldName  string
	nestedArgs []argUnit
	ctx        registry.Message
}

func (u *nestedArgUnit) prepareRequired(p printer.Printer, refName string, done map[registry.Message]struct{}) []string {
	var ret []string
	for _, au := range u.nestedArgs {
		ret = append(ret, au.prepareRequired(p, fmt.Sprintf("%s.%s", refName, strcase.ToCamel(u.fieldName)), done)...)
	}
	return ret
}

func (u *nestedArgUnit) produces() []registry.Message {
	var ret []registry.Message
	done := map[registry.Message]struct{}{}
	for _, au := range u.nestedArgs {
		for _, msg := range au.produces() {
			if _, ok := done[msg]; !ok {
				ret = append(ret, msg)
				done[msg] = struct{}{}
			}
		}
	}
	return ret
}

func (u *nestedArgUnit) printProtoUsage(p printer.Printer, idx int) {
	p.P(strcase.ToCamel(u.fieldName), "Gen ", strcase.ToSnake(u.fieldName), " = ", idx+1, ";")
}

func (u *nestedArgUnit) getFieldName() string {
	return u.fieldName
}

func (u *nestedArgUnit) printProtoDependency(p printer.Printer, idx int) {
	p.P("message ", strcase.ToCamel(u.fieldName), "Gen", " {")
	p.Tab()
	for idx, au := range u.nestedArgs {
		au.printProtoDependency(p, idx)
	}
	for idx, au := range u.nestedArgs {
		au.printProtoUsage(p, idx)
	}
	p.UnTab()
	p.P("}")
}

func (u *nestedArgUnit) tryMerge(nau *nestedArgUnit) error {
	if u.fieldName != nau.fieldName {
		return errors.New("cannot merge dis-similar nau")
	}
	for _, v := range nau.nestedArgs {
		found := false
		if fau, ok := v.(*fieldArgUnit); ok {
			for _, au := range u.nestedArgs {
				if fau.fieldName == au.getFieldName() {
					if thisFau, ok := au.(*fieldArgUnit); ok {
						if fau.fieldType == thisFau.fieldType {
							thisFau.producerStack = append(thisFau.producerStack, fau.producerStack...)
						} else {
							return errors.New("field types mismatch")
						}
					} else {
						return errors.New("cannot merge field arg unit with non field arg")
					}
				}
				found = true
			}
			if !found{
				u.nestedArgs = append(u.nestedArgs, fau)
			}
		} else if childNau, ok := v.(*nestedArgUnit); ok {
			childFound := false
			for _, au := range u.nestedArgs {
				if childNau.fieldName == au.getFieldName() {
					childFound = true
					if existingChildNau, ok := au.(*nestedArgUnit); ok {
						err := existingChildNau.tryMerge(childNau)
						if err != nil {
							return err
						}
					} else if _, ok := au.(*fieldArgUnit); ok {
						return errors.New("cannot merge nau to fau")
					}
					break
				}
			}
			if !childFound {
				u.nestedArgs = append(u.nestedArgs, childNau)
			}
		}
	}
	return nil
}

type fieldArgUnit struct {
	fieldName     string
	fieldType     registry.FieldType
	producerStack [][]fieldMessageDependency
}

func (f *fieldArgUnit) prepareRequired(p printer.Printer, refName string, done map[registry.Message]struct{}) []string {
	var ret []string
	for _, fmds := range f.producerStack {
		for idx, fmd := range fmds {
			if _, ok := done[fmd.msg]; ok {
				continue
			}
			p.P(strcase.ToLowerCamel(fmd.msg.Name()), " := &", fmd.msg.Package(), ".", fmd.msg.Name(), "{}")
			done[fmd.msg] = struct{}{}
			if idx == 0 {
				ret = append(ret, strcase.ToLowerCamel(fmd.msg.Name()))
			}
		}
	}
	for _, fmds := range f.producerStack {
		var retRef []string
		for idx, fmd := range fmds {
			if idx == 0 {
				retRef = append(retRef, strcase.ToLowerCamel(fmd.msg.Name()))
			}
			retRef = append(retRef, strcase.ToCamel(fmd.fieldName))

		}
		p.P(strings.Join(retRef, "."), " = ", refName, ".", strcase.ToCamel(f.fieldName))
	}
	return ret
}

func (f *fieldArgUnit) produces() []registry.Message {
	var ret []registry.Message
	for _, p := range f.producerStack {
		ret = append(ret, p[0].msg)
	}
	return ret
}

func (f *fieldArgUnit) printProtoUsage(p printer.Printer, idx int) {
	switch f.fieldType {
	case registry.FieldTypeString:
		p.P("string ", f.fieldName, " = ", idx+1, ";")
	case registry.FieldTypeInt64:
		p.P("int64 ", f.fieldName, " = ", idx+1, ";")
	case registry.FieldTypeDouble:
		p.P("float64 ", f.fieldName, " = ", idx+1, ";")
	case registry.FieldTypeBool:
		p.P("bool ", f.fieldName, " = ", idx+1, ";")
	}
}

func (f *fieldArgUnit) getFieldName() string {
	return f.fieldName
}

func (f *fieldArgUnit) printProtoDependency(p printer.Printer, idx int) {
}
