/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package node

import (
	"fmt"

	"github.com/ijsingh82/qV2Cli/pkg/node"

	"github.com/spf13/cobra"
)

// crtNodeCmd represents the crtNode command
var crtNodeCmd = &cobra.Command{
	Use:              "create",
	TraverseChildren: true,
	Short:            "Create a New Node with account info for a new org.",
	Long: `
Create a New Node with account info for a new org.

qV2Cli node create [nodeName] [password]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
			return fmt.Errorf("add args %s", err.Error())
		}
		if len(args[0]) > 11 {
			return fmt.Errorf("name must be less than 10 in length")

		}
		if len(args[1]) < 11 {
			return fmt.Errorf("password must be greater than 10 in length")

		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		pass := args[0]
		instName := args[0]
		v, err := node.Create(instName, pass)
		if err != nil {
			return err
		}
		fmt.Println("The Node Validator account Address in keystore : ", v.ValidatorAccount)
		fmt.Println("The Node path is :", v.NodePath)

		return nil
	},
}

func init() {

	nodeCmd.AddCommand(crtNodeCmd)

}
