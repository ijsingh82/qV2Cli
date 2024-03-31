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
var addOrgCmd = &cobra.Command{
	Use:              "addOrg",
	TraverseChildren: true,
	Short:            "Add an new Org with a Node to the chain.",
	Long: `    /** @notice interface to add a new organization to the network
	* @param _orgId unique organization id
	* @param _enodeId enode id linked to the organization
	* @param _ip IP of node
	* @param _port tcp port of node
	* @param _raftport raft port of node
	* @param _account account id. this will have the org admin privileges
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
		address, err := cmd.Flags().GetString("address")
		if err != nil {
			return err
		}
		r, err := inter.AddOrg(key, guardAdd, enodeId, ip, orgId, address, port, raftport)
		if err != nil {
			return err
		}
		fmt.Printf("This is the transaction receipt for Adding the role: %#v\n", r)
		fmt.Println()
		return nil
	},
}

func init() {

	addOrgCmd.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied.")
	addOrgCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied.")
	addOrgCmd.PersistentFlags().StringP("enodeId", "e", "", "The enode address for the new node add.")
	addOrgCmd.PersistentFlags().StringP("orgId", "o", "", "The orgId of the new org.")
	addOrgCmd.PersistentFlags().StringP("ip", "i", "", "The ip address for the new node.")
	addOrgCmd.PersistentFlags().Uint16P("port", "p", 0, "The port for the new node.")
	addOrgCmd.PersistentFlags().Uint16P("raftport", "r", 0, "The port for the new node.")
	addOrgCmd.PersistentFlags().StringP("address", "a", "", "the address for the admin account.")

	addOrgCmd.MarkPersistentFlagRequired("enodeId")
	addOrgCmd.MarkPersistentFlagRequired("ip")
	addOrgCmd.MarkPersistentFlagRequired("port")
	addOrgCmd.MarkPersistentFlagRequired("raftport")
	addOrgCmd.MarkPersistentFlagRequired("orgId")
	addOrgCmd.MarkPersistentFlagRequired("address")
	addOrgCmd.MarkPersistentFlagRequired("guardKey")
	addOrgCmd.MarkPersistentFlagRequired("guardAdd")

	interCmd.AddCommand(addOrgCmd)

}
