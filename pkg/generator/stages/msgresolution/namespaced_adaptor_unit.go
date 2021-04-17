package msgresolution

import (
	"fmt"
	"strings"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/carousell/aggproto/pkg/registry"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
)

type messageFieldAdaptorUnit struct {
	underlying               registry.Field
	fieldName                string
	fieldMessageDependencies []fieldMessageDependency
	repeated                 bool
}

func (m *messageFieldAdaptorUnit) isAdaptorUnit() {
	panic("should never be called")
}

func (m *messageFieldAdaptorUnit) printProtoDefinitions(p printer.Printer) {
}

func (m *messageFieldAdaptorUnit) printAsProtoField(p printer.Printer, idx int) {
	switch m.underlying.Type() {
	case registry.FieldTypeString:
		p.P("string ", m.fieldName, " = ", idx, ";")
	case registry.FieldTypeBool:
		p.P("bool ", m.fieldName, " = ", idx, ";")
	case registry.FieldTypeInt64:
		p.P("int64 ", m.fieldName, " = ", idx, ";")
	default:
		panic("unhandled field type")
	}
}

func (m *messageFieldAdaptorUnit) printAsAdaptorCode(p printer.Printer, referenceName string, parents []string) {
	fieldName := strcase.ToLowerCamel(m.fieldMessageDependencies[len(m.fieldMessageDependencies)-1].fieldName)
	p.P(referenceName, ".", strcase.ToCamel(m.fieldName), " = ", fieldName, ".", strcase.ToCamel(m.underlying.Name()))
}

func (m *messageFieldAdaptorUnit) dependencies() [][]fieldMessageDependency {
	return [][]fieldMessageDependency{m.fieldMessageDependencies}
}

func makeNamespacedMessageAdaptorUnit(r registry.Registry, ofd *dsl.NamespacedMessageFieldDescriptor, ifd *dsl.NamespacedMessageFieldDescriptor) (adaptorUnit, error) {
	msgs := r.ListMessages(registry.LMOWithPrefixMatch(ifd.NamespacedField))
	if len(msgs) != 1 {
		panic(fmt.Sprintf("No or more than one messages retrieved for %s", ifd.NamespacedField))
	}
	msg := msgs[0]
	ifdSplits := strings.Split(ifd.NamespacedField, ".")
	ofdSplits := strings.Split(ofd.NamespacedField, ".")

	dependencyStack := []fieldMessageDependency{{msg.Name(), msg, false}}
	var resolvedField registry.Field
	resolvedMsg := msg
	for idx, fd := range ifdSplits {
		repeated := false
		if strings.HasSuffix(fd, "[]") {
			repeated = true
			fd = strings.Trim(fd, "[]")
		}
		found := false
		// todo get unresolved tail from somewhere
		if idx == 0 && resolvedMsg.Package() == fd {
			continue
		}
		if idx <= 1 && resolvedMsg.Name() == fd {
			continue
		}
		for _, field := range resolvedMsg.Fields() {
			if field.Name() == fd {
				found = true
				if (repeated || field.Repeated()) && !(repeated && field.Repeated()) {
					return nil, errors.Errorf("repeated fields must be marked in ifd %s", ifd.NamespacedField)
				}
				if field.Type() != registry.FieldTypeMessage {
					if idx != len(ifdSplits)-1 {
						panic("partially resolved to a non message")
					}
					resolvedField = field
				} else {
					resolvedMsg = field.Message()
					dependencyStack = append(dependencyStack, fieldMessageDependency{fieldName: field.Name(), message: resolvedMsg, repeated: field.Repeated()})
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
		lastOfdSplit := ofdSplits[len(ofdSplits)-1]
		repeated := false
		if strings.HasSuffix(lastOfdSplit, "[]") {
			repeated = true
			lastOfdSplit = strings.Trim(lastOfdSplit, "[]")
		}
		unit = &messageFieldAdaptorUnit{underlying: resolvedField, fieldName: lastOfdSplit, fieldMessageDependencies: dependencyStack, repeated: repeated}
		for i := len(ofdSplits) - 2; i >= 0; i -= 1 {
			currSplit := ofdSplits[i]
			repeated := false
			if strings.HasSuffix(currSplit, "[]") {
				repeated = true
				currSplit = strings.Trim(currSplit, "[]")
			}
			unit = &nestedAdaptorUnit{fieldName: currSplit, nestedUnit: []adaptorUnit{unit}, repeated: repeated}
		}
		return unit, nil
	}
	// todo implement explode message
	panic("unhandled case")
}
