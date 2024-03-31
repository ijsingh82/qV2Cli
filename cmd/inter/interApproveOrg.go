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
var aprvOrgCmd = &cobra.Command{
	Use:              "approveOrg",
	TraverseChildren: true,
	Short:            "Add an new Org with a Node to the chain.",
	Long: `    /** @notice interface to approve a newly added organization
	* @param _orgId unique organization id
	* @param _enodeId enode id linked to the organization
	* @param _ip IP of node
	* @param _port tcp port of node
	* @param _raftport raft port of node
	* @param _account account id this will have the org admin privileges
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
		admin, err := cmd.Flags().GetString("admin")
		if err != nil {
			return err
		}
		r, err := inter.ApproveOrg(key, guardAdd, enodeId, ip, orgId, admin, port, raftport)
		if err != nil {
			return err
		}

		fmt.Printf("This is the transaction receipt for approve the org:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil
	},
}

func init() {

	aprvOrgCmd.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied.")
	aprvOrgCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied.")
	aprvOrgCmd.PersistentFlags().StringP("enodeId", "e", "", "The enode address for the new node add.")
	aprvOrgCmd.PersistentFlags().StringP("ip", "i", "", "The ip address for the new node.")
	aprvOrgCmd.PersistentFlags().Uint16P("port", "p", 0, "The port for the new node.")
	aprvOrgCmd.PersistentFlags().Uint16P("raftport", "r", 0, "The port for the new node.")
	aprvOrgCmd.PersistentFlags().StringP("orgId", "o", "", "The orgId of the new org.")
	aprvOrgCmd.PersistentFlags().StringP("admin", "a", "", "the address for the admin account.")

	aprvOrgCmd.MarkPersistentFlagRequired("enodeId")
	aprvOrgCmd.MarkPersistentFlagRequired("ultPrnt")
	aprvOrgCmd.MarkPersistentFlagRequired("orgId")
	aprvOrgCmd.MarkPersistentFlagRequired("raftport")
	aprvOrgCmd.MarkPersistentFlagRequired("ip")
	aprvOrgCmd.MarkPersistentFlagRequired("port")
	aprvOrgCmd.MarkPersistentFlagRequired("guardAdd")
	aprvOrgCmd.MarkPersistentFlagRequired("guardKey")

	interCmd.AddCommand(aprvOrgCmd)

}
