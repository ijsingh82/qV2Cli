/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package depl

import (
	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/spf13/cobra"
)

// acctCmd represents the acct command
var deplCmd = &cobra.Command{
	Use:              "depl",
	TraverseChildren: true,
	Short:            "The functions to deploy the permission chain contracts.",
	Long:             `The functions to deploy the permission chain contracts.`,
	Args:             cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

func init() {
	cmd.RootCmd.AddCommand(deplCmd)

}
