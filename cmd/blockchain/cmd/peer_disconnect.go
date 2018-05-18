package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect from blockchain network",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("disconnect called")
	},
}

func init() {
	peerCmd.AddCommand(disconnectCmd)
}
