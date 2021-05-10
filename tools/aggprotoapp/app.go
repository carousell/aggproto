//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package main

import (
	"flag"
	"log"

	"github.com/carousell/aggproto/app"
)

var (
	specDir     string
	registryDir string
)

func init() {
	flag.StringVar(&specDir, "spec", "", "directory for api specs")
	flag.StringVar(&registryDir, "registry", "", "directory with generated registry files")
	flag.Parse()
}

func main() {
	err := app.Run(specDir, registryDir)
	if err != nil {
		log.Fatal(err)
	}
}
