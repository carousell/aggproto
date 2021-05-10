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

import (
	"strings"

	"github.com/pkg/errors"
)

func parseArgDescriptors(lines ...string) ([]ArgDescriptor, error) {
	var ret []ArgDescriptor
	for _, l := range lines {
		if strings.Contains(l, "=") {
			aliasDesc, err := parseAliasArgDescriptors(l)
			if err != nil {
				return nil, err
			}
			ret = append(ret, aliasDesc)
		} else if strings.Contains(l, "<-") {
			pipedDesc, err := parsePipedArgDescriptor(l)
			if err != nil {
				return nil, err
			}
			ret = append(ret, pipedDesc)
		}
	}
	return ret, nil
}

func parsePipedArgDescriptor(l string) (*PipedArgDescriptor, error) {
	splits := strings.Split(l, "<-")
	if len(splits) != 2 {
		return nil, errors.Errorf("piped arg descriptors must have one <- operator got %d for %s", len(splits), l)
	}
	return &PipedArgDescriptor{input: splits[1], output: splits[0]}, nil
}

func parseAliasArgDescriptors(l string) (*AliasArgDescriptor, error) {
	splits := strings.Split(l, "=")
	if len(splits) != 2 {
		return nil, errors.Errorf("alias arg descriptors must have one = operator got %d for  %s", len(splits), l)
	}
	return &AliasArgDescriptor{input: splits[0], output: splits[1]}, nil
}
