/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package role

import (
	"strconv"

	"github.com/ijsingh82/qV2Cli/pkg/role"

	"fmt"

	"github.com/spf13/cobra"
)

// acctDetailsCmd represents the acctDetails command
var roleInfoIndex = &cobra.Command{
	Use:              "fromIndex",
	TraverseChildren: true,
	Short:            "get role details from role index.",
	Long: `    
/** @notice returns the role details for a passed role index
* @param _rIndex - unique identifier for the role being added
* @return role id
* @return org id
* @return access type
* @return bool to indicate if the role is a voter role
* @return bool to indicate if the role is active

qV2Cli role fromIndex [index]`,
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

		d, stat, err := role.FromIndex(int64(index))
		if err != nil {
			return err
		}
		fmt.Printf("The Role Details are:\n%+v\n", d)
		fmt.Printf("The Role status is: %+v\n", stat)
		return nil
	},
}

func init() {

	roleCmd.AddCommand(roleInfoIndex)

}
