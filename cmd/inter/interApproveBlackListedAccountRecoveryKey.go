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
var aprvBlkLstAcctRcvryKey = &cobra.Command{
	Use:              "approveBlackListedAccountRecoveryKey",
	TraverseChildren: true,
	Short:            "Approve recovery of a blacklisted account.",
	Long: `    /** @notice interface to approve blacklisted node recovery
	* @param _orgId unique id of the organization to which the account belongs
	* @param _account account id being recovered
	*/`,
	RunE: func(cmd *cobra.Command, args []string) error {

		address, err := cmd.Flags().GetString("address")
		if err != nil {
			return err
		}

		orgId, err := cmd.Flags().GetString("orgId")
		if err != nil {
			return err
		}

		guardKeyFile, err := cmd.Flags().GetString("keyfile")
		if err != nil {
			return err
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			return err
		}

		r, err := inter.StartBlackListedAccountRecoveryKey(guardKeyFile, password, address, orgId)
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for approving recovery of account: %#v\n", r)
		fmt.Println()
		return nil
	},
}

func init() {

	aprvBlkLstAcctRcvryKey.PersistentFlags().StringP("address", "a", "", "The address to start the recovery.")
	aprvBlkLstAcctRcvryKey.PersistentFlags().StringP("orgId", "o", "", "The org to add the role to certChain.")
	aprvBlkLstAcctRcvryKey.PersistentFlags().StringP("keyfile", "f", "", "The guardian keystore path.")
	aprvBlkLstAcctRcvryKey.PersistentFlags().StringP("password", "p", "", "The guardian keystore password.")

	aprvBlkLstAcctRcvryKey.MarkPersistentFlagRequired("address")
	aprvBlkLstAcctRcvryKey.MarkPersistentFlagRequired("orgId")
	aprvBlkLstAcctRcvryKey.MarkPersistentFlagRequired("keyfile")

	interCmd.AddCommand(aprvBlkLstAcctRcvryKey)

}
