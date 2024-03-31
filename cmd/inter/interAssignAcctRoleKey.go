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
var asgnAcctRoleKeyCmd = &cobra.Command{
	Use:              "assignAcctRoleKey",
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

		r, err := inter.AssignRoleKey(keyfile, password, newAcct, roleID, orgId)
		if err != nil {
			return err
		}

		fmt.Printf("This is the transaction receipt for Adding the role: %#v\n", r)
		fmt.Println()
		return nil
	},
}

func init() {

	asgnAcctRoleKeyCmd.PersistentFlags().StringP("roleId", "r", "", "The role id to be assigned to account.")
	asgnAcctRoleKeyCmd.PersistentFlags().StringP("keyfile", "k", "", "The authorizer keyfile.")
	asgnAcctRoleKeyCmd.PersistentFlags().StringP("password", "p", "", "The authorizer keyfile password.")
	asgnAcctRoleKeyCmd.PersistentFlags().StringP("orgId", "o", "", "The org Id the role ID belongs to in Chain.")
	asgnAcctRoleKeyCmd.PersistentFlags().StringP("newAcc", "n", "", "The new account being approved.")

	asgnAcctRoleKeyCmd.MarkPersistentFlagRequired("keyfile")
	asgnAcctRoleKeyCmd.MarkPersistentFlagRequired("newAcc")
	asgnAcctRoleKeyCmd.MarkPersistentFlagRequired("orgId")
	asgnAcctRoleKeyCmd.MarkPersistentFlagRequired("roleId")

	interCmd.AddCommand(asgnAcctRoleKeyCmd)

}
