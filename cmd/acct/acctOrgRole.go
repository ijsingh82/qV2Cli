/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package acct

import (
	"fmt"

	"github.com/ijsingh82/qV2Cli/pkg/accounts"
	"github.com/ijsingh82/qV2Cli/pkg/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// acctDetailsCmd represents the acctDetails command
var orgRoleCmd = &cobra.Command{
	Use:              "orgRole",
	TraverseChildren: true,
	Short:            "Get the Org and role for account.",
	Long: `    
/** @notice returns the account details for a given account if account is valid/active
* @param _account account id
* @return org id of the account
* @return role linked to the account
*/

qV2Cli acct orgRole [account]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		if common.IsHexAddress(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid address Specified: %s", args[0])
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		acctAdd := args[0]
		d, err := accounts.GetAccountOrgDetails(acctAdd)
		if err != nil {
			return err
		}
		fmt.Printf("The Org Role Info is :\n %+v\n", utils.PPrint(d))
		return nil
	},
}

func init() {

	acctCmd.AddCommand(orgRoleCmd)

}
