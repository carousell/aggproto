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
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/generator/stage"
	"github.com/carousell/aggproto/pkg/generator/stages/msgresolution"
	"github.com/carousell/aggproto/pkg/registry"
)

type OperationResolver interface {
	Resolve(api dsl.ApiDescriptor, meta dsl.Meta, ctx msgresolution.AdaptorContext, descriptor dsl.OpDescriptor) []OperationContext
}

type OperationContext interface {
	stage.Stage
	InputDependency() registry.Message
	// splitting the interface ?
	ClientReferenceName() string
	RequiredImport() string
	ClientDependency() string
	PrintConstructorUsage(p printer.Printer)
	Produces() registry.Message
	Consumes() registry.Message
}
