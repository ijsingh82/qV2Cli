/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package impl

import (
	"strconv"

	"github.com/ijsingh82/qV2Cli/pkg/impl"

	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// getOrgsCmd represents the getOrgs command
var implTxAllowedCmd = &cobra.Command{
	Use:              "txAllow",
	TraverseChildren: true,
	Short:            "Check if a transaction is allowed form an account.",
	Long: `    
/** @notice checks if the account is allowed to transact or not
* @param _sender source account
* @param _target target account
* @param _value value being transferred
* @param _gasPrice gas price
* @param _gasLimit gas limit
* @param _payload payload for transactions on contracts
* @return bool indicating if the account is allowed to transact or not
*/

qV2Cli impl txAllow [sender] [target] [value] [payload]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(4)(cmd, args); err != nil {
			return err
		}
		if !common.IsHexAddress(args[0]) {
			return fmt.Errorf("invalid sender address Specified: %s", args[0])
		}
		if !common.IsHexAddress(args[1]) {
			return fmt.Errorf("invalid target address Specified: %s", args[0])
		}
		_, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("not a number: %s", args[0])
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		fromAdd := args[0]
		toAdd := args[1]
		value, _ := strconv.Atoi(args[3])
		data := args[3]

		allowed, err := impl.IsTxAllowed(fromAdd, toAdd, int64(value), []byte(data))
		if err != nil {
			return err
		}
		fmt.Printf("The transaction is allowed %v\n", allowed)

		return nil
	},
}

func init() {

	implCmd.AddCommand(implTxAllowedCmd)

}
