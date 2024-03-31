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
var addAdminNodeKeyCmd = &cobra.Command{
	Use:              "addAdminNodeKey",
	TraverseChildren: true,
	Short:            "Add an Admin Node to the chain.",
	Long: `    /** @notice interface to add new node to an admin organization
	* @param _enodeId enode id of the node to be added
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

		r, err := inter.AddNodeAsAdminKey(keyfile, password, enodeId, ip, port, raftport)
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for Adding the Admin Node:\n %+v\n", utils.PPrint(r))
		fmt.Println()

		return nil
	},
}

func init() {

	addAdminNodeKeyCmd.PersistentFlags().StringP("enodeId", "e", "", "The enode address for the new node add.")
	addAdminNodeKeyCmd.PersistentFlags().StringP("ip", "i", "", "The ip address for the new node.")
	addAdminNodeKeyCmd.PersistentFlags().Uint16P("port", "p", 0, "The port for the new node.")
	addAdminNodeKeyCmd.PersistentFlags().Uint16P("raftport", "r", 0, "The port for the new node.")
	addAdminNodeKeyCmd.PersistentFlags().StringP("keyfile", "k", "", "The path to the keyfile for authorization.")
	addAdminNodeKeyCmd.PersistentFlags().StringP("password", "w", "", "The password for the keyfile for authorization(optional)")

	addAdminNodeKeyCmd.MarkPersistentFlagRequired("enodeId")
	addAdminNodeKeyCmd.MarkPersistentFlagRequired("raftport")
	addAdminNodeKeyCmd.MarkPersistentFlagRequired("ip")
	addAdminNodeKeyCmd.MarkPersistentFlagRequired("port")
	addAdminNodeKeyCmd.MarkPersistentFlagRequired("keyfile")

	interCmd.AddCommand(addAdminNodeKeyCmd)

}
