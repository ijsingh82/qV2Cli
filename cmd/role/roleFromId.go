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
var roleInfoId = &cobra.Command{
	Use:              "fromId",
	TraverseChildren: true,
	Args:             cobra.MatchAll(cobra.MinimumNArgs(2), cobra.OnlyValidArgs),
	Short:            "get role details from role id",
	Long: `    
/** @notice returns the role details for a passed role id and org
* @param _roleId - unique identifier for the role being added
* @param _orgId - org id to which the role belongs
* @return role id
* @return org id
* @return access type
* @return bool to indicate if the role is a voter role
* @return bool to indicate if the role is active
*/

qV2Cli role fromId [roleId] [orgId]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		roleId := args[0]
		orgId := args[1]

		d, stat, err := role.FromId(roleId, orgId)
		if err != nil {
			return err
		}
		fmt.Printf("The Role Details are:\n%+v\n", d)
		fmt.Printf("The Role status is: %+v\n", stat)
		return nil
	},
}

func init() {

	roleCmd.AddCommand(roleInfoId)

}
