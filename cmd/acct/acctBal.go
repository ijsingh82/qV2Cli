/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package acct

import (
	"fmt"

	"github.com/ijsingh82/qV2Cli/pkg/accounts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// getBalCmd Get Balance of account from chian.
var getBalCmd = &cobra.Command{
	Use: "bal",
	// TraverseChildren: true,
	Short: "Get Balance of account from chian.",
	Long: `
Get Balance of account from chian.

qV2Cli acct bal [account]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		if common.IsHexAddress(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid address Specified: %s", args[0])
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		bal, err := accounts.GetBal(args[0])
		if err != nil {
			return err
		}
		fmt.Println("The Balance is : ", bal)

		return nil

	},
}

func init() {

	acctCmd.AddCommand(getBalCmd)

}
