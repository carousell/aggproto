package stage

import "github.com/carousell/aggproto/pkg/generator/printer"

type Stage interface {
	PrintProto(p printer.Factory)
	PrintCode(p printer.Factory)
}
