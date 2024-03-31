/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package acct

import (
	"github.com/ijsingh82/qV2Cli/pkg/accounts"

	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// acctValidCmd represents the acctValid command
var validCmd = &cobra.Command{
	Use:              "valid",
	TraverseChildren: true,
	Short:            "Check if an account is valid for an org on the chain.",
	Long: `  
/** @notice checks if the passed account exists and if exists does it belong to the passed organization.
* @param _account - account id
* @param _orgId - org id
* @return bool true if the account does not exists or exists and belongs
* @return passed org
*/

qV2Cli acct valid [account] [orgId]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
			return err
		}

		if common.IsHexAddress(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid address Specified: %s", args[0])
	}, RunE: func(cmd *cobra.Command, args []string) error {

		acctAdd := args[0]
		orgId := args[1]
		v, err := accounts.ValidatedAccount(acctAdd, orgId)
		if err != nil {
			return err
		}
		fmt.Println("True if the account does not exists or exists and belongs passed org :", v)

		return nil
	},
}

func init() {

	acctCmd.AddCommand(validCmd)

}
