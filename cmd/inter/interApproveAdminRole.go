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
var aprvAdminRoleCmd = &cobra.Command{
	Use:              "approveAdminRole",
	TraverseChildren: true,
	Short:            "Approve an address for the admin role to the specified org in the chain.",
	Long: `    /** @notice interface to approve network admin/org admin role assigment
	this can be executed by network admin accounts only
  * @param _orgId unique id of the organization to which the account belongs
  * @param _account account id
  */`,
	RunE: func(cmd *cobra.Command, args []string) error {

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

		r, err := inter.ApproveAdminRole(keyfile, password, newAcct, orgId)
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for approving Admin role: %#v\n", r)

		fmt.Println()
		return nil
	},
}

func init() {

	aprvAdminRoleCmd.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied for Chain.")
	aprvAdminRoleCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied for Chain.")
	aprvAdminRoleCmd.PersistentFlags().StringP("orgId", "o", "", "The org Id the role ID belongs to in Chain.")
	aprvAdminRoleCmd.PersistentFlags().StringP("newAcc", "n", "", "The new account being approved.")

	aprvAdminRoleCmd.MarkPersistentFlagRequired("guardAdd")
	aprvAdminRoleCmd.MarkPersistentFlagRequired("newAcc")
	aprvAdminRoleCmd.MarkPersistentFlagRequired("orgId")
	aprvAdminRoleCmd.MarkPersistentFlagRequired("guardKey")

	interCmd.AddCommand(aprvAdminRoleCmd)

}
