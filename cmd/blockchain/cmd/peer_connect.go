package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to blockchain network",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("connect called")
	},
}

func init() {
	peerCmd.AddCommand(connectCmd)
}
