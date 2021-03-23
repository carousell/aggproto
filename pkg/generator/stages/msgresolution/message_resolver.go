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

func (m *msgResolver) Resolve(api dsl.ApiDescriptor, fds []dsl.FieldDescriptor) AdaptorContext {
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
			adaptorUnits = append(adaptorUnits, makeNamespacedMessageAdaptorUnit(m.r, ofd, ifd))
		default:
			panic("unsupported input field descriptor")
		}
	}
	return &adaptorContext{apiDescriptor: api, adaptorUnits: adaptorUnits.tryMerge()}
}
