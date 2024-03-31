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
var addRoleCmd = &cobra.Command{
	Use:              "addRole",
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
		r, err := inter.AddNewRole(key, guardAdd, roleID, orgId, voter, admin, accLvl)
		if err != nil {
			return err
		}

		fmt.Printf("This is the transaction receipt for Adding the role:\n %+v\n", utils.PPrint(r))
		fmt.Println()

		return nil
	},
}

func init() {

	addRoleCmd.PersistentFlags().StringP("roleId", "i", "", "The name for the new role id.")
	addRoleCmd.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied for certChain.")
	addRoleCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied for certChain.")
	addRoleCmd.PersistentFlags().StringP("orgId", "o", "", "The org to add the role to certChain.")
	addRoleCmd.PersistentFlags().Int64P("accLvl", "l", 7, "The level of access for the role. (defualt - Transact and contract deploy only)")
	addRoleCmd.PersistentFlags().BoolP("voter", "v", false, "true or false for role voter (defualt false)")
	addRoleCmd.PersistentFlags().BoolP("admin", "a", false, "true or false for role is admin (defualt false)")

	addRoleCmd.MarkPersistentFlagRequired("roleId")
	addRoleCmd.MarkPersistentFlagRequired("orgId")
	addRoleCmd.MarkPersistentFlagRequired("guardAdd")
	addRoleCmd.MarkPersistentFlagRequired("guardKey")

	interCmd.AddCommand(addRoleCmd)

}
