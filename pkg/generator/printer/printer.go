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

func New() Printer {
	return &printer{new(bytes.Buffer), ""}
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
	default:
		log.Printf("unknown type in printer %T\n", v)
	}
}
