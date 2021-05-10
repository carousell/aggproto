//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package stage

import "github.com/carousell/aggproto/pkg/generator/printer"

// todo maybe split interface
type Stage interface {
	PrintProto(p printer.Factory)
	PrintCode(p printer.Factory) error
	PrintCodeUsage(p printer.Printer)
	AddStageDependency(stage Stage)
	GetStageDependencies() []Stage
}
