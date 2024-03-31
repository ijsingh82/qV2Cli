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

// checkAdminOrgCmd Check if an account is an admin for the passed org or ultimate parent.
var checkAdminOrgCmd = &cobra.Command{
	Use:              "checkOrgAdmin",
	TraverseChildren: true,
	Short:            "Check if an account is an admin for the passed org or ultimate parent. ",
	Long: `
/** @notice checks if the account is a org admin for the passed org or
for the ultimate parent organization 
NOTE: "" for no ultParent.
* @param _account account id
* @param _orgId org id
* @param _ultParent master org id or
*/  

qV2Cli acct checkOrgAdmin [address] [orgId] [ultParent]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(3)(cmd, args); err != nil {
			return fmt.Errorf("add args %s", err.Error())
		}
		if common.IsHexAddress(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid address Specified: %s", args[0])

	},
	RunE: func(cmd *cobra.Command, args []string) error {

		acct := args[0]
		org := args[1]
		ultId := args[2]

		stat, err := accounts.CheckOrgAdmin(acct, org, ultId)
		if err != nil {
			return err
		}
		fmt.Printf("The Account is an Org Admin for OrgId %v : %v\n", org, stat)
		return nil
	},
}

func init() {

	acctCmd.AddCommand(checkAdminOrgCmd)
}
