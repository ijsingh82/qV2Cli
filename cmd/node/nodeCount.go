/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package node

import (
	"fmt"

	"github.com/ijsingh82/qV2Cli/pkg/node"

	"github.com/spf13/cobra"
)

// acctDetailsCmd represents the acctDetails command
var nodeCountCmd = &cobra.Command{
	Use:              "count",
	TraverseChildren: true,
	Short:            "Get the Node Count from the chain.",
	Long: `
	/** @notice returns the total number of enodes in the network
	* @return number of nodes
	*/
	
	qV2Cli node count`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(0), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {

		d, err := node.NodeCount()
		if err != nil {
			return err
		}
		fmt.Printf("The Nodes Count is : %v\n", d)

		return nil
	},
}

func init() {
	nodeCmd.AddCommand(nodeCountCmd)
}
