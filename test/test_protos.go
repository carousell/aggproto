package test

import (
	_ "embed"
)

//go:embed testdata/pkga/a.proto.fdp
var PkgaAProto []byte
//go:embed testdata/pkgb/b.proto.fdp
var PkgbBProto []byte
