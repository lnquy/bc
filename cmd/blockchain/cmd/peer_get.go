package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var peerGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a peer detail or list all peers",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func init() {
	peerCmd.AddCommand(peerGetCmd)
}
