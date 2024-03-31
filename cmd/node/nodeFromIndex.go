/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package node

import (
	"fmt"
	"strconv"

	"github.com/ijsingh82/qV2Cli/pkg/node"

	"github.com/spf13/cobra"
)

// acctDetailsCmd represents the acctDetails command
var nodeFromIndexCmd = &cobra.Command{
	Use:              "fromIndex",
	TraverseChildren: true,
	Short:            "Get the Node Details from an index.",
	Long: `    
/** @notice fetches the node details given the index of the enode
* @param _nodeIndex node index
* @return org id
* @return enode id
* @return ip of the node
* @return port of the node
* @return raftport of the node
* @return status of the node
*/


qV2Cli node fromIndex [index]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		_, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("not a number: %s", args[0])
		}

		return nil

	},
	RunE: func(cmd *cobra.Command, args []string) error {
		index, _ := strconv.Atoi(args[0])

		d, stat, err := node.FromIndex(int64(index))
		if err != nil {
			return err
		}
		fmt.Printf("The Nodes Details is :\n%+v\n", d)
		fmt.Printf("The Nodes Status is :\n%+v\n", stat)

		return nil
	},
}

func init() {

	nodeCmd.AddCommand(nodeFromIndexCmd)

}
