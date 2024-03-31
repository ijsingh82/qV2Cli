/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package acct

import (
	"fmt"

	"github.com/ijsingh82/qV2Cli/pkg/accounts"
	"github.com/ijsingh82/qV2Cli/pkg/model"
	"github.com/ijsingh82/qV2Cli/pkg/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// acctDetailsCmd represents the acctDetails command
var detailsCmd = &cobra.Command{
	Use:              "detail",
	TraverseChildren: true,
	Short:            "Get the Account Details from an account.",
	Long: `
/** @notice returns the account details for a given account
* @param _account account id
* @return account id
* @return org id of the account
* @return role linked to the account
* @return status of the account
* @return bool indicating if the account is an org admin
*/

qV2Cli acct detail [account]`,
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
		d, err := accounts.GetAccountDetails(acctAdd)
		if err != nil {
			return err
		}

		access, err := model.RoleAccesString(d.Status)
		if err != nil {
			return err
		}

		fmt.Printf("The Accounts Details are :\n %+v\n", utils.PPrint(d))
		fmt.Printf("The the Account Status : %v\n", access)

		return nil
	},
}

func init() {

	acctCmd.AddCommand(detailsCmd)

}
