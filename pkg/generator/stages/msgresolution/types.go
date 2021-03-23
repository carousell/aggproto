package msgresolution

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator"
	"github.com/carousell/aggproto/pkg/registry"
)

type MessageResolver interface {
	Resolve(apiDescriptor dsl.ApiDescriptor, messageFullName []dsl.FieldDescriptor) AdaptorContext
}

type AdaptorContext interface {
	generator.Context
	Dependencies() []registry.Message
}
