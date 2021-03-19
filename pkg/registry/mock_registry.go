package registry

type mockRegistry struct {
	onAddMessages   func(...Message)
	onAddOperations func(...UnaryOperation)
	cache           map[string]Message
}

func Mock() *mockRegistry {
	return &mockRegistry{cache: make(map[string]Message)}
}
func (mr *mockRegistry) OnAddMessages(fn func(...Message)) *mockRegistry {
	mr.onAddMessages = fn
	return mr
}
func (mr *mockRegistry) OnAddOperations(fn func(...UnaryOperation)) *mockRegistry {
	mr.onAddOperations = fn
	return mr
}
func (mr *mockRegistry) addAll(msgs ...Message) {
	for _, msg := range msgs {
		mr.cache[msg.FullName()] = msg
		mr.addAll(msg.Definitions()...)
	}
}

func (mr *mockRegistry) AddMessages(msgs ...Message) {
	if mr.onAddMessages != nil {
		mr.onAddMessages(msgs...)
	}
	mr.addAll(msgs...)
}

func (mr *mockRegistry) AddOperations(ops ...UnaryOperation) {
	if mr.onAddOperations != nil {
		mr.onAddOperations(ops...)
	}
}

func (mr *mockRegistry) ListMessages(options ...ListMessageOption) []Message {
	lmo := listMessageOptions{}
	for _, op := range options {
		lmo = op(lmo)
	}
	if lmo.exactFullName != nil {
		if msg, ok := mr.cache[*lmo.exactFullName]; ok {
			return []Message{msg}
		}
	}
	return nil
}

func (mr *mockRegistry) ListOperations() []UnaryOperation {
	return nil
}
