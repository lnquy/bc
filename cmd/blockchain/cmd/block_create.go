package cmd

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/lnquy/blockchain/block"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new block",
	Long:  ``,
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
	latestBlock, err := glbLedger.GetLatestBlock()
	if err != nil {
		if err.Error() != "the blockchain is empty" {
			return fmt.Errorf("failed to get latest block: %s", err)
		}
		latestBlock = block.GenesisBlock()
		if _, err = glbLedger.AddBlock(latestBlock); err != nil {
			return fmt.Errorf("failed to create genesis block: %s", err)
		}
		log.Info("genesis block created")
	}

	bl := block.NewBlock(latestBlock.ID, latestBlock.Hash, []byte(data))
	if _, err = glbLedger.AddBlock(bl); err != nil {
		return fmt.Errorf("failed to create block: %s", err)
	}
	log.Infof("block created:\n%s", bl)
	return nil
}
