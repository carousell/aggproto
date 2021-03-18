package msgresolution

import "github.com/carousell/aggproto/pkg/registry"

type MessageResolver interface {
	Resolve(messageFullName ...string) AdaptorContext
}
type AdaptorInputContext interface {
	isAdaptorInputContext()
}
type StaticAdaptorInputContext interface {
	AdaptorInputContext
}
type MessageAdaptorInputContext interface {
	AdaptorInputContext
	Message() registry.Message
}
type RepeatedAdaptorInputContext interface {
	AdaptorInputContext
	Adaptor() AdaptorInputContext
}

type AdaptorContext interface {
	Inputs() []AdaptorInputContext
	Output() registry.Message
}
