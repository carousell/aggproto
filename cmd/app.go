package cmd

import (
	"github.com/carousell/aggproto/app"
	"github.com/spf13/cobra"
)

var (
	specDir string
	appCmd = &cobra.Command{
		Use: "app",
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Run("","")
		},
	}
)

func init() {
	appCmd.Flags().StringVar(&specDir, "spec", "","Directory storing the api-specs")
	//_ = appCmd.MarkFlagRequired("spec")
	rootCmd.AddCommand(appCmd)
}
