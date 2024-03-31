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

// acctStatCmd represents the acctStat command
// Account Status types
// 0 - Not in list
// 1 - Account pending approval
// 2 - Active
// 3 - Inactive
// 4 - Suspended
// 5 - Blacklisted
// 6 - Revoked
// 7 - Recovery Initiated for blacklisted accounts and pending approval
//
//	from network admins
var statCmd = &cobra.Command{
	Use:              "stat",
	TraverseChildren: true,
	Short:            "Check an accounts status on the chain.",
	Long: `
/** @notice returns the account status for a given account
* @param _account account id
* @return account status
*/

qV2Cli acct stat [account]`,
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

		stat, err := accounts.GetStatusOfAccount(acctAdd)
		if err != nil {
			return err
		}
		fmt.Printf("The Accounts Status is : %#v\n", stat.Status)

		return nil
	},
}

func init() {

	acctCmd.AddCommand(statCmd)

}
