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
var addSubOrgKeyCmd = &cobra.Command{
	Use:              "addSubOrgKey",
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
		ultPrnt, err := cmd.Flags().GetString("ultPrnt")
		if err != nil {
			return err
		}

		r, err := inter.AddSubOrgKey(keyfile, password, enodeId, ip, orgId, ultPrnt, port, raftport)
		if err != nil {
			return err
		}

		fmt.Printf("This is the transaction receipt for Adding the sub org:\n %+v\n", utils.PPrint(r))
		fmt.Println()
		return nil
	},
}

func init() {

	addSubOrgKeyCmd.PersistentFlags().StringP("enodeId", "e", "", "The enode address for the new node add.")
	addSubOrgKeyCmd.PersistentFlags().StringP("ip", "i", "", "The ip address for the new node.")
	addSubOrgKeyCmd.PersistentFlags().Uint16P("port", "p", 0, "The port for the new node.")
	addSubOrgKeyCmd.PersistentFlags().Uint16P("raftport", "r", 0, "The port for the new node.")
	addSubOrgKeyCmd.PersistentFlags().StringP("keyfile", "f", "", "The path to the keyfile for authorization.")
	addSubOrgKeyCmd.PersistentFlags().StringP("password", "s", "", "The password for the keyfile for authorization(optional)")
	addSubOrgKeyCmd.PersistentFlags().StringP("orgId", "o", "", "The orgId of the new org.")
	addSubOrgKeyCmd.PersistentFlags().StringP("ultPrnt", "u", "", "The ultimate parents org Id of the sub org.")

	addSubOrgKeyCmd.MarkPersistentFlagRequired("enodeId")
	addSubOrgKeyCmd.MarkPersistentFlagRequired("ultPrnt")
	addSubOrgKeyCmd.MarkPersistentFlagRequired("orgId")
	addSubOrgKeyCmd.MarkPersistentFlagRequired("raftport")
	addSubOrgKeyCmd.MarkPersistentFlagRequired("ip")
	addSubOrgKeyCmd.MarkPersistentFlagRequired("port")
	addSubOrgKeyCmd.MarkPersistentFlagRequired("keyfile")

	interCmd.AddCommand(addSubOrgKeyCmd)

}
