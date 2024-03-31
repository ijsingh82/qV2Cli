/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package inter

import (
	"fmt"

	"github.com/ijsingh82/qV2Cli/pkg/inter"
	"github.com/ijsingh82/qV2Cli/pkg/utils"

	"github.com/spf13/cobra"
)

// aprvRoleCmd represents the aprvRole command
var asgnAdminRoleKeyCmd = &cobra.Command{
	Use:              "assignAdminRoleKey",
	TraverseChildren: true,
	Short:            "Assign an address for the admin role to the specified org in the chain.",
	Long: `    /** @notice interface to assign network admin/org admin role to an account
	this can be executed by network admin accounts only
  * @param _orgId unique id of the organization to which the account belongs
  * @param _account account id
  * @param _roleId role id to be assigned to the account
  */`,
	RunE: func(cmd *cobra.Command, args []string) error {

		roleID, err := cmd.Flags().GetString("roleId")
		if err != nil {
			return err
		}

		orgId, err := cmd.Flags().GetString("orgId")
		if err != nil {
			return err
		}
		newAcct, err := cmd.Flags().GetString("newAcc")
		if err != nil {
			return err
		}

		keyfile, err := cmd.Flags().GetString("keyfile")
		if err != nil {
			return err
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			return err
		}

		r, err := inter.AssignAdminRoleKey(keyfile, password, newAcct, roleID, orgId)
		if err != nil {
			return err
		}

		fmt.Printf("This is the transaction receipt for assigning Admin role:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil
	},
}

func init() {

	asgnAdminRoleKeyCmd.PersistentFlags().StringP("roleId", "r", "", "The role id to be assigned to account.")
	asgnAdminRoleKeyCmd.PersistentFlags().StringP("keyfile", "k", "", "The authorizer keyfile.")
	asgnAdminRoleKeyCmd.PersistentFlags().StringP("password", "p", "", "The authorizer keyfile password.")
	asgnAdminRoleKeyCmd.PersistentFlags().StringP("orgId", "o", "", "The org Id the role ID belongs to in Chain.")
	asgnAdminRoleKeyCmd.PersistentFlags().StringP("newAcc", "n", "", "The new account being approved.")

	asgnAdminRoleKeyCmd.MarkPersistentFlagRequired("keyfile")
	asgnAdminRoleKeyCmd.MarkPersistentFlagRequired("newAcc")
	asgnAdminRoleKeyCmd.MarkPersistentFlagRequired("orgId")
	asgnAdminRoleKeyCmd.MarkPersistentFlagRequired("roleId")

	interCmd.AddCommand(asgnAdminRoleKeyCmd)

}
