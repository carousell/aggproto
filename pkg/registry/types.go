package registry

type Message interface {
	Fields() []Field
	Definitions() []Message
	Parent() Message
	Package() string
	Name() string
	FullName() string
}

type FieldType int

func (ft FieldType) GoTypeString() string {
	switch ft {
	case FieldTypeInt64:
		return "int64"
	case FieldTypeBool:
		return "bool"
	case FieldTypeDouble:
		return "float64"
	case FieldTypeString:
		return "string"
	case FieldTypeEnum:
		return "int"
	default:
		panic("unimplemented ")
	}
}

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
	Name() string
	Context() Service
	FullName() string
}
type Service interface {
	UnaryOperations() []UnaryOperation
	Package() string
	Name() string
}

type ListMessageOption func(options ListMessageOptions) ListMessageOptions

type ListServiceOption func(options ListServiceOptions) ListServiceOptions

type Registry interface {
	ListMessages(...ListMessageOption) []Message
	ListOperations(...ListServiceOption) []UnaryOperation
}
