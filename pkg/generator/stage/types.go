package stage

import "github.com/carousell/aggproto/pkg/generator/printer"

// todo maybe split interface
type Stage interface {
	PrintProto(p printer.Factory)
	PrintCode(p printer.Factory)
	PrintCodeUsage(p printer.Printer)
	AddStageDependency(stage Stage)
	GetStageDependencies() []Stage
}
