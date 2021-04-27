package main

import (
	"flag"
	"log"

	"github.com/carousell/aggproto/app"
)
var(
	specDir string
	registryDir string
)

func init() {
	flag.StringVar(&specDir,"spec","","directory for api specs")
	flag.StringVar(&registryDir,"registry","","directory with generated registry files")
	flag.Parse()
}

func main() {
	err := app.Run(specDir, registryDir)
	if err != nil {
		log.Fatal(err)
	}
}
