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
var roleAccess = &cobra.Command{
	Use:              "access",
	TraverseChildren: true,
	Short:            "Get the role Access level.",
	Long: `    
/** @notice checks if the role access level if it exists for the given org or master org
* @param _roleId - unique identifier for the role being added
* @param _orgId - org id to which the role belongs
* @param _ultParent - master org id
* @return true or false

qV2Cli role access [roleId] [orgId] [ultParent]`,
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

		acc, err := role.Access(roleId, orgId, ultPrnt)
		if err != nil {
			return err
		}
		fmt.Printf("The Account access for RoleId %v, OrgId %v, Ultimate Parent %v, %v\n", roleId, orgId, ultPrnt, acc)

		return nil
	},
}

func init() {

	roleCmd.AddCommand(roleAccess)

}
