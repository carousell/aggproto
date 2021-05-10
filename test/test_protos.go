//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package test

import (
	_ "embed"
)

//go:embed testdata/pkga/a.proto.fdp
var PkgaAProto []byte

//go:embed testdata/pkgb/b.proto.fdp
var PkgbBProto []byte
