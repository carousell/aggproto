package msgresolution

import (
	"fmt"
	"strings"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/registry"
	"github.com/iancoleman/strcase"
)

type messageFieldAdaptorUnit struct {
	underlying               registry.Field
	fieldName                string
	fieldMessageDependencies []fieldMessageDependency
}

func (m *messageFieldAdaptorUnit) isAdaptorUnit() {
	panic("should never be called")
}

func (m *messageFieldAdaptorUnit) printProtoDefinitions(p printer.Printer, fieldIdx int) {
}

func (m *messageFieldAdaptorUnit) printAsProtoField(p printer.Printer, idx int) {
	switch m.underlying.Type() {
	case registry.FieldTypeString:
		p.P("string ", m.fieldName, " = ", idx, ";")
	case registry.FieldTypeBool:
		p.P("bool ", m.fieldName, " = ", idx, ";")
	default:
		panic("unhandled field type")
	}
}

func (m *messageFieldAdaptorUnit) printAsAdaptorCode(p printer.Printer, referenceName string) {
	fieldName := strcase.ToLowerCamel(m.fieldMessageDependencies[len(m.fieldMessageDependencies)-1].fieldName)
	p.P(referenceName, ".", strcase.ToCamel(m.fieldName), " = ", fieldName, ".", strcase.ToCamel(m.fieldName))
}

func (m *messageFieldAdaptorUnit) dependencies() [][]fieldMessageDependency {
	return [][]fieldMessageDependency{m.fieldMessageDependencies}
}

func makeNamespacedMessageAdaptorUnit(r registry.Registry, ofd *dsl.NamespacedMessageFieldDescriptor, ifd *dsl.NamespacedMessageFieldDescriptor) adaptorUnit {
	msgs := r.ListMessages(registry.LMOWithPrefixMatch(ifd.NamespacedField))
	if len(msgs) != 1 {
		panic(fmt.Sprintf("No or more than one messages retrieved for %s", ifd.NamespacedField))
	}
	msg := msgs[0]
	ifdSplits := strings.Split(ifd.NamespacedField, ".")
	ofdSplits := strings.Split(ofd.NamespacedField, ".")

	dependencyStack := []fieldMessageDependency{{msg.Name(), msg}}
	var resolvedField registry.Field
	resolvedMsg := msg
	for idx, fd := range ifdSplits {
		found := false
		if resolvedMsg.Package() == fd {
			continue
		}
		if resolvedMsg.Name() == fd {
			continue
		}
		for _, field := range resolvedMsg.Fields() {
			if field.Name() == fd {
				found = true
				if field.Type() != registry.FieldTypeMessage {
					if idx != len(ifdSplits)-1 {
						panic("partially resolved to a non message")
					}
					resolvedField = field
				} else {
					resolvedMsg = field.Message()
					dependencyStack = append(dependencyStack, fieldMessageDependency{fieldName: field.Name(), message: resolvedMsg})
				}
				break
			}
		}
		if !found {
			panic("did not resolve to a known field")
		}
	}
	var unit adaptorUnit
	if resolvedField != nil {
		unit = &messageFieldAdaptorUnit{underlying: resolvedField, fieldName: ofdSplits[len(ofdSplits)-1], fieldMessageDependencies: dependencyStack}
		for i := len(ofdSplits) - 2; i >= 0; i -= 1 {
			unit = &nestedAdaptorUnit{fieldName: ofdSplits[i], nestedUnit: []adaptorUnit{unit}}
		}
		return unit
	}
	// todo implement explode message
	panic("unhandled case")
}
