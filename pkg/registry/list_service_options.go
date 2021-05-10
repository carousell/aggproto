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

type ListServiceOptions struct {
	MatchReturn Message
	OpNameIn    []string
}

func LSOWithOutputMessage(msg Message) ListServiceOption {
	return func(options ListServiceOptions) ListServiceOptions {
		options.MatchReturn = msg
		return options
	}
}

func LSOWithOperationNameIn(names []string) ListServiceOption {
	return func(options ListServiceOptions) ListServiceOptions {
		options.OpNameIn = names
		return options
	}
}
