package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var peerCmd = &cobra.Command{
	Use:   "peer",
	Short: "Manage peers in blockchain network",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("peer called")
	},
}

func init() {
	rootCmd.AddCommand(peerCmd)
}
