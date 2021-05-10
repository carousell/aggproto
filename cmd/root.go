package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{}
)

func Run() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("error executing command %v", err)
	}
}
