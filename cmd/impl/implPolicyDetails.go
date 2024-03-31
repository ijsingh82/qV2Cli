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
var implPolicyCmd = &cobra.Command{
	Use:              "policy",
	TraverseChildren: true,
	Short:            "Get the network Policy.",
	Long: `    
/** @notice returns various permissions policy related parameters
* @return adminOrg admin org id
* @return adminRole default network admin role
* @return orgAdminRole default org admin role
* @return networkBoot network boot status
*/.


qV2Cli impl policy`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(0), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {

		p, err := impl.GetPolicy()
		if err != nil {
			return err
		}
		fmt.Printf("The Policy Details are : \n%+v\n", p)

		return nil
	},
}

func init() {

	implCmd.AddCommand(implPolicyCmd)

}
