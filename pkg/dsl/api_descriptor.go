package dsl

type ApiDescriptor interface {
	Name() string
	Version() int
	Group() string
}

type apiDescriptor struct {
	name    string
	version int
	group   string
}

func (a *apiDescriptor) Name() string {
	return a.name
}

func (a *apiDescriptor) Version() int {
	return a.version
}
func (a *apiDescriptor) Group() string {
	return a.group
}
