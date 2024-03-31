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
var upgInterCmd = &cobra.Command{
	Use:              "inter",
	TraverseChildren: true,
	Args:             cobra.MatchAll(cobra.MaximumNArgs(0), cobra.OnlyValidArgs),
	Short:            "Get the interface address from the upgradable contract.",
	Long: `
/** @notice function to fetch the current interface contract address
* @return permissions interface contract address
*/

qV2Cli upg impl`,
	RunE: func(cmd *cobra.Command, args []string) error {

		add, err := upg.Interface()
		if err != nil {
			return err
		}
		fmt.Printf("The Permission Interface Address is : %v\n", add)
		return nil
	},
}

func init() {
	upgCmd.AddCommand(upgInterCmd)
}
