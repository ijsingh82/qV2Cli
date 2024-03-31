/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package inter

import (
	"fmt"
	"math/big"

	"github.com/ijsingh82/qV2Cli/pkg/inter"
	"github.com/ijsingh82/qV2Cli/pkg/utils"

	"github.com/spf13/cobra"
)

var updOrgStatusCmd = &cobra.Command{
	Use:              "updateOrgStat",
	TraverseChildren: true,
	Short:            "update the org status.",
	Long: `    /** @notice interface to update the org status
	* @param _orgId unique id of the organization
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

		r, err := inter.UpdateOrgStatus(key, guardAdd, orgId, big.NewInt(action))
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for updating org status:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil

	},
}

func init() {

	updOrgStatusCmd.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied.")
	updOrgStatusCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied.")
	updOrgStatusCmd.PersistentFlags().StringP("orgId", "o", "", "The org id to update.")
	updOrgStatusCmd.PersistentFlags().Int64P("action", "a", 0, "The action to take 1 for suspending an org and 2 for revoke of suspension")

	updOrgStatusCmd.MarkPersistentFlagRequired("action")
	updOrgStatusCmd.MarkPersistentFlagRequired("guardAdd")
	updOrgStatusCmd.MarkPersistentFlagRequired("guardKey")
	updOrgStatusCmd.MarkPersistentFlagRequired("orgId")

	interCmd.AddCommand(updOrgStatusCmd)

}
