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
	prepareRequiredAssign(p printer.Printer, refName string, repIdx []string)
	prepareRequiredDefine(p printer.Printer, refName string, done map[registry.Message]struct{}, repIdx []string) []string
}

type nestedArgUnit struct {
	fieldName  string
	nestedArgs []argUnit
	ctx        registry.Message
	repeated   bool
}

func (u *nestedArgUnit) prepareRequiredDefine(p printer.Printer, refName string, done map[registry.Message]struct{}, repIdx []string) []string {
	var ret []string
	for _, au := range u.nestedArgs {
		ret = append(ret, au.prepareRequiredDefine(p, refName, done, repIdx)...)
	}
	return ret
}

func (u *nestedArgUnit) prepareRequiredAssign(p printer.Printer, refName string, repIdx []string) {
	if u.repeated {
		repIdx = append(repIdx, "i")
		lenName := fmt.Sprintf("len(%s.%s)", refName, strcase.ToCamel(u.fieldName))
		refName = fmt.Sprintf("%s.%s[i]", refName, strcase.ToCamel(u.fieldName))
		p.P("for i := 0; i < ", lenName, "; i += 1 {")
		p.Tab()
		for _, au := range u.nestedArgs {
			au.prepareRequiredAssign(p, refName, repIdx)
		}
		p.UnTab()
		p.P("}")
	} else {
		refName = fmt.Sprintf("%s.%s", refName, strcase.ToCamel(u.fieldName))
		for _, au := range u.nestedArgs {
			au.prepareRequiredAssign(p, refName, repIdx)
		}
	}
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
	if u.repeated {
		p.P("repeated ", strcase.ToCamel(u.fieldName), "Gen ", strcase.ToSnake(u.fieldName), " = ", idx+1, ";")
	} else {
		p.P(strcase.ToCamel(u.fieldName), "Gen ", strcase.ToSnake(u.fieldName), " = ", idx+1, ";")
	}
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
			if !found {
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
	repeated      bool
}

func (f *fieldArgUnit) prepareRequiredDefine(p printer.Printer, refName string, done map[registry.Message]struct{}, repIdx []string) []string {
	var ret []string
	for _, fmds := range f.producerStack {
		for idx, fmd := range fmds {
			if _, ok := done[fmd.msg]; ok {
				continue
			}
			if fmd.msg != nil {
				p.P(strcase.ToLowerCamel(fmd.fieldName), " := &", fmd.msg.Package(), ".", fmd.msg.Name(), "{}")
				done[fmd.msg] = struct{}{}
			}
			if idx == 0 {
				ret = append(ret, strcase.ToLowerCamel(fmd.fieldName))
			}
		}
	}
	return ret
}

func (f *fieldArgUnit) prepareRequiredAssign(p printer.Printer, refName string, repIdx []string) {
	for _, fmds := range f.producerStack {
		var retRef []string
		numRep := 0
		for idx, fmd := range fmds {
			var refSuffix string
			if fmd.repeated && len(repIdx) > numRep {
				refSuffix = fmt.Sprintf("[%s]", repIdx[numRep])
				numRep += 1
			}
			if idx == 0 {
				retRef = append(retRef, fmt.Sprintf("%s%s", strcase.ToLowerCamel(fmd.fieldName), refSuffix))
			} else {
				retRef = append(retRef, fmt.Sprintf("%s%s", strcase.ToCamel(fmd.fieldName), refSuffix))
			}
		}
		p.P(strings.Join(retRef, "."), " = ", refName, ".", strcase.ToCamel(f.fieldName))
	}
}

func (f *fieldArgUnit) produces() []registry.Message {
	var ret []registry.Message
	for _, p := range f.producerStack {
		ret = append(ret, p[0].msg)
	}
	return ret
}

func (f *fieldArgUnit) printProtoUsage(p printer.Printer, idx int) {
	var prefix string
	if f.repeated {
		prefix = "repeated "
	}
	switch f.fieldType {
	case registry.FieldTypeString:
		p.P(prefix, "string ", f.fieldName, " = ", idx+1, ";")
	case registry.FieldTypeInt64:
		p.P(prefix, "int64 ", f.fieldName, " = ", idx+1, ";")
	case registry.FieldTypeDouble:
		p.P(prefix, "float64 ", f.fieldName, " = ", idx+1, ";")
	case registry.FieldTypeBool:
		p.P(prefix, "bool ", f.fieldName, " = ", idx+1, ";")
	}
}

func (f *fieldArgUnit) getFieldName() string {
	return f.fieldName
}

func (f *fieldArgUnit) printProtoDependency(p printer.Printer, idx int) {
}
