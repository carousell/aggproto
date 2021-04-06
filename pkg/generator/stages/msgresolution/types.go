package msgresolution

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/stage"
	"github.com/carousell/aggproto/pkg/registry"
)

type MessageResolver interface {
	Resolve(apiDescriptor dsl.ApiDescriptor, messageFullName []dsl.FieldDescriptor) AdaptorContext
}

type AdaptorContext interface {
	stage.Stage
	Dependencies() []registry.Message
}
