/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package role

import (
	"strconv"

	"github.com/ijsingh82/qV2Cli/pkg/role"

	"fmt"

	"github.com/spf13/cobra"
)

// acctDetailsCmd represents the acctDetails command
var roleTxAllowed = &cobra.Command{
	Use:              "txAllow",
	TraverseChildren: true,
	Short:            "Check if transaction is allowed.",
	Long: `    
/** @notice checks if the role can perform the transactin
* @param _roleId - unique identifier for the role being added
* @param _orgId - org id to which the role belongs
* @param _ultParent - master org id
* @param _txType - type of transaction
* @return true or false
*/
NOTE : Ultimate Parent can be "".

qV2Cli role txAllow [roleId] [orgId] [ultParent] [_txType]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(4)(cmd, args); err != nil {
			return err
		}
		_, err := strconv.Atoi(args[3])
		if err != nil {
			return fmt.Errorf("not a number: %s", args[3])
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		roleId := args[0]
		orgId := args[1]
		ultPrnt := args[2]
		txType, _ := strconv.Atoi(args[3])

		allowed, err := role.TxAllowed(roleId, orgId, ultPrnt, int64(txType))
		if err != nil {
			return err
		}
		fmt.Printf("The Transaction for RoleId %v, OrgId %v, Ultimate Parent %v, and Transaction Type %v, is :  %v\n", roleId, orgId, ultPrnt, txType, allowed)
		return nil

	},
}

func init() {

	roleCmd.AddCommand(roleTxAllowed)

}
