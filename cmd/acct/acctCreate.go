/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package acct

import (
	"fmt"

	"github.com/ijsingh82/qV2Cli/pkg/accounts"

	"github.com/spf13/cobra"
)

// crtAccCmd represents the crtAcc command
var crtAccCmd = &cobra.Command{
	Use:              "create",
	TraverseChildren: true,
	Short:            "Create a New account for use.",
	Long: `
Create a New account info for use. Password must be more than 10 chars. Name must be less than 10.

qV2Cli acct create [name] [password]`,
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
		name := args[0]
		pass := args[1]

		acc0, err := accounts.CreateAccount(pass, name)
		if err != nil {
			return err
		}
		fmt.Println("The New Account is in the keystore.", acc0.Address)
		fmt.Println("The New Account is in the keystore path.", acc0.URL.Path)

		return nil
	},
}

func init() {

	acctCmd.AddCommand(crtAccCmd)

}
