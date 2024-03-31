/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package acct

import (
	"fmt"
	"os"

	"github.com/ijsingh82/qV2Cli/pkg/accounts"

	"github.com/spf13/cobra"
)

// getBalKeyCmd Get Balance of account from chian using keystore.
var getBalKeyCmd = &cobra.Command{
	Use:              "balKey",
	TraverseChildren: true,
	Short:            "Get Balance of account from chian using keystore.",
	Long: `
Get Balance of account from chian using keystore.
NOTE use "" for no password.
	
qV2Cli acct balKey [keystorePath] [password]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
			return err
		}
		_, err := os.Stat(args[0])
		if err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		keyfilePath := args[0]
		passowrd := args[1]

		bal, err := accounts.GetBalKeyfile(keyfilePath, passowrd)
		if err != nil {
			return err
		}
		fmt.Println("The Balance is : ", bal)

		return nil

	},
}

func init() {

	acctCmd.AddCommand(getBalKeyCmd)

}
