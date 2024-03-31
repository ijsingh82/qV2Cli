/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/ijsingh82/qV2Cli/pkg/org"

	"github.com/spf13/cobra"
)

// acctDetailsCmd represents the acctDetails command
var orgInfoIdCmd = &cobra.Command{
	Use:              "fromId",
	TraverseChildren: true,
	Short:            "Get the org Details from a org id. ",
	Long: ` 
/** @notice returns org info for a given org id
* @param _orgId org id
* @return org id
* @return parent org id
* @return ultimate parent id
* @return level in the org tree
* @return status
*/

qV2Cli org fromId`,
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		orgId := args[0]

		d, stat, err := org.FromID(orgId)
		if err != nil {
			return err
		}

		fmt.Printf("The Org info is :\n%+v\n", d)
		fmt.Printf("The Org status is : %v\n", stat)
		return nil
	},
}

func init() {

	orgCmd.AddCommand(orgInfoIdCmd)

}
