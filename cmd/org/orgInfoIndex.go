/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/ijsingh82/qV2Cli/pkg/org"

	"github.com/spf13/cobra"
)

// acctDetailsCmd represents the acctDetails command
var orgInfoIndexCmd = &cobra.Command{
	Use:              "fromIndex",
	TraverseChildren: true,
	Short:            "Get the org Details from an index.",
	Long: ` 
/** @notice returns org info for a given org index
* @param _orgIndex org index
* @return org id
* @return parent org id
* @return ultimate parent id
* @return level in the org tree
* @return status

qV2Cli org fromIndex [index]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		_, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("not a number: %s", args[0])
		}

		return nil

	}, RunE: func(cmd *cobra.Command, args []string) error {
		index, _ := strconv.Atoi(args[0])

		d, stat, err := org.FromIndex(int64(index))
		if err != nil {
			return err
		}
		fmt.Printf("The Org info is :\n%+v\n", d)
		fmt.Printf("The Org status is : %v\n", stat)
		return nil
	},
}

func init() {

	orgCmd.AddCommand(orgInfoIndexCmd)

}
