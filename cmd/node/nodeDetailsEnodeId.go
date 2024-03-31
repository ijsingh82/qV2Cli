/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package node

import (
	"fmt"

	"github.com/ijsingh82/qV2Cli/pkg/node"
	"github.com/ijsingh82/qV2Cli/pkg/utils"

	"github.com/spf13/cobra"
)

// acctDetailsCmd represents the acctDetails command
var nodeFromEnodeCmd = &cobra.Command{
	Use:              "fromId",
	TraverseChildren: true,
	Short:            "Get the Node Details from Enode Id",
	Long: `    /** @notice fetches the node details given an enode id
	* @param _enodeId full enode id
	* @return org id
	* @return enode id
	* @return status of the node
	*/
	
	qV2Cli node fromId [enodeId]`,
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		nodeId := args[0]

		d, status, err := node.FromEnode(nodeId)
		if err != nil {
			return err
		}
		fmt.Printf("The Nodes Details is :\n %+v\n", utils.PPrint(d))
		fmt.Printf("The Nodes Status is :%+v\n", status)

		return nil
	},
}

func init() {
	nodeFromEnodeCmd.PersistentFlags().StringP("nodeId", "n", "", "The address to verify for validity for the org on the certChain.")
	nodeFromEnodeCmd.MarkPersistentFlagRequired("nodeId")

	nodeCmd.AddCommand(nodeFromEnodeCmd)

}
