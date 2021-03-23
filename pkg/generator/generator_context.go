package generator

import "github.com/carousell/aggproto/pkg/generator/printer"

type Context interface {
	PrintProto(p printer.Printer)
	PrintCode(p printer.Printer)
}
