/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package acct

import (
	"github.com/ijsingh82/qV2Cli/pkg/accounts"

	"fmt"

	"github.com/spf13/cobra"
)

// acctDetailsCmd represents the acctDetails command
var adminExistsCmd = &cobra.Command{
	Use:              "orgAdminExists",
	TraverseChildren: true,
	Short:            "Check of org admin account exists for the org id.",
	Long: `    
/** @notice checks if org admin account exists for the passed org id
* @param _orgId - org id
* @return true if the org admin account exists and is approved

qV2Cli acct orgAdminExists [orgId]`,
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		org := args[0]
		d, err := accounts.OrgAdminExists(org)
		if err != nil {
			return err
		}
		fmt.Printf("The org Admin exists : %v\n", d)

		return nil
	},
}

func init() {

	acctCmd.AddCommand(adminExistsCmd)

}
