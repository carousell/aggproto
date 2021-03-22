package msgresolution

import (
	"strings"

	"github.com/carousell/aggproto/pkg/dsl"
)

type staticPrimitiveAdaptorUnit struct {
	fieldName   string
	primitiveFD dsl.PrimitiveFieldDescriptor
}

func (s *staticPrimitiveAdaptorUnit) isAdaptorUnit() {
}

func makeStaticPrimitiveAdaptorUnit(ofd *dsl.NamespacedMessageFieldDescriptor, ifd dsl.PrimitiveFieldDescriptor) adaptorUnit {
	splits := strings.Split(ofd.NamespacedField, ".")
	var unit adaptorUnit
	for i := len(splits) - 1; i >= 0; i -= 1 {
		if unit == nil {
			unit = &staticPrimitiveAdaptorUnit{fieldName: splits[i], primitiveFD: ifd}
		} else {
			unit = &nestedAdaptorUnit{name: splits[i], nestedUnit: []adaptorUnit{unit}}
		}
	}
	return unit
}
