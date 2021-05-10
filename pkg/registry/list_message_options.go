//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package registry

type ListMessageOptions struct {
	ExactFullName *string
	PrefixMatch   *string
}

// does message sub definition based resolution
func LMOWithFullName(name string) ListMessageOption {
	return func(options ListMessageOptions) ListMessageOptions {
		options.ExactFullName = &name
		return options
	}
}

// used for field name resolution;
func LMOWithPrefixMatch(prefix string) ListMessageOption {
	return func(options ListMessageOptions) ListMessageOptions {
		options.PrefixMatch = &prefix
		return options
	}
}
