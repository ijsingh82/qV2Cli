/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package impl

import (
	"github.com/ijsingh82/qV2Cli/pkg/impl"

	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// getOrgsCmd represents the getOrgs command
var implNetAdminCmd = &cobra.Command{
	Use:              "isNetAdmin",
	TraverseChildren: true,
	Short:            "Check if account is a network admin.",
	Long: `
/** @notice function to check if passed account is an network admin account
* @param _account account id
* @return true/false
*/

qV2Cli impl isNetAdmin [account]`,
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
		netAdmin, err := impl.IsNetAdmin(acctAdd)
		if err != nil {
			return err
		}
		fmt.Printf("The Account %v, is network Admin %v\n", acctAdd, netAdmin)

		return nil
	},
}

func init() {

	implCmd.AddCommand(implNetAdminCmd)

}
