package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

var rootCmd = &cobra.Command{
	Use:  "toolset",
	Long: "tower project dev command line",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use --help for usage description.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
