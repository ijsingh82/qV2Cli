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
var addRoleKeyCmd = &cobra.Command{
	Use:              "addRoleKey",
	TraverseChildren: true,
	Short:            "Add a role to the specified org in the chain.",
	Long: `    /** @notice interface to add a new role definition to an organization
	* @param _roleId unique id for the role
	* @param _orgId unique id of the organization to which the role belongs
	* @param _access account access type for the role
	* @param _voter bool indicates if the role is voter role or not
	* @param _admin bool indicates if the role is an admin role
	* @dev account access type can have of the following four values:
		  0 - Read only
		  1 - Transact access
		  2 - Contract deployment access. Can transact as well
		  3 - Full access
		  4 - contract call
		  5 - value transfer and contract call
		  6 - value transfer and contract deploy
		  7 - contract call and deploy
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
		accLvl, err := cmd.Flags().GetInt64("accLvl")
		if err != nil {
			return err
		}
		voter, err := cmd.Flags().GetBool("voter")
		if err != nil {
			return err
		}
		admin, err := cmd.Flags().GetBool("admin")
		if err != nil {
			return err
		}
		guardKeyFile, err := cmd.Flags().GetString("keyFile")
		if err != nil {
			return err
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			return err
		}

		r, err := inter.AddNewRoleKey(guardKeyFile, password, roleID, orgId, voter, admin, accLvl)
		if err != nil {
			return err
		}

		fmt.Printf("This is the transaction receipt for Adding the role: %#v\n", r)
		fmt.Println()

		return nil
	},
}

func init() {

	addRoleKeyCmd.PersistentFlags().StringP("roleId", "r", "", "The name for the new role id.")
	addRoleKeyCmd.PersistentFlags().StringP("keyFile", "k", "", "The guardian keystore path.")
	addRoleKeyCmd.PersistentFlags().StringP("password", "p", "", "The guardian keystore password.")
	addRoleKeyCmd.PersistentFlags().StringP("orgId", "o", "", "The org to add the role to certChain.")
	addRoleKeyCmd.PersistentFlags().Int64P("accLvl", "l", 7, "The level of access for the role. (defualt - Transact and contract deploy only)")
	addRoleKeyCmd.PersistentFlags().BoolP("voter", "v", false, "true or false for role voter (defualt false)")
	addRoleKeyCmd.PersistentFlags().BoolP("admin", "a", false, "true or false for role is admin (defualt false)")

	addRoleKeyCmd.MarkPersistentFlagRequired("keyFile")
	addRoleKeyCmd.MarkPersistentFlagRequired("roleId")
	addRoleKeyCmd.MarkPersistentFlagRequired("orgId")

	interCmd.AddCommand(addRoleKeyCmd)

}
