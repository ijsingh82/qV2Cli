/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package inter

import (
	"fmt"

	"github.com/ijsingh82/qV2Cli/pkg/inter"

	"github.com/spf13/cobra"
)

// aprvRoleCmd represents the aprvRole command
var asgnAdminRoleCmd = &cobra.Command{
	Use:              "assignAdminRole",
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

		guardAdd, err := cmd.Flags().GetString("guardAdd")
		if err != nil {
			return err
		}
		guardKey, err := cmd.Flags().GetString("guardKey")
		if err != nil {
			return err
		}

		r, err := inter.AssignAdminRole(guardKey, guardAdd, newAcct, roleID, orgId)
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for assigning Admin role: %#v\n", r)

		fmt.Println()
		return nil
	},
}

func init() {

	asgnAdminRoleCmd.PersistentFlags().StringP("roleId", "i", "", "The role id to be assigned to admin account.")
	asgnAdminRoleCmd.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied for Chain.")
	asgnAdminRoleCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied for Chain.")
	asgnAdminRoleCmd.PersistentFlags().StringP("orgId", "o", "", "The org Id the role ID belongs to in Chain.")
	asgnAdminRoleCmd.PersistentFlags().StringP("newAcc", "n", "", "The new account being approved.")

	asgnAdminRoleCmd.MarkPersistentFlagRequired("guardAdd")
	asgnAdminRoleCmd.MarkPersistentFlagRequired("newAcc")
	asgnAdminRoleCmd.MarkPersistentFlagRequired("orgId")
	asgnAdminRoleCmd.MarkPersistentFlagRequired("guardKey")
	asgnAdminRoleCmd.MarkPersistentFlagRequired("roleId")

	interCmd.AddCommand(asgnAdminRoleCmd)

}
