//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package msgresolution

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/stage"
	"github.com/carousell/aggproto/pkg/registry"
)

type MessageResolver interface {
	Resolve(apiDescriptor dsl.ApiDescriptor, meta dsl.Meta, messageFullName []dsl.FieldDescriptor) (AdaptorContext, error)
}

type AdaptorContext interface {
	stage.Stage
	Dependencies() []registry.Message
}
