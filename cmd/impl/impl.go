/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package impl

import (
	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/spf13/cobra"
)

// acctCmd represents the acct command
var implCmd = &cobra.Command{
	Use:              "impl",
	TraverseChildren: true,
	Short:            "The functions for the Implementation contract.",
	Long:             `The functions for the Implementation contract.`,
	Args:             cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

func init() {
	cmd.RootCmd.AddCommand(implCmd)

}
