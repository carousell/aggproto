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
)

type InputResolver interface {
	Resolve(api dsl.ApiDescriptor, meta dsl.Meta, input []dsl.ArgDescriptor, opCtxs []opresolution.OperationContext, context msgresolution.AdaptorContext) (*InputContext, error)
}
