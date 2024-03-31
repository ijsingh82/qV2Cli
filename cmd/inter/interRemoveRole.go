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

// addRoleCmd represents the addRole command
var removeRoleCmd = &cobra.Command{
	Use:              "removeRole",
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
		key, err := cmd.Flags().GetString("guardKey")
		if err != nil {
			return err
		}
		guardAdd, err := cmd.Flags().GetString("guardAdd")
		if err != nil {
			return err
		}
		orgId, err := cmd.Flags().GetString("orgId")
		if err != nil {
			return err
		}

		r, err := inter.RmvRole(key, guardAdd, roleID, orgId)
		if err != nil {
			return err
		}

		fmt.Printf("This is the transaction receipt for removing the role:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil
	},
}

func init() {

	removeRoleCmd.PersistentFlags().StringP("roleId", "i", "", "The name for the new role id.")
	removeRoleCmd.PersistentFlags().StringP("orgId", "o", "", "The org to add the role to certChain.")
	removeRoleCmd.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied for certChain.")
	removeRoleCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied for certChain.")

	removeRoleCmd.MarkPersistentFlagRequired("roleId")
	removeRoleCmd.MarkPersistentFlagRequired("orgId")
	removeRoleCmd.MarkPersistentFlagRequired("guardAdd")
	removeRoleCmd.MarkPersistentFlagRequired("guardKey")

	interCmd.AddCommand(removeRoleCmd)

}
