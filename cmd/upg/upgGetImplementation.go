/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package upg

import (
	"fmt"

	"github.com/ijsingh82/qV2Cli/pkg/upg"

	"github.com/spf13/cobra"
)

// acctDetailsCmd represents the acctDetails command
var upgImplCmd = &cobra.Command{
	Use:              "impl",
	TraverseChildren: true,
	Args:             cobra.MatchAll(cobra.MaximumNArgs(0), cobra.OnlyValidArgs),
	Short:            "Get the current permission implementation contract.",
	Long: ` 
/** @notice function to fetch the current implementation address
* @return permissions implementation contract address
*/

qV2Cli upg impl`,
	RunE: func(cmd *cobra.Command, args []string) error {
		add, err := upg.Implementation()
		if err != nil {
			return err
		}
		fmt.Printf("The Permission Implementaion Address is : %v\n", add)
		return nil
	},
}

func init() {
	upgCmd.AddCommand(upgImplCmd)
}
