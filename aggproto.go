package main

import (
	"github.com/carousell/aggproto/cmd"
	"os"
)

func main() {
	os.Args = append(os.Args, "sync")
	cmd.Run()
}
