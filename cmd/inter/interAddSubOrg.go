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
var addSubOrgCmd = &cobra.Command{
	Use:              "addSubOrg",
	TraverseChildren: true,
	Short:            "Add a SubOrg to an org.",
	Long: `    /** @notice interface to add sub org under an org
	* @param _pOrgId parent org id under which the sub org is being added
	* @param _orgId unique id for the sub organization
	* @param _enodeId enode id linked to the sjb organization
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
		ultPrnt, err := cmd.Flags().GetString("ultPrnt")
		if err != nil {
			return err
		}
		r, err := inter.AddSubOrg(key, guardAdd, enodeId, ip, orgId, ultPrnt, port, raftport)
		if err != nil {
			return err
		}

		fmt.Printf("This is the transaction receipt for Adding the sub org:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil
	},
}

func init() {

	addSubOrgCmd.PersistentFlags().StringP("guardAdd", "g", "", "The guardian Address authorzied.")
	addSubOrgCmd.PersistentFlags().StringP("guardKey", "k", "", "The guardian private key no 0x authorzied.")
	addSubOrgCmd.PersistentFlags().StringP("enodeId", "e", "", "The enode address for the new node add.")
	addSubOrgCmd.PersistentFlags().StringP("ip", "i", "", "The ip address for the new node.")
	addSubOrgCmd.PersistentFlags().Uint16P("port", "p", 0, "The port for the new node.")
	addSubOrgCmd.PersistentFlags().Uint16P("raftport", "t", 0, "The port for the new node.")
	addSubOrgCmd.PersistentFlags().StringP("orgId", "o", "", "The orgId of the new org.")
	addSubOrgCmd.PersistentFlags().StringP("ultPrnt", "u", "", "The ultimate parents org Id of the sub org.")

	addSubOrgCmd.MarkPersistentFlagRequired("enodeId")
	addSubOrgCmd.MarkPersistentFlagRequired("ultPrnt")
	addSubOrgCmd.MarkPersistentFlagRequired("orgId")
	addSubOrgCmd.MarkPersistentFlagRequired("raftport")
	addSubOrgCmd.MarkPersistentFlagRequired("ip")
	addSubOrgCmd.MarkPersistentFlagRequired("port")
	addSubOrgCmd.MarkPersistentFlagRequired("guardAdd")
	addSubOrgCmd.MarkPersistentFlagRequired("guardKey")

	interCmd.AddCommand(addSubOrgCmd)

}
