package cmd

import (
	"fmt"
	"strconv"

	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a block detail or list all blocks",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return getAllBlocks()
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid block ID: %s", args[0])
		}
		return getBlockDetail(uint64(id))
	},
}

func init() {
	blockCmd.AddCommand(getCmd)
}

func getBlockDetail(id uint64) error {
	bl, err := glbLedger.GetBlock(id)
	if err != nil {
		return fmt.Errorf("failed to get block #%d: %s", id, err)
	}
	log.Infof("%s", bl.ID, bl)
	return nil
}

func getAllBlocks() error {
	blocks, err := glbLedger.Dump()
	if err != nil {
		return fmt.Errorf("failed to get all blocks: %s", err)
	}
	for _, bl := range blocks {
		log.Info(bl)
	}
	log.Infof("total: %d block(s)", len(blocks))
	return nil
}
