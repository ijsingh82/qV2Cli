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

// addRoleCmd represents the addRole command
var removeRoleKeyCmd = &cobra.Command{
	Use:              "removeRoleKey",
	TraverseChildren: true,
	Short:            "Remove a role from the chain.",
	Long: `    /** @notice interface to remove a role definition from an organization
	* @param _roleId unique id for the role
	* @param _orgId unique id of the organization to which the role belongs
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

		keyfile, err := cmd.Flags().GetString("keyfile")
		if err != nil {
			return err
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			return err
		}

		r, err := inter.RmvRoleKey(keyfile, password, roleID, orgId)
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for Adding the role: %#v\n", r)
		fmt.Println()
		return nil
	},
}

func init() {

	removeRoleKeyCmd.PersistentFlags().StringP("roleId", "i", "", "The name for the new role id.")
	removeRoleKeyCmd.PersistentFlags().StringP("orgId", "o", "", "The org to add the role to certChain.")
	removeRoleKeyCmd.PersistentFlags().StringP("keyfile", "f", "", "The guardian keystore path.")
	removeRoleKeyCmd.PersistentFlags().StringP("password", "p", "", "The guardian keystore password.")

	removeRoleKeyCmd.MarkPersistentFlagRequired("roleId")
	removeRoleKeyCmd.MarkPersistentFlagRequired("orgId")
	removeRoleKeyCmd.MarkPersistentFlagRequired("keyfile")

	interCmd.AddCommand(removeRoleKeyCmd)

}
