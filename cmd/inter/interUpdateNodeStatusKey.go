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

// addAdminCmd represents the addAdmin command
var updNodeStatusKeyCmd = &cobra.Command{
	Use:              "updateNodeStatKey",
	TraverseChildren: true,
	Short:            "update the node status.",
	Long: `    /** @notice interface to update node status
	* @param _orgId unique id of the organization to which the account belongs
	* @param _enodeId enode id being dded to the org
	* @param _ip IP of node
	* @param _port tcp port of node
	* @param _raftport raft port of node
	* @param _action 1-deactivate, 2-activate back, 3-blacklist the node
	*/`,
	RunE: func(cmd *cobra.Command, args []string) error {

		enodeId, err := cmd.Flags().GetString("enodeId")
		if err != nil {
			return err
		}
		ip, err := cmd.Flags().GetString("ip")
		if err != nil {
			return err
		}
		port, err := cmd.Flags().GetUint16("port")
		if err != nil {
			return err
		}
		raftport, err := cmd.Flags().GetUint16("raftport")
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
		orgId, err := cmd.Flags().GetString("orgId")
		if err != nil {
			return err
		}

		r, err := inter.UpdateNodeStatusKey(keyfile, password, enodeId, ip, orgId, port, raftport, big.NewInt(action))
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for updating the node:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil
	},
}

func init() {

	updNodeStatusKeyCmd.PersistentFlags().StringP("keyfile", "k", "", "The keyfile to use for transaction.")
	updNodeStatusKeyCmd.PersistentFlags().StringP("password", "w", "", "The password for the keyfile if set.")
	updNodeStatusKeyCmd.PersistentFlags().StringP("enodeId", "e", "", "The enode address for the new node add.")
	updNodeStatusKeyCmd.PersistentFlags().StringP("orgId", "o", "", "The orgId of the new org.")
	updNodeStatusKeyCmd.PersistentFlags().StringP("ip", "i", "", "The ip address for the new node.")
	updNodeStatusKeyCmd.PersistentFlags().Uint16P("port", "p", 0, "The port for the new node.")
	updNodeStatusKeyCmd.PersistentFlags().Uint16P("raftport", "t", 0, "The port for the new node.")
	updNodeStatusKeyCmd.PersistentFlags().Int64P("action", "a", 0, "The action to take  1-deactivate, 2-activate back, 3-blacklist the node")

	updNodeStatusKeyCmd.MarkPersistentFlagRequired("enodeId")
	updNodeStatusKeyCmd.MarkPersistentFlagRequired("ip")
	updNodeStatusKeyCmd.MarkPersistentFlagRequired("port")
	updNodeStatusKeyCmd.MarkPersistentFlagRequired("raftport")
	updNodeStatusKeyCmd.MarkPersistentFlagRequired("orgId")
	updNodeStatusKeyCmd.MarkPersistentFlagRequired("keyfile")
	updNodeStatusKeyCmd.MarkPersistentFlagRequired("action")

	interCmd.AddCommand(updNodeStatusKeyCmd)

}
