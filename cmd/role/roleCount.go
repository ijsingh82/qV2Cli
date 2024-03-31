/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package role

import (
	"fmt"

	"github.com/ijsingh82/qV2Cli/pkg/role"

	"github.com/spf13/cobra"
)

// acctDetailsCmd represents the acctDetails command
var roleCountCmd = &cobra.Command{
	Use:              "count",
	TraverseChildren: true,
	Args:             cobra.MatchAll(cobra.MaximumNArgs(0), cobra.OnlyValidArgs),
	Short:            "Get the role Count.",
	Long: ` 
/** @notice returns the total number of roles in the network
* @return total number of roles
*/

qV2Cli role count`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cnt, err := role.Count()
		if err != nil {
			return err
		}
		fmt.Printf("The Number Of Roles is : %v\n", cnt)
		return nil
	},
}

func init() {
	roleCmd.AddCommand(roleCountCmd)
}
