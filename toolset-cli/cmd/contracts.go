package cmd

import (
	"fmt"
	"git.jetbrains.space/artdecoction/wt/tower/toolset-cli/cmd/contracts"
	"github.com/spf13/cobra"
)

var serviceFlag string
var skipLintingFlag bool
var skipBreakingFlag bool

const possibleActionsInfo = "Possible actions: generate"

func init() {
	rootCmd.AddCommand(contractsCmd)
	contractsCmd.Flags().StringVarP(&serviceFlag, "service", "s", "all", "Narrows scope to specified service")
	contractsCmd.Flags().BoolVarP(&skipLintingFlag, "skip-linting", "", false, "Disables proto definition linting")
	contractsCmd.Flags().BoolVarP(&skipBreakingFlag, "skip-breaking", "", false, "Disables proto definition breaking check")
}

var contractsCmd = &cobra.Command{
	Use:   "contracts",
	Short: "Manage contracts definitions",
	Long: `Tool to manage proto contracts definitions.
Underneath buf is used. For your convenience you do not need to install any proto compilers on your own.
Instead the proto compilation is done inside docker container, so you need working docker on your machine.
` + possibleActionsInfo + "\n",
	Args: cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "generate":
			contracts.Generate(serviceFlag, skipLintingFlag, skipBreakingFlag)
		default:
			fmt.Println("Please specify valid action. " + possibleActionsInfo)
		}
	},
}
