package msgresolution

import (
	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/registry"
)

type MessageResolver interface {
	Resolve(messageFullName []dsl.FieldDescriptor) AdaptorContext
}

type AdaptorContext interface {
	Dependencies() []registry.Message
}
