package dsl

type FieldDescriptor interface {
	Input() InputFieldDescriptor
	Output() OutputFieldDescriptor
}

type InputFieldDescriptor interface {
	isInputFieldDescriptor()
}
type OutputFieldDescriptor interface {
	isOutputFieldDescriptor()
}
type NamespacedMessageFieldDescriptor struct {
	NamespacedField string
}

func (n *NamespacedMessageFieldDescriptor) isOutputFieldDescriptor() {
}

func (n *NamespacedMessageFieldDescriptor) isInputFieldDescriptor() {
}

type PrimitiveFieldDescriptor interface {
	InputFieldDescriptor
	isPrimitiveFieldDescriptor()
}

type StringValueFieldDescriptor struct {
	Value string
}

func (s *StringValueFieldDescriptor) isInputFieldDescriptor() {
}

func (s *StringValueFieldDescriptor) isPrimitiveFieldDescriptor() {
}

type IntValueFieldDescriptor struct {
	Value int64
}

func (i *IntValueFieldDescriptor) isInputFieldDescriptor() {
}

func (i *IntValueFieldDescriptor) isPrimitiveFieldDescriptor() {
}

type BoolValueFieldDescriptor struct {
	Value bool
}

func (b *BoolValueFieldDescriptor) isInputFieldDescriptor() {
}

func (b *BoolValueFieldDescriptor) isPrimitiveFieldDescriptor() {
}

type FloatValueFieldDescriptor struct {
	Value float64
}

func (f *FloatValueFieldDescriptor) isInputFieldDescriptor() {
}

func (f *FloatValueFieldDescriptor) isPrimitiveFieldDescriptor() {
}
