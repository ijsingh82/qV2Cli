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

var updNodeStatusCmd = &cobra.Command{
	Use:              "updateNodeStat",
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

		key, err := cmd.Flags().GetString("guardKey")
		if err != nil {
			return err
		}
		guardAdd, err := cmd.Flags().GetString("guardAdd")
		if err != nil {
			return err
		}
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

		orgId, err := cmd.Flags().GetString("orgId")
		if err != nil {
			return err
		}
		action, err := cmd.Flags().GetInt64("action")
		if err != nil {
			return err
		}

		r, err := inter.UpdateNodeStatus(key, guardAdd, orgId, enodeId, ip, port, raftport, big.NewInt(action))
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for updating the node: %#v\n", r)
		fmt.Println()
		return nil

	},
}

func init() {

	updNodeStatusCmd.PersistentFlags().StringP("enodeId", "e", "", "The enode address for the new node add.")
	updNodeStatusCmd.PersistentFlags().StringP("ip", "i", "", "The ip address for the new node.")
	updNodeStatusCmd.PersistentFlags().Uint16P("port", "p", 0, "The port for the new node.")
	updNodeStatusCmd.PersistentFlags().Uint16P("raftport", "t", 0, "The port for the new node.")
	updNodeStatusCmd.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied.")
	updNodeStatusCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied.")
	updNodeStatusCmd.PersistentFlags().StringP("orgId", "o", "", "The org id to update.")
	updNodeStatusCmd.PersistentFlags().Int64P("action", "a", 0, "The action to take  1-deactivate, 2-activate back, 3-blacklist the node")

	updNodeStatusCmd.MarkPersistentFlagRequired("action")
	updNodeStatusCmd.MarkPersistentFlagRequired("guardAdd")
	updNodeStatusCmd.MarkPersistentFlagRequired("guardKey")
	updNodeStatusCmd.MarkPersistentFlagRequired("orgId")
	updNodeStatusCmd.MarkPersistentFlagRequired("enodeId")
	updNodeStatusCmd.MarkPersistentFlagRequired("ip")
	updNodeStatusCmd.MarkPersistentFlagRequired("port")
	updNodeStatusCmd.MarkPersistentFlagRequired("raftport")

	interCmd.AddCommand(updNodeStatusCmd)

}
