package inresolution

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/generator/stage"
	"github.com/carousell/aggproto/pkg/registry"
	"github.com/iancoleman/strcase"
)

type InputContext struct {
	apiDescriptor dsl.ApiDescriptor
	required      []registry.Message
	meta          dsl.Meta
}

func (i *InputContext) AddStageDependency(stage stage.Stage) {
	panic("input stage cannot have a dependency")
}

func (i *InputContext) GetStageDependencies() []stage.Stage {
	return nil
}

func printMessage(p printer.Printer, m registry.Message) {
	p.P("message ", fmt.Sprintf("%sGen {", m.Name()))
	p.Tab()
	for _, subDef := range m.Definitions() {
		printMessage(p, subDef)
	}
	for idx, f := range m.Fields() {
		switch f.Type() {
		case registry.FieldTypeString:
			p.P("string ", f.Name(), " = ", idx+1, ";")
		case registry.FieldTypeInt64:
			p.P("int64 ", f.Name(), " = ", idx+1, ";")
		case registry.FieldTypeDouble:
			p.P("float64 ", f.Name(), " = ", idx+1, ";")
		case registry.FieldTypeBool:
			p.P("bool ", f.Name(), " = ", idx+1, ";")
		case registry.FieldTypeMessage:
			p.P(strcase.ToCamel(f.Name()), "Gen ", strcase.ToSnake(f.Name()), " = ", idx+1, ";")
		default:
			panic("unsupported type " + f.Name())
		}
	}
	p.UnTab()
	p.P("}")
}

func (i *InputContext) PrintProto(factory printer.Factory) {
	p := factory.Get(fmt.Sprintf("%s.proto", i.apiDescriptor.FileName()))
	p.P("message ", fmt.Sprintf("%sRequest", strcase.ToCamel(i.apiDescriptor.Name())), " {")
	p.Tab()
	for _, r := range i.required {
		printMessage(p, r)
	}
	for idx, r := range i.required {
		p.P(r.Name(), "Gen ", strcase.ToSnake(r.Name()), " = ", idx+1, ";")
	}
	p.UnTab()
	p.P("}")
}

func (i *InputContext) PrintCodeUsage(p printer.Printer) {
	var retReferences []string
	for _, r := range i.required {
		retReferences = append(retReferences, strcase.ToLowerCamel(r.Name()))
	}
	if len(retReferences) == 0 {
		return
	}
	p.P(strings.Join(retReferences, ", "), " := transform", strcase.ToCamel(i.apiDescriptor.Name()), "Request(req)")
}
func (i *InputContext) PrintCode(printerFactory printer.Factory) {
	p := printerFactory.Get(fmt.Sprintf("%s_input.go", i.apiDescriptor.FileName()))
	p.P("package ", i.apiDescriptor.Group(), "_v", i.apiDescriptor.Version())
	reqClassName := fmt.Sprintf("%sRequest", strcase.ToCamel(i.apiDescriptor.Name()))
	p.P()
	prepareImports(p, i.meta, i.required)
	p.P()
	p.P("func ", "transform", reqClassName, "(req *", reqClassName, ") ", prepareRequiredReturns(i.required), " {")
	p.Tab()
	var retReferences []string
	for _, r := range i.required {
		retReferences = append(retReferences, prepareRequired(p, r, fmt.Sprintf("req.%s", r.Name())))
	}
	p.P("return ", strings.Join(retReferences, ", "))
	p.UnTab()
	p.P("}")
}

func prepareRequired(p printer.Printer, msg registry.Message, referenceName string) string {
	fieldName := strcase.ToLowerCamel(msg.Name())
	p.P(fieldName, " := &", msg.Package(), ".", msg.Name(), "{}")
	for _, f := range msg.Fields() {
		switch f.Type() {
		case registry.FieldTypeMessage:
			prepareRequired(p, f.Message(), fmt.Sprintf("%s.%s", referenceName, f.Name()))
		default:
			p.P(fieldName, ".", strcase.ToCamel(f.Name()), " = ", referenceName, ".", strcase.ToCamel(f.Name()))
		}
	}
	return fieldName
}

func prepareImports(p printer.Printer, meta dsl.Meta, required []registry.Message) {
	packages := map[string]struct{}{}
	for _, r := range required {
		packages[r.Package()] = struct{}{}
	}
	if len(packages) > 0 {
		p.P("import (")
		p.Tab()
		for pkg, _ := range packages {
			p.P("\"", meta.GoPackage, "/", pkg, "\"")
		}
		p.UnTab()
		p.P(")")
	}
}

func prepareRequiredReturns(required []registry.Message) string {
	buf := new(bytes.Buffer)
	if len(required) > 1 {
		buf.WriteString("(")
	}
	var ret []string
	for _, r := range required {
		ret = append(ret, fmt.Sprintf("*%s.%s", r.Package(), r.Name()))
	}
	buf.WriteString(strings.Join(ret, ", "))
	if len(required) > 1 {
		buf.WriteString(")")
	}
	return buf.String()
}
