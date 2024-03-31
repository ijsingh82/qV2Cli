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

// addAdminCmd represents the addAdmin command
var addAdminCmd = &cobra.Command{
	Use:              "addAdmin",
	TraverseChildren: true,
	Short:            "Add an Admin Account to the chain.",
	Long: `    /** @notice interface to add accounts to an admin organization
	* @param _acct account address to be added
	*/`,
	RunE: func(cmd *cobra.Command, args []string) error {

		key, err := cmd.Flags().GetString("guardKey")
		if err != nil {
			return err
		}
		guardAdd, err := cmd.Flags().GetString("guardAdd")
		if err != nil {
			return err
		}
		auth, err := cmd.Flags().GetString("newAdmin")
		if err != nil {
			return err
		}

		r, err := inter.AddAdmin(key, guardAdd, auth)
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for Adding the admin:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil
	},
}

func init() {

	addAdminCmd.PersistentFlags().StringP("guardAdd", "a", "", "The guardian Address authorzied.")
	addAdminCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied.")
	addAdminCmd.PersistentFlags().StringP("newAdmin", "n", "", "The address to add as admin.")

	addAdminCmd.MarkPersistentFlagRequired("newAdmin")
	addAdminCmd.MarkPersistentFlagRequired("guardKey")
	addAdminCmd.MarkPersistentFlagRequired("guardAdd")

	interCmd.AddCommand(addAdminCmd)

}
