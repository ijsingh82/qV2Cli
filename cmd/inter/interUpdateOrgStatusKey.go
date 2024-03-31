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

var updOrgStatusKeyCmd = &cobra.Command{
	Use:              "updateOrgStatKey",
	TraverseChildren: true,
	Short:            "update the org status.",
	Long: `    /** @notice interface to update the org status
	* @param _orgId unique id of the organization
	* @param _action 1 for suspending an org and 2 for revoke of suspension
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
		keyfile, err := cmd.Flags().GetString("keyfile")
		if err != nil {
			return err
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			return err
		}

		r, err := inter.UpdateOrgStatusKey(keyfile, password, orgId, big.NewInt(action))
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for updating org status:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil
	},
}

func init() {

	updOrgStatusKeyCmd.PersistentFlags().StringP("orgId", "o", "", "The org id to update.")
	updOrgStatusKeyCmd.PersistentFlags().Int64P("action", "a", 0, "The action to take 1 for suspending an org and 2 for revoke of suspension")
	updOrgStatusKeyCmd.PersistentFlags().StringP("keyfile", "f", "", "The path to the keyfile for authorization.")
	updOrgStatusKeyCmd.PersistentFlags().StringP("password", "p", "", "The password for the keyfile for authorization(oupdOrgStatusCmd")

	updOrgStatusKeyCmd.MarkPersistentFlagRequired("action")
	updOrgStatusKeyCmd.MarkPersistentFlagRequired("orgId")
	updOrgStatusKeyCmd.MarkPersistentFlagRequired("keyfile")

	interCmd.AddCommand(updOrgStatusKeyCmd)

}
