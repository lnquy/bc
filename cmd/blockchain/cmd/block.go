package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "Manage the blocks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("block called")
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)
}
