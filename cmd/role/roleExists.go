/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package role

import (
	"github.com/ijsingh82/qV2Cli/pkg/role"

	"fmt"

	"github.com/spf13/cobra"
)

// acctDetailsCmd represents the acctDetails command
var roleExists = &cobra.Command{
	Use:              "exists",
	TraverseChildren: true,
	Short:            "check if the role exists.",
	Long: `
/** @notice checks if the role exists for the given org or master org
* @param _roleId - unique identifier for the role being added
* @param _orgId - org id to which the role belongs
* @param _ultParent - master org id
* @return true or false
*/

qV2Cli role exists [roleId] [orgId] [ultParent]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(3)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		roleId := args[0]
		orgId := args[1]
		ultPrnt := args[2]

		exsits, err := role.Exists(roleId, orgId, ultPrnt)
		if err != nil {
			return err
		}
		fmt.Printf("Does The Role Exists for RoleId %v, OrgId %v, Ultimate Parent %v, %v\n", roleId, orgId, ultPrnt, exsits)

		return nil
	},
}

func init() {
	roleCmd.AddCommand(roleExists)

}
