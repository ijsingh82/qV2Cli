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
var implOrgAdminCmd = &cobra.Command{
	Use:              "isOrgAdmin",
	TraverseChildren: true,
	Short:            "Check if account is a org admin.",
	Long: `    
/** @notice function to check if passed account is an org admin account
* @param _account account id
* @param _orgId organization id
* @return true/false
*/

qV2Cli impl isOrgAdmin [account] [orgId]`,

	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
			return err
		}

		if common.IsHexAddress(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid address Specified: %s", args[0])

	},
	RunE: func(cmd *cobra.Command, args []string) error {

		acctAdd := args[0]
		orgId := args[1]
		orgAdmin, err := impl.IsOrgAdmin(acctAdd, orgId)
		if err != nil {
			return err
		}
		fmt.Printf("The Account %v is Org Admin for org %v, %v \n", acctAdd, orgId, orgAdmin)
		return nil
	},
}

func init() {

	implCmd.AddCommand(implOrgAdminCmd)

}
