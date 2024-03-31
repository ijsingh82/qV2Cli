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

// addAdminCmd represents the addAdmin command
var addAdminNodeCmd = &cobra.Command{
	Use:              "addAdminNode",
	TraverseChildren: true,
	Short:            "Add an Admin Node to the chain.",
	Long: `    /** @notice interface to add new node to an admin organization
	* @param _enodeId enode id of the node to be added
	* @param _ip IP of node
	* @param _port tcp port of node
	* @param _raftport raft port of node
	*/`,
	RunE: func(cmd *cobra.Command, args []string) error {

		key, err := cmd.Flags().GetString("guardKey")
		if err != nil {
			fmt.Println("Get guarKey Error")
			return err
		}
		guardAdd, err := cmd.Flags().GetString("guardAdd")
		if err != nil {
			fmt.Println("Get guradAdd Error")
			return err
		}
		enodeId, err := cmd.Flags().GetString("enodeId")
		if err != nil {
			fmt.Println("Get enodeId Error")
			return err
		}
		ip, err := cmd.Flags().GetString("ip")
		if err != nil {
			fmt.Println("Get ip Error")
			return err
		}
		port, err := cmd.Flags().GetUint16("port")
		if err != nil {
			fmt.Println("Get port Error")
			return err
		}
		raftport, err := cmd.Flags().GetUint16("raftport")
		if err != nil {
			fmt.Println("Get raftport Error")
			return err
		}

		r, err := inter.AddNodeAsAdmin(key, guardAdd, enodeId, ip, port, raftport)
		if err != nil {
			fmt.Println("Get addAddressAsAdmin Error")
			return err
		}

		fmt.Printf("This is the transaction receipt for Adding the role: %#v\n", r)
		fmt.Println()

		return nil
	},
}

func init() {

	addAdminNodeCmd.PersistentFlags().StringP("guardAdd", "a", "", "The guardian Address authorzied.")
	addAdminNodeCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied.")
	addAdminNodeCmd.PersistentFlags().StringP("enodeId", "e", "", "The enode address for the new node add.")
	addAdminNodeCmd.PersistentFlags().StringP("ip", "i", "", "The ip address for the new node.")
	addAdminNodeCmd.PersistentFlags().Uint16P("port", "p", 0, "The port for the new node.")
	addAdminNodeCmd.PersistentFlags().Uint16P("raftport", "r", 0, "The port for the new node.")

	addAdminNodeCmd.MarkPersistentFlagRequired("enodeId")
	addAdminNodeCmd.MarkPersistentFlagRequired("raftport")
	addAdminNodeCmd.MarkPersistentFlagRequired("port")
	addAdminNodeCmd.MarkPersistentFlagRequired("ip")
	addAdminNodeCmd.MarkPersistentFlagRequired("guardKey")
	addAdminNodeCmd.MarkPersistentFlagRequired("guardAdd")

	interCmd.AddCommand(addAdminNodeCmd)

}
