/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package acct

import (
	"fmt"
	"strconv"

	"github.com/ijsingh82/qV2Cli/pkg/accounts"
	"github.com/ijsingh82/qV2Cli/pkg/model"
	"github.com/ijsingh82/qV2Cli/pkg/utils"

	"github.com/spf13/cobra"
)

// acctDetailsCmd represents the acctDetails command
var detailsFrmIndxCmd = &cobra.Command{
	Use:              "fromIndex",
	TraverseChildren: true,
	Short:            "Get the Account Details from an index.",
	Long: `
/** @notice returns the account details a given account index
* @param  _aIndex account index
* @return account id
* @return org id of the account
* @return role linked to the account
* @return status of the account
* @return bool indicating if the account is an org admin
*/

qV2Cli acct fromIndex [index]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		_, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("not a number: %s", args[0])
		}

		return nil

	},
	RunE: func(cmd *cobra.Command, args []string) error {

		index, _ := strconv.Atoi(args[0])

		d, err := accounts.GetAccountDetailsFromIndex(int64(index))
		if err != nil {
			return err
		}
		access, err := model.RoleAccesString(d.Status)
		if err != nil {
			return err
		}

		fmt.Printf("The Accounts Details are :\n %+v\n", utils.PPrint(d))
		fmt.Printf("The the Account Status is: %v\n", access)

		return nil
	},
}

func init() {

	acctCmd.AddCommand(detailsFrmIndxCmd)

}
