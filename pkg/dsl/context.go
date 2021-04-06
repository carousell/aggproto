package dsl

type Context struct {
	Api       ApiDescriptor
	Output    []FieldDescriptor
	Operation OpDescriptor
}
