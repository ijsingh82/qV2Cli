/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package acct

import (
	"fmt"

	"github.com/ijsingh82/qV2Cli/pkg/accounts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// acctRoleCmd represents the acctRole command
var acctroleCmd = &cobra.Command{
	Use:              "role",
	TraverseChildren: true,
	Short:            "Get the role for the account on chain.",
	Long: `    
/** @notice returns the role id linked to the passed account
* @param _account account id
* @return role id
*/

qV2Cli acct role [account]`,
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

		role, err := accounts.GetRoleForAccount(acctAdd)
		if err != nil {
			return err
		}
		fmt.Printf("The Accounts role is : %#v\n", role)

		return nil
	},
}

func init() {

	acctCmd.AddCommand(acctroleCmd)

}
