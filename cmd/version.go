package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version  = "1.0.2"
	codename = "NodeX"
	intro    = "Backend For Panel VPN"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print current version of NodeX",
		Run: func(cmd *cobra.Command, args []string) {
			showVersion()
		},
	})
}

func showVersion() {
	fmt.Printf("%s %s (%s) \n", codename, version, intro)
}
