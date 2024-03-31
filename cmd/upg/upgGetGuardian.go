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
var upgGuardCmd = &cobra.Command{
	Use:              "guard",
	TraverseChildren: true,
	Args:             cobra.MatchAll(cobra.MaximumNArgs(0), cobra.OnlyValidArgs),
	Short:            "Get the guradian info from the upgradable contract.",
	Long: ` 
/** @notice function to fetch the guardian account address
* @return _guardian guardian account address
*/

qV2Cli upg guard`,
	RunE: func(cmd *cobra.Command, args []string) error {
		add, err := upg.Guardian()
		if err != nil {
			return err
		}
		fmt.Printf("The Guardian Address is : %v\n", add)
		return nil
	},
}

func init() {
	upgCmd.AddCommand(upgGuardCmd)
}
