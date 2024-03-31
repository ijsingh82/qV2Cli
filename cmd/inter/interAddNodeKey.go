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

// addAdminCmd represents the addAdmin command
var addNodeKeyCmd = &cobra.Command{
	Use:              "addNodeKey",
	TraverseChildren: true,
	Short:            "Add an new a Node to an org.",
	Long: `    /** @notice interface to add a new node to the organization
	* @param _orgId unique id of the organization to which the account belongs
	* @param _enodeId enode id being dded to the org
	* @param _ip IP of node
	* @param _port tcp port of node
	* @param _raftport raft port of node
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

		r, err := inter.AddNodeKey(keyfile, password, enodeId, ip, orgId, port, raftport)
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for Adding the node:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil
	},
}

func init() {

	addNodeKeyCmd.PersistentFlags().StringP("password", "w", "", "The password for the keystore.")
	addNodeKeyCmd.PersistentFlags().StringP("keyfile", "k", "", "The path to the keystore.")
	addNodeKeyCmd.PersistentFlags().StringP("enodeId", "e", "", "The enode address for the new node add.")
	addNodeKeyCmd.PersistentFlags().StringP("orgId", "o", "", "The orgId of the new org.")
	addNodeKeyCmd.PersistentFlags().StringP("ip", "i", "", "The ip address for the new node.")
	addNodeKeyCmd.PersistentFlags().Uint16P("port", "p", 0, "The port for the new node.")
	addNodeKeyCmd.PersistentFlags().Uint16P("raftport", "r", 0, "The port for the new node.")

	addNodeKeyCmd.MarkPersistentFlagRequired("enodeId")
	addNodeKeyCmd.MarkPersistentFlagRequired("ip")
	addNodeKeyCmd.MarkPersistentFlagRequired("port")
	addNodeKeyCmd.MarkPersistentFlagRequired("raftport")
	addNodeKeyCmd.MarkPersistentFlagRequired("orgId")
	addNodeKeyCmd.MarkPersistentFlagRequired("keyfile")

	interCmd.AddCommand(addNodeKeyCmd)

}
