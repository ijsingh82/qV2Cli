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
var aprvBlkLstAcctRcvry = &cobra.Command{
	Use:              "approveBlackListedAccountRecovery",
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

		r, err := inter.ApproveBlackListedAccountRecovery(key, guardAdd, address, orgId)
		if err != nil {
			return err
		}

		fmt.Printf("This is the transaction receipt for approving recovery of account:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil
	},
}

func init() {

	aprvBlkLstAcctRcvry.PersistentFlags().StringP("address", "a", "", "The address of the account to recover.")
	aprvBlkLstAcctRcvry.PersistentFlags().StringP("orgId", "o", "", "The org the account is assigned to")
	aprvBlkLstAcctRcvry.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied for certChain.")
	aprvBlkLstAcctRcvry.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied for certChain.")

	aprvBlkLstAcctRcvry.MarkPersistentFlagRequired("address")
	aprvBlkLstAcctRcvry.MarkPersistentFlagRequired("orgId")
	aprvBlkLstAcctRcvry.MarkPersistentFlagRequired("guardAdd")
	aprvBlkLstAcctRcvry.MarkPersistentFlagRequired("guardKey")

	interCmd.AddCommand(aprvBlkLstAcctRcvry)

}
