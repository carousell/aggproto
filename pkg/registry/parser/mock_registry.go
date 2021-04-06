package parser

import (
	"strings"

	"github.com/carousell/aggproto/pkg/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type mockRegistry struct {
	onAddMessages   func(...*MessageContainer)
	onAddOperations func(...*UnaryOperationContainer)
	cache           map[string]registry.Message
	onListMessage   []func(options registry.ListMessageOptions) ([]registry.Message, bool)
	onListOperation []func(options registry.ListServiceOptions) ([]registry.UnaryOperation, bool)
}

func (mr *mockRegistry) save() error {
	return nil
}

func Mock() *mockRegistry {
	return &mockRegistry{cache: make(map[string]registry.Message)}
}
func (mr *mockRegistry) OnAddMessages(fn func(...*MessageContainer)) *mockRegistry {
	mr.onAddMessages = fn
	return mr
}
func (mr *mockRegistry) OnAddOperations(fn func(...*UnaryOperationContainer)) *mockRegistry {
	mr.onAddOperations = fn
	return mr
}
func (mr *mockRegistry) OnListMessageMatchPrefix(prefix string, ret []registry.Message) *mockRegistry {
	mr.onListMessage = append(mr.onListMessage, func(options registry.ListMessageOptions) (messages []registry.Message, b bool) {
		if options.PrefixMatch != nil && strings.HasPrefix(*options.PrefixMatch, prefix) {
			return ret, true
		}
		return nil, false
	})
	return mr
}
func (mr *mockRegistry) OnListOperationMatchReturn(match registry.Message, ret []registry.UnaryOperation) *mockRegistry {
	mr.onListOperation = append(mr.onListOperation, func(options registry.ListServiceOptions) (operations []registry.UnaryOperation, b bool) {
		if options.MatchReturn == match {
			return ret, true
		}
		return nil, false
	})
	return mr
}
func (mr *mockRegistry) addAll(msgs ...*MessageContainer) {
	for _, msg := range msgs {
		mr.cache[msg.FullName()] = msg
		mr.addAll(msg.MessageDefinitions...)
	}
}

func (mr *mockRegistry) addMessages(msgs ...*MessageContainer) {
	if mr.onAddMessages != nil {
		mr.onAddMessages(msgs...)
	}
	mr.addAll(msgs...)
}

func (mr *mockRegistry) addOperations(ops ...*UnaryOperationContainer) {
	if mr.onAddOperations != nil {
		mr.onAddOperations(ops...)
	}
}

func (mr *mockRegistry) addFileDescriptor(fd *descriptorpb.FileDescriptorProto) {
}

func (mr *mockRegistry) ListMessages(options ...registry.ListMessageOption) []registry.Message {
	lmo := registry.ListMessageOptions{}
	for _, op := range options {
		lmo = op(lmo)
	}
	if mr.onListMessage != nil {
		for _, fn := range mr.onListMessage {
			ret, handled := fn(lmo)
			if handled {
				return ret
			}
		}
	}
	if lmo.ExactFullName != nil {
		if msg, ok := mr.cache[*lmo.ExactFullName]; ok {
			return []registry.Message{msg}
		}
	}
	return nil
}

func (mr *mockRegistry) ListOperations(options ...registry.ListServiceOption) []registry.UnaryOperation {
	lso := registry.ListServiceOptions{}
	for _, op := range options {
		lso = op(lso)
	}
	for _, fn := range mr.onListOperation {
		ret, handled := fn(lso)
		if handled {
			return ret
		}
	}
	return nil
}
