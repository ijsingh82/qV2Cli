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

var aprvBlkListNodeRcvry = &cobra.Command{
	Use:              "approveBlackListNodeRecovery",
	TraverseChildren: true,
	Short:            "approve the recovery of a node.",
	Long: `    /** @notice interface to approve blacklisted node recoevry
	* @param _orgId unique id of the organization to which the account belongs
	* @param _enodeId enode id being recovered
	* @param _ip IP of node
	* @param _port tcp port of node
	* @param _raftport raft port of node
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

		r, err := inter.ApproveBlackListedNodeRecovery(key, guardAdd, orgId, enodeId, ip, port, raftport)
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for updating the node: %#v\n", r)
		fmt.Println()
		return nil

	},
}

func init() {

	aprvBlkListNodeRcvry.PersistentFlags().StringP("enodeId", "e", "", "The enode address for the node.")
	aprvBlkListNodeRcvry.PersistentFlags().StringP("ip", "i", "", "The ip address for the node.")
	aprvBlkListNodeRcvry.PersistentFlags().Uint16P("port", "p", 0, "The port for the node.")
	aprvBlkListNodeRcvry.PersistentFlags().Uint16P("raftport", "r", 0, "The port for the node.")
	aprvBlkListNodeRcvry.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied.")
	aprvBlkListNodeRcvry.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key authorzied.")
	aprvBlkListNodeRcvry.PersistentFlags().StringP("orgId", "o", "", "The org id to update.")

	aprvBlkListNodeRcvry.MarkPersistentFlagRequired("guardAdd")
	aprvBlkListNodeRcvry.MarkPersistentFlagRequired("guardKey")
	aprvBlkListNodeRcvry.MarkPersistentFlagRequired("orgId")
	aprvBlkListNodeRcvry.MarkPersistentFlagRequired("port")
	aprvBlkListNodeRcvry.MarkPersistentFlagRequired("raftport")
	aprvBlkListNodeRcvry.MarkPersistentFlagRequired("enodeId")
	aprvBlkListNodeRcvry.MarkPersistentFlagRequired("ip")

	interCmd.AddCommand(aprvBlkListNodeRcvry)

}
