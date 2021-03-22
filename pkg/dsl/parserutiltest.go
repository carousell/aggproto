package dsl

func GetFieldDescriptors(inputs ...string) []FieldDescriptor {
	return parseOutputFields(inputs...)
}
