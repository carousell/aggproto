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

type ApiSpec struct {
	Api struct {
		Name    string
		Group   string
		Version int
	}
	Meta struct {
		GoPackage string
	}
	Input      []string
	Output     []string
	Operations []string
}

func (s *ApiSpec) ToContext() (*Context, error) {
	argDesc, err := parseArgDescriptors(s.Input...)
	if err != nil {
		return nil, err
	}
	return &Context{
		Api: &apiDescriptor{
			name:    s.Api.Name,
			group:   s.Api.Group,
			version: s.Api.Version,
		},
		Operation: OpDescriptor{AllowedOperations: s.Operations},
		Output:    parseOutputFields(s.Output...),
		Meta:      Meta{GoPackage: s.Meta.GoPackage},
		Input:     argDesc,
	}, nil
}
