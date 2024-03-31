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

// acctDetailsCmd represents the acctDetails command
var countCmd = &cobra.Command{
	Use:              "count",
	TraverseChildren: true,
	Short:            "Get the Number of accounts active on chain.",
	Long: `
Get the Number of accounts active on chain.
/** @notice returns the total number of accounts
* @return total number accounts
*/

qV2Cli [acct] [count]`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(0), cobra.OnlyValidArgs),

	RunE: func(cmd *cobra.Command, args []string) error {

		cnt, err := accounts.GetAccountCount()
		if err != nil {
			return err
		}
		fmt.Printf("The Accounts Count is : %v\n", cnt)
		return nil
	},
}

func init() {
	acctCmd.AddCommand(countCmd)
}
