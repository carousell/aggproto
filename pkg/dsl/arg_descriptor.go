//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
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
