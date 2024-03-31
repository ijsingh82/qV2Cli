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

var updAcctStatusCmd = &cobra.Command{
	Use:              "updateAccountStat",
	TraverseChildren: true,
	Short:            "update the account status.",
	Long: `    /** @notice interface to update account status
	this can be executed by org admin accounts only
  * @param _orgId unique id of the organization to which the account belongs
  * @param _account account id
  * @param _action 1-suspending 2-activating back 3-blacklisting
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
		address, err := cmd.Flags().GetString("address")
		if err != nil {
			return err
		}
		orgId, err := cmd.Flags().GetString("orgId")
		if err != nil {
			return err
		}
		action, err := cmd.Flags().GetInt64("action")
		if err != nil {
			return err
		}

		r, err := inter.UpdateAccountStatus(key, guardAdd, orgId, address, big.NewInt(action))
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for updating the Account Status: %#v\n", r)
		fmt.Println()
		return nil

	},
}

func init() {

	updAcctStatusCmd.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied.")
	updAcctStatusCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied.")
	updAcctStatusCmd.PersistentFlags().StringP("orgId", "o", "", "The org id to update.")
	updAcctStatusCmd.PersistentFlags().StringP("address", "d", "", "The address of the account to update")
	updAcctStatusCmd.PersistentFlags().Int64P("action", "a", 0, "The action to take 1-suspending 2-activating back 3-blacklisting")

	updAcctStatusCmd.MarkPersistentFlagRequired("action")
	updAcctStatusCmd.MarkPersistentFlagRequired("guardAdd")
	updAcctStatusCmd.MarkPersistentFlagRequired("guardKey")
	updAcctStatusCmd.MarkPersistentFlagRequired("orgId")
	updAcctStatusCmd.MarkPersistentFlagRequired("address")

	interCmd.AddCommand(updAcctStatusCmd)

}
