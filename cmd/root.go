/*
Copyright © 2024 Inderjit Singh

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	_ "embed"
	"log"

	"github.com/ijsingh82/qV2Cli/pkg/config"

	"os"

	"github.com/spf13/cobra"
)

// var RpcConnection string
var Conf *config.Config

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "qV2Cli",
	Short: "Cli to interact with the quorum permission V2 Contracts.",
	Long: `Cli to interact with the quorum permission V2 Contracts.
	Long descriptions are from the solidity docs.`,
	ValidArgs: []string{"node"},
	Args:      cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	// // Uncomment the following line if your bare application
	// // has an action associated with it:
	// 	RunE: func(cmd *cobra.Command, args []string) error {
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	RootCmd.CompletionOptions.DisableDefaultCmd = true

	b, err := os.ReadFile("./permission-config.json")
	if err != nil {
		log.Fatalf("error with load err : %v", err)

	}
	Con, err := config.Load(b)
	if err != nil {
		log.Fatalf("error with load err : %v", err)
	}
	Conf = Con
}
