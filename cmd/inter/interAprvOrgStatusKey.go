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

var aprvOrgStatusKeyCmd = &cobra.Command{
	Use:              "approveOrgStatKey",
	TraverseChildren: true,
	Short:            "update the org status.",
	Long: `    /** @notice interface to approve org status change
	* @param _orgId unique id for the sub organization
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

		r, err := inter.ApproveOrgStatusKey(keyfile, password, orgId, big.NewInt(action))
		if err != nil {
			return err
		}

		fmt.Printf("This is the transaction receipt for Adding the role: %#v\n", r)
		fmt.Println()
		return nil
	},
}

func init() {

	aprvOrgStatusKeyCmd.PersistentFlags().StringP("orgId", "o", "", "The org id to update.")
	aprvOrgStatusKeyCmd.PersistentFlags().Int64P("action", "a", 0, "The action to take 1 for suspending an org and 2 for revoke of suspension")
	aprvOrgStatusKeyCmd.PersistentFlags().StringP("keyfile", "k", "", "The path to the keyfile for authorization.")
	aprvOrgStatusKeyCmd.PersistentFlags().StringP("password", "p", "", "The password for the keyfile for authorization(oupdOrgStatusCmd")

	aprvOrgStatusKeyCmd.MarkPersistentFlagRequired("action")
	aprvOrgStatusKeyCmd.MarkPersistentFlagRequired("orgId")
	aprvOrgStatusKeyCmd.MarkPersistentFlagRequired("keyfile")

	interCmd.AddCommand(aprvOrgStatusKeyCmd)

}
