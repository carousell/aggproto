package registry

type Message interface {
	Fields() []Field
	Definitions() []Message
	Parent() Message
	Package() string
	Name() string
	FullName()string
}

type FieldType int

const (
	FieldTypeInt32 FieldType = iota
	FieldTypeInt64
	FieldTypeBool
	FieldTypeDouble
	FieldTypeString
	FieldTypeMessage
	FieldTypeEnum
)

type Field interface {
	Type() FieldType
	Name() string
	Message() Message
	Context() Message
	Repeated() bool
}
type UnaryOperation interface {
	Input() Message
	Output() Message
	Context() Service
}
type Service interface {
	UnaryOperations() []UnaryOperation
	Package() string
}

type ListMessageOption func(options listMessageOptions) listMessageOptions

type Registry interface {
	AddMessages(...Message)
	AddOperations(...UnaryOperation)
	ListMessages(...ListMessageOption) []Message
	ListOperations() []UnaryOperation
}
