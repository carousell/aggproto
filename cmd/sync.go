package cmd

import (
	"github.com/carousell/aggproto/cmd/sync"
	"github.com/spf13/cobra"
)

var (
	registryDir string
	apiSpecsDir string
	protoOutDir string
	goOutDir    string
	syncCmd     = &cobra.Command{
		Use: "sync",
		RunE: func(cmd *cobra.Command, args []string) error {
			return sync.APIs(registryDir, apiSpecsDir, protoOutDir, goOutDir)
		},
	}
)

func init() {
	syncCmd.Flags().StringVar(&registryDir, "registry_path", "", "Directory storing the compiled proto descriptors")
	_ = syncCmd.MarkFlagRequired("registry_path")
	syncCmd.Flags().StringVar(&apiSpecsDir, "api_specs_path", "", "Directory storing the requests for generating APIs")
	_ = syncCmd.MarkFlagRequired("api_path")
	syncCmd.Flags().StringVar(&protoOutDir, "proto_out_path", "", "Directory to store Protobuf output from generation")
	_ = syncCmd.MarkFlagRequired("proto_out_path")
	syncCmd.Flags().StringVar(&goOutDir, "go_out_path", "", "Directory to store go output from generation")
	_ = syncCmd.MarkFlagRequired("go_out_path")
	rootCmd.AddCommand(syncCmd)
}
