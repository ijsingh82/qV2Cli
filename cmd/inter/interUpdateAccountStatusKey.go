/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package inter

import (
	"fmt"
	"math/big"

	"github.com/ijsingh82/qV2Cli/pkg/inter"

	"github.com/spf13/cobra"
)

var updAcctStatusKeyCmd = &cobra.Command{
	Use:              "updateAccountStatKey",
	TraverseChildren: true,
	Short:            "update the account status.",
	Long: `    /** @notice interface to update account status
	this can be executed by org admin accounts only
  * @param _orgId unique id of the organization to which the account belongs
  * @param _account account id
  * @param _action 1-suspending 2-activating back 3-blacklisting
  */`,
	RunE: func(cmd *cobra.Command, args []string) error {

		orgId, err := cmd.Flags().GetString("orgId")
		if err != nil {
			return err
		}
		action, err := cmd.Flags().GetInt64("action")
		if err != nil {
			return err
		}
		address, err := cmd.Flags().GetString("address")
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

		r, err := inter.UpdateAccountStatusKey(keyfile, password, orgId, address, big.NewInt(action))
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for Adding the role: %#v\n", r)
		fmt.Println()
		return nil
	},
}

func init() {

	updAcctStatusKeyCmd.PersistentFlags().StringP("orgId", "o", "", "The org id to update.")
	updAcctStatusKeyCmd.PersistentFlags().StringP("address", "a", "", "The account address to update.")
	updAcctStatusKeyCmd.PersistentFlags().Int64P("action", "u", 0, "The action to take 1 for suspending an org and 2 for revoke of suspension")
	updAcctStatusKeyCmd.PersistentFlags().StringP("keyfile", "k", "", "The path to the keyfile for authorization.")
	updAcctStatusKeyCmd.PersistentFlags().StringP("password", "p", "", "The password for the keyfile for authorization(oupdOrgStatusCmd")

	updAcctStatusKeyCmd.MarkPersistentFlagRequired("action")
	updAcctStatusKeyCmd.MarkPersistentFlagRequired("orgId")
	updAcctStatusKeyCmd.MarkPersistentFlagRequired("keyfile")
	updAcctStatusKeyCmd.MarkPersistentFlagRequired("address")

	interCmd.AddCommand(updAcctStatusKeyCmd)

}
