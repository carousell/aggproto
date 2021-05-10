package dsl

type ArgDescriptor interface {
	Input() string
	Output() string
}

type AliasArgDescriptor struct {
	input  string
	output string
}

func (a *AliasArgDescriptor) Input() string {
	return a.input
}

func (a *AliasArgDescriptor) Output() string {
	return a.output
}

type PipedArgDescriptor struct {
	input  string
	output string
}

func (p *PipedArgDescriptor) Input() string {
	return p.input
}

func (p *PipedArgDescriptor) Output() string {
	return p.output
}
