/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package impl

import (
	"github.com/ijsingh82/qV2Cli/pkg/impl"

	"fmt"

	"github.com/spf13/cobra"
)

// getOrgsCmd represents the getOrgs command
var netStatCmd = &cobra.Command{
	Use:              "bootStat",
	TraverseChildren: true,
	Short:            "Get the network boot status.",
	Long: `    
/** @notice function to fetch network boot status
* @return bool network boot status
*/

qV2Cli impl bootStat`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(0), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {

		stat, err := impl.BootStat()
		if err != nil {
			return err
		}
		fmt.Printf("The Network Boot Status for Rpc is : %v\n", stat)
		return nil
	},
}

func init() {

	implCmd.AddCommand(netStatCmd)

}
