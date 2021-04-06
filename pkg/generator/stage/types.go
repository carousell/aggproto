package stage

import "github.com/carousell/aggproto/pkg/generator/printer"

type Stage interface {
	PrintProto(p printer.Printer)
	PrintCode(p printer.Printer)
}
