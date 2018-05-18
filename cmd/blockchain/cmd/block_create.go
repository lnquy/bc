package cmd

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/lnquy/bc/block"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new block",
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		data := ""
		if len(args) > 0 {
			data = args[0]
		}
		return createBlock(data)
	},
}

func init() {
	blockCmd.AddCommand(createCmd)
}

func createBlock(data string) error {
	latestBlock, err := ledger.GetLatestBlock()
	if err != nil && err.Error() == "the blockchain is empty" {
		latestBlock = block.GenesisBlock()
		if _, err = ledger.AddBlock(latestBlock); err != nil {
			return fmt.Errorf("failed to create genesis block: %s", err)
		}
		log.Info("genesis block created")
	}

	bl := block.NewBlock(latestBlock.ID, latestBlock.Hash, []byte(data))
	if _, err = ledger.AddBlock(bl); err != nil {
		return fmt.Errorf("failed to create block: %s", err)
	}
	log.Infof("block created:\n%s", bl)
	return nil
}
