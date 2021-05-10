//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package printer

type factory struct {
	state map[string]Printer
}

func (f *factory) Out() map[string]string {
	out := map[string]string{}
	for k, v := range f.state {
		out[k] = v.String()
	}
	return out
}

type Factory interface {
	Get(printerName string) Printer
	Out() map[string]string
}

func New() *factory {
	return &factory{state: map[string]Printer{}}
}
func (f *factory) Get(printerName string) Printer {
	if p, ok := f.state[printerName]; ok {
		return p
	}
	f.state[printerName] = newPrinter()
	return f.state[printerName]
}
