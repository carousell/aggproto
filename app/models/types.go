//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package models

import "fmt"

type APIDescriptor struct {
	Name    string
	Group   string
	Version int
}

type SpecInfo struct {
	APIDescriptor APIDescriptor
	Filename      string
}

func (si *SpecInfo) Id() string {
	return fmt.Sprintf("%s_v%d/%s", si.APIDescriptor.Group, si.APIDescriptor.Version, si.APIDescriptor.Name)
}

type InputType int

const (
	Unknown InputType = iota
	PrimitiveIntType
	PrimitiveStringType
	PrimitiveDoubleType
	PrimitiveBoolType
	MessageType
)

type APIContextOutput struct {
	Input struct {
		Value string
		Type  InputType
	}
	Output struct {
		Value string
	}
}

type APIContext struct {
	API  APIDescriptor
	Meta struct {
		GoPackage string
	}
	Operations struct {
		AllowedOperations []string
	}
	Outputs []APIContextOutput
}

type MessageField struct {
	Name     string
	Type     string
	Message  *Message
	Selected bool
	Repeated bool
}

type Message struct {
	PackageName string
	Name        string
	Fields      []*MessageField
}
