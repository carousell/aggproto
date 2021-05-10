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
