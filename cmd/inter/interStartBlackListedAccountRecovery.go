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
var strtBlkLstAcctRcvry = &cobra.Command{
	Use:              "startBlackListedAccountRecovery",
	TraverseChildren: true,
	Short:            "Start recovery of a blacklisted account.",
	Long: `    /** @notice interface to initiate blacklisted account recovery
	* @param _orgId unique id of the organization to which the account belongs
	* @param _account account id being recovered
	*/`,
	RunE: func(cmd *cobra.Command, args []string) error {

		address, err := cmd.Flags().GetString("address")
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

		r, err := inter.StartBlackListedAccountRecovery(key, guardAdd, address, orgId)
		if err != nil {
			return err
		}

		fmt.Printf("This is the transaction receipt for starting recovery of account: %#v\n", r)
		fmt.Println()
		return nil
	},
}

func init() {

	strtBlkLstAcctRcvry.PersistentFlags().StringP("address", "a", "", "The address of the account to recover.")
	strtBlkLstAcctRcvry.PersistentFlags().StringP("orgId", "o", "", "The org the account is assigned to")
	strtBlkLstAcctRcvry.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied for certChain.")
	strtBlkLstAcctRcvry.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied for certChain.")

	strtBlkLstAcctRcvry.MarkPersistentFlagRequired("address")
	strtBlkLstAcctRcvry.MarkPersistentFlagRequired("orgId")
	strtBlkLstAcctRcvry.MarkPersistentFlagRequired("guardAdd")
	strtBlkLstAcctRcvry.MarkPersistentFlagRequired("guardKey")

	interCmd.AddCommand(strtBlkLstAcctRcvry)

}
