package dsl

func GetFieldDescriptors(inputs ...string) []FieldDescriptor {
	return parseOutputFields(inputs...)
}

func GetApiDescriptor(group, name string, version int) ApiDescriptor {
	return &apiDescriptor{name, version, group}
}
