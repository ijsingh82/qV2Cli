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
var strtBlkLstAcctRcvryKey = &cobra.Command{
	Use:              "startBlackListedAccountRecoveryKey",
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
		fmt.Printf("This is the transaction receipt for starting recovery of account:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil
	},
}

func init() {

	strtBlkLstAcctRcvryKey.PersistentFlags().StringP("address", "a", "", "The address to start the recovery.")
	strtBlkLstAcctRcvryKey.PersistentFlags().StringP("orgId", "o", "", "The org to add the role to certChain.")
	strtBlkLstAcctRcvryKey.PersistentFlags().StringP("keyfile", "k", "", "The guardian keystore path.")
	strtBlkLstAcctRcvryKey.PersistentFlags().StringP("password", "p", "", "The guardian keystore password.")

	strtBlkLstAcctRcvryKey.MarkPersistentFlagRequired("address")
	strtBlkLstAcctRcvryKey.MarkPersistentFlagRequired("orgId")
	strtBlkLstAcctRcvryKey.MarkPersistentFlagRequired("keyfile")

	interCmd.AddCommand(strtBlkLstAcctRcvryKey)

}
