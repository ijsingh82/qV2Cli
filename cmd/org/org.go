/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package cmd

import (
	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/spf13/cobra"
)

// acctCmd represents the acct command
var orgCmd = &cobra.Command{
	Use:              "org",
	TraverseChildren: true,
	Short:            "The functions for the Org contract.",
	Long:             `The functions for the Org contract.`,
	Args:             cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

func init() {
	cmd.RootCmd.AddCommand(orgCmd)

}
