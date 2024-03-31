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

// addAdminCmd represents the addAdmin command
var addAdminKeyCmd = &cobra.Command{
	Use:              "addAdminKey",
	TraverseChildren: true,
	Short:            "Add an Admin Account to the chain.",
	Long: `    /** @notice interface to add accounts to an admin organization
	* @param _acct account address to be added
	*/`,
	RunE: func(cmd *cobra.Command, args []string) error {
		auth, err := cmd.Flags().GetString("newAdmin")
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

		r, err := inter.AddAdminKey(keyfile, password, auth)
		if err != nil {
			return err

		}
		fmt.Printf("This is the transaction receipt for Adding the role: %#v\n", r)
		fmt.Println()

		return nil
	},
}

func init() {

	addAdminKeyCmd.PersistentFlags().StringP("newAdmin", "n", "", "The address to add as admin.")
	addAdminKeyCmd.PersistentFlags().StringP("keyfile", "k", "", "The path to the keyfile for authorization.")
	addAdminKeyCmd.PersistentFlags().StringP("password", "p", "", "The password for the keyfile for authorization(optional)")

	addAdminKeyCmd.MarkPersistentFlagRequired("newAdmin")
	addAdminKeyCmd.MarkPersistentFlagRequired("keyfile")

	interCmd.AddCommand(addAdminKeyCmd)

}
