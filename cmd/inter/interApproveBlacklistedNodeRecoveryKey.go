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

var aprvBlkListNodeRcvryKey = &cobra.Command{
	Use:              "approveBlackListNodeRecoveryKey",
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

		key, err := cmd.Flags().GetString("keyfile")
		if err != nil {
			return err
		}
		guardAdd, err := cmd.Flags().GetString("password")
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

		r, err := inter.ApproveBlackListedNodeRecoveryKey(key, guardAdd, orgId, enodeId, ip, port, raftport)
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for approving node recovery:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil

	},
}

func init() {

	aprvBlkListNodeRcvryKey.PersistentFlags().StringP("enodeId", "e", "", "The enode address for the node.")
	aprvBlkListNodeRcvryKey.PersistentFlags().StringP("ip", "i", "", "The ip address for the node.")
	aprvBlkListNodeRcvryKey.PersistentFlags().Uint16P("port", "p", 0, "The port for the node.")
	aprvBlkListNodeRcvryKey.PersistentFlags().Uint16P("raftport", "r", 0, "The port for the node.")
	aprvBlkListNodeRcvryKey.PersistentFlags().StringP("keyfile", "k", "", "The keyfile path to the keystore for transaction.")
	aprvBlkListNodeRcvryKey.PersistentFlags().StringP("password", "w", "", "The password for the keystore.")
	aprvBlkListNodeRcvryKey.PersistentFlags().StringP("orgId", "o", "", "The org.")

	aprvBlkListNodeRcvryKey.MarkPersistentFlagRequired("ip")
	aprvBlkListNodeRcvryKey.MarkPersistentFlagRequired("keyfile")
	aprvBlkListNodeRcvryKey.MarkPersistentFlagRequired("enodeId")
	aprvBlkListNodeRcvryKey.MarkPersistentFlagRequired("orgId")
	aprvBlkListNodeRcvryKey.MarkPersistentFlagRequired("port")
	aprvBlkListNodeRcvryKey.MarkPersistentFlagRequired("raftport")

	interCmd.AddCommand(aprvBlkListNodeRcvryKey)

}
