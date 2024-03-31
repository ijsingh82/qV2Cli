/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package impl

import (
	"github.com/ijsingh82/qV2Cli/pkg/impl"
	"github.com/ijsingh82/qV2Cli/pkg/utils"

	"fmt"

	"github.com/spf13/cobra"
)

// getOrgsCmd represents the getOrgs command
var pendOpCmd = &cobra.Command{
	Use:              "pendOp",
	TraverseChildren: true,
	Short:            "Get any pending ops for org.",
	Long: `    
/** @notice function to fetch detail of any pending approval activities
for network admin organization
* @param _orgId unique id of the organization to which the account belongs
*/

qV2Cli impl pendingOp [orgId]`,
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),

	RunE: func(cmd *cobra.Command, args []string) error {

		orgId := args[0]
		op, err := impl.GetPendingOp(orgId)
		if err != nil {
			return err
		}
		t, err := op.PendingOpStatusString()
		if err != nil {
			return err
		}
		fmt.Printf("The Pending Operation Details are :\n %+v\n", utils.PPrint(op))
		fmt.Printf("The Operation Type is :\n%v \n", t)

		return nil
	},
}

func init() {

	implCmd.AddCommand(pendOpCmd)

}
