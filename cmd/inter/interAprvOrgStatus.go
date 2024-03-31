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

var aprvOrgStatusCmd = &cobra.Command{
	Use:              "approveOrgStat",
	TraverseChildren: true,
	Short:            "update the org status.",
	Long: `    /** @notice interface to approve org status change
	* @param _orgId unique id for the sub organization
	* @param _action 1 for suspending an org and 2 for revoke of suspension
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
		orgId, err := cmd.Flags().GetString("orgId")
		if err != nil {
			return err
		}
		action, err := cmd.Flags().GetInt64("action")
		if err != nil {
			return err
		}

		r, err := inter.ApproveOrgStatus(key, guardAdd, orgId, big.NewInt(action))
		if err != nil {
			return err
		}

		fmt.Printf("This is the transaction receipt for Adding the role: %#v\n", r)
		fmt.Println()
		return nil
	},
}

func init() {

	aprvOrgStatusCmd.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied.")
	aprvOrgStatusCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied.")
	aprvOrgStatusCmd.PersistentFlags().StringP("orgId", "o", "", "The org id to update.")
	aprvOrgStatusCmd.PersistentFlags().Int64P("action", "a", 0, "The action to take 1 for suspending an org and 2 for revoke of suspension")

	aprvOrgStatusCmd.MarkPersistentFlagRequired("action")
	aprvOrgStatusCmd.MarkPersistentFlagRequired("orgId")
	aprvOrgStatusCmd.MarkPersistentFlagRequired("guardAdd")
	aprvOrgStatusCmd.MarkPersistentFlagRequired("guardKey")

	interCmd.AddCommand(aprvOrgStatusCmd)

}
