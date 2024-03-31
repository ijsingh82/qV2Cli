/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package impl

import (
	"fmt"

	"github.com/ijsingh82/qV2Cli/pkg/impl"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// acctValidCmd represents the acctValid command
var implValidCmd = &cobra.Command{
	Use:              "valid",
	TraverseChildren: true,
	Short:            "Check if an account is valid for an org on the chain.",
	Long: `
/** @notice function to validate the account for access change operation
* @param _account account id
* @param _orgId organization id
* @return true/false
*/

qV2Cli impl valid [account] [orgId]`,
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
		v, err := impl.ValidateAccount(acctAdd, orgId)
		if err != nil {
			return err
		}
		fmt.Println("True if the account does not exists or exists and belongs passed org :", v)

		return nil
	},
}

func init() {

	implCmd.AddCommand(implValidCmd)

}
