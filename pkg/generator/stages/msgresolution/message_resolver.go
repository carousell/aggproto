package msgresolution

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/registry"
)

type msgResolver struct {
	r registry.Registry
}

type adaptorBuilder struct {
	adaptorUnits []adaptorUnit
}

func (n *nestedAdaptorUnit) isAdaptorUnit() {
	panic("Should never be called")
}

func (m *msgResolver) Resolve(fds []dsl.FieldDescriptor) AdaptorContext {
	var adaptorUnits adaptorUnits
	for _, fd := range fds {
		ofd, ok := fd.Output().(*dsl.NamespacedMessageFieldDescriptor)
		if !ok {
			panic("unsupported output field descriptor")
		}
		switch ifd := fd.Input().(type) {
		case dsl.PrimitiveFieldDescriptor:
			adaptorUnits = append(adaptorUnits, makeStaticPrimitiveAdaptorUnit(ofd, ifd))

		case *dsl.NamespacedMessageFieldDescriptor:
		default:
			panic("unsupported input field descriptor")
		}
	}
	return &adaptorContext{adaptorUnits: adaptorUnits.tryMerge()}
}
