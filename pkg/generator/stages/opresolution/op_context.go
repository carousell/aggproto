//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package opresolution

import (
	"fmt"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/generator/stage"
	"github.com/carousell/aggproto/pkg/registry"
	"github.com/iancoleman/strcase"
)

type opContext struct {
	operation       registry.UnaryOperation
	api             dsl.ApiDescriptor
	meta            dsl.Meta
	dependentStages []stage.Stage
}

func (o *opContext) Produces() registry.Message {
	return o.operation.Output()
}

func (o *opContext) Consumes() registry.Message {
	return o.operation.Input()
}

func (o *opContext) RequiredImport() string {
	return fmt.Sprintf("%s/%s", o.meta.GoPackage, o.operation.Context().Package())
}

func (o *opContext) ClientDependency() string {
	return fmt.Sprintf("%s %s.%sClient", strcase.ToLowerCamel(o.operation.Context().Name()), o.operation.Context().Package(), o.operation.Context().Name())
}

func (o *opContext) PrintConstructorUsage(p printer.Printer) {
	p.P(strcase.ToLowerCamel(o.operation.Name()), "Client: &", strcase.ToLowerCamel(o.operation.Name()), "Client{", strcase.ToLowerCamel(o.operation.Context().Name()), "},")
}

func (o *opContext) ClientReferenceName() string {
	return fmt.Sprintf("%sClient *%sClient", strcase.ToLowerCamel(o.operation.Name()), strcase.ToLowerCamel(o.operation.Name()))
}

func (o *opContext) AddStageDependency(stage stage.Stage) {
	o.dependentStages = append(o.dependentStages, stage)
}

func (o *opContext) GetStageDependencies() []stage.Stage {
	return o.dependentStages
}

func (o *opContext) PrintCodeUsage(p printer.Printer) {
	p.P(strcase.ToLowerCamel(o.operation.Output().Name()), ", err := s.", strcase.ToLowerCamel(o.operation.Name()), "Client.",
		strcase.ToLowerCamel(o.operation.Name()), "(ctx, ", strcase.ToLowerCamel(o.operation.Input().Name()), ")")
	p.P("if err != nil {")
	p.Tab()
	p.P("return nil, err")
	p.UnTab()
	p.P("}")
}

func (o *opContext) PrintProto(p printer.Factory) {
}

func (o *opContext) PrintCode(printerFactory printer.Factory) error {
	p := printerFactory.Get(fmt.Sprintf("%s_%s_operation.go", o.api.FileName(), strcase.ToSnake(o.operation.Name())))
	p.P("package ", o.api.Group(), "_v", o.api.Version())
	p.P()
	prepareImports(p, o.meta, o.operation)
	p.P()
	p.P("type ", strcase.ToLowerCamel(o.operation.Name()), "Client struct {")
	p.Tab()
	p.P("client ", o.operation.Context().Package(), ".", o.operation.Context().Name(), "Client")
	p.UnTab()
	p.P("}")
	p.P()
	p.P("func (c *", strcase.ToLowerCamel(o.operation.Name()), "Client) ", strcase.ToLowerCamel(o.operation.Name()),
		"(ctx context.Context, req *", o.operation.Input().Package(), ".", o.operation.Input().Name(), ") ( *",
		o.operation.Output().Package(), ".", o.operation.Output().Name(), ", error) {")
	p.Tab()
	p.P("resp, err := c.client.", o.operation.Name(), "(ctx, req)")
	p.P("if err != nil {")
	p.Tab()
	p.P("return nil, errors.Wrap(err, \"error invoking ", o.operation.Name(), "\")")
	p.UnTab()
	p.P("}")
	p.P("return resp, nil")
	p.UnTab()
	p.P("}")
	return nil
}

func prepareImports(p printer.Printer, meta dsl.Meta, operation registry.UnaryOperation) {
	p.P("import (")
	p.Tab()
	p.P("\"context\"")
	p.P()
	p.P("\"github.com/pkg/errors\"")
	p.P()
	p.P("\"", meta.GoPackage, "/", operation.Context().Package(), "\"")
	p.UnTab()
	p.P(")")
}

func (o *opContext) InputDependency() registry.Message {
	return o.operation.Input()
}
