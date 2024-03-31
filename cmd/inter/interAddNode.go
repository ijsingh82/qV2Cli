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
var addNodeCmd = &cobra.Command{
	Use:              "addNode",
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

		r, err := inter.AddNode(key, guardAdd, enodeId, ip, orgId, port, raftport)
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for Adding the node:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil
	},
}

func init() {

	addNodeCmd.PersistentFlags().StringP("guardAdd", "a", "", "The guardian Address authorzied.")
	addNodeCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied.")
	addNodeCmd.PersistentFlags().StringP("enodeId", "e", "", "The enode address for the new node add.")
	addNodeCmd.PersistentFlags().StringP("orgId", "o", "", "The orgId of the new org.")
	addNodeCmd.PersistentFlags().StringP("ip", "i", "", "The ip address for the new node.")
	addNodeCmd.PersistentFlags().Uint16P("port", "p", 0, "The port for the new node.")
	addNodeCmd.PersistentFlags().Uint16P("raftport", "r", 0, "The port for the new node.")

	addNodeCmd.MarkPersistentFlagRequired("enodeId")
	addNodeCmd.MarkPersistentFlagRequired("ip")
	addNodeCmd.MarkPersistentFlagRequired("port")
	addNodeCmd.MarkPersistentFlagRequired("raftport")
	addNodeCmd.MarkPersistentFlagRequired("orgId")
	addNodeCmd.MarkPersistentFlagRequired("guardKey")
	addNodeCmd.MarkPersistentFlagRequired("guardAdd")

	interCmd.AddCommand(addNodeCmd)

}
