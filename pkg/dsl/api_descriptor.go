package dsl

import "fmt"

type ApiDescriptor interface {
	Name() string
	Version() int
	Group() string
	FileName() string
}

type apiDescriptor struct {
	name    string
	version int
	group   string
}

func (a *apiDescriptor) FileName() string {
	return fmt.Sprintf("%s_v%d/%s", a.group, a.version, a.name)
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
