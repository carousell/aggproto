//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package cmd

import (
	"github.com/carousell/aggproto/app"
	"github.com/spf13/cobra"
)

var (
	specDir string
	appCmd  = &cobra.Command{
		Use: "app",
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Run("", "")
		},
	}
)

func init() {
	appCmd.Flags().StringVar(&specDir, "spec", "", "Directory storing the api-specs")
	//_ = appCmd.MarkFlagRequired("spec")
	rootCmd.AddCommand(appCmd)
}
