package msgresolution

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/stage"
	"github.com/carousell/aggproto/pkg/registry"
)

type MessageResolver interface {
	Resolve(apiDescriptor dsl.ApiDescriptor, meta dsl.Meta, messageFullName []dsl.FieldDescriptor) (AdaptorContext, error)
}

type AdaptorContext interface {
	stage.Stage
	Dependencies() []registry.Message
}
