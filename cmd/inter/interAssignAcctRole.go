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
var asgnRoleCmd = &cobra.Command{
	Use:              "assignAcctRole",
	TraverseChildren: true,
	Short:            "Assign an address for a role to the specified org in the chain.",
	Long: `    /** @notice interface to assigns a role id to the account give
	* @param _account account id
	* @param _orgId organization id to which the account belongs
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

		r, err := inter.AssignRole(keyfile, password, newAcct, roleID, orgId)
		if err != nil {
			return err
		}

		fmt.Printf("This is the transaction receipt for assigning account role:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil
	},
}

func init() {

	asgnRoleCmd.PersistentFlags().StringP("roleId", "r", "", "The role id to be assigned to account.")
	asgnRoleCmd.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied for Chain.")
	asgnRoleCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied for Chain.")
	asgnRoleCmd.PersistentFlags().StringP("orgId", "o", "", "The org Id the role ID belongs to in Chain.")
	asgnRoleCmd.PersistentFlags().StringP("newAcc", "n", "", "The new account being approved.")

	asgnRoleCmd.MarkPersistentFlagRequired("guardAdd")
	asgnRoleCmd.MarkPersistentFlagRequired("newAcc")
	asgnRoleCmd.MarkPersistentFlagRequired("orgId")
	asgnRoleCmd.MarkPersistentFlagRequired("guardKey")
	asgnRoleCmd.MarkPersistentFlagRequired("roleId")

	interCmd.AddCommand(asgnRoleCmd)

}
