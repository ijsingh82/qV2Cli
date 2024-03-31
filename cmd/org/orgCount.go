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
var orgCountCmd = &cobra.Command{
	Use:              "count",
	TraverseChildren: true,
	Args:             cobra.MatchAll(cobra.MaximumNArgs(0), cobra.OnlyValidArgs),
	Short:            "Get the count of orgs on the chain.",
	Long: `  
/** @notice returns the total number of orgs in the network
* @return master org id
*/


qV2Cli org count`,
	RunE: func(cmd *cobra.Command, args []string) error {

		cnt, err := org.Count()
		if err != nil {
			return err
		}
		fmt.Printf("The Number Of Orgs is : %v\n", cnt)

		return nil
	},
}

func init() {

	orgCmd.AddCommand(orgCountCmd)

}
