/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package impl

import (
	"net"
	"strconv"

	"github.com/ijsingh82/qV2Cli/pkg/impl"

	"fmt"

	"github.com/spf13/cobra"
)

// getOrgsCmd represents the getOrgs command
var implConnAllowedCmd = &cobra.Command{
	Use:              "connAllow",
	TraverseChildren: true,
	Short:            "Check if connection is allowed form a node",
	Long: `
/** @notice checks if the node is allowed to connect or not
* @param _enodeId enode id
* @param _ip IP of node
* @param _port tcp port of node
* @return bool indicating if the node is allowed to connect or not
*/

qV2Cli impl connAllow [enodeId] [ip] [port]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(3)(cmd, args); err != nil {
			return err
		}

		_, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("not a number: %s", args[0])
		}
		if net.ParseIP(args[1]) == nil {
			return fmt.Errorf("not an ip: %s", args[0])
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		enodeId := args[0]
		ip := args[1]
		port, _ := strconv.Atoi(args[2])
		allowed, err := impl.ConnAllowed(enodeId, ip, uint16(port))
		if err != nil {
			return err
		}

		fmt.Printf("The Connections for Enode ID  %v, Ip address %v, Port %v is allowed %v\n", enodeId, ip, port, allowed)

		return nil
	},
}

func init() {
	// implConnAllowedCmd.PersistentFlags().StringP("enodeId", "e", "", "The node EnodeId to check certChain.")
	// implConnAllowedCmd.PersistentFlags().StringP("ip", "i", "", "The ip address for the node certChain.")
	// implConnAllowedCmd.PersistentFlags().Uint16P("port", "p", 0, "The ip address for the node certChain.")

	// implConnAllowedCmd.MarkPersistentFlagRequired("enodeId")
	// implConnAllowedCmd.MarkPersistentFlagRequired("ip")
	// implConnAllowedCmd.MarkPersistentFlagRequired("port")

	implCmd.AddCommand(implConnAllowedCmd)

}
