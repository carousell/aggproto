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

import (
	"bytes"
	"fmt"
	"log"
)

type Printer interface {
	P(v ...interface{})
	String() string
	Tab()
	UnTab()
}

func newPrinter() Printer {
	p := &printer{new(bytes.Buffer), ""}
	return p
}

type printer struct {
	*bytes.Buffer
	indent string
}

func (p *printer) Tab() {
	p.indent += "\t"
}
func (p *printer) UnTab() {
	if p.indent != "" {
		p.indent = p.indent[1:]
	}
}

func (p *printer) P(v ...interface{}) {
	if len(v) == 0 {
		p.WriteByte('\n')
	}
	p.WriteString(p.indent)
	for _, atom := range v {
		p.writeAtom(atom)
	}
	p.WriteByte('\n')
}

func (p *printer) writeAtom(v interface{}) {
	switch v := v.(type) {
	case string:
		p.WriteString(v)
	case *string:
		p.WriteString(*v)
	case int:
		fmt.Fprint(p, v)
	case int64:
		fmt.Fprint(p, v)
	default:
		log.Printf("unknown type in printer %T\n", v)
	}
}
