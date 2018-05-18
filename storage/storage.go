package storage

import "github.com/lnquy/bc/block"

type Ledger interface {
	AddBlock(*block.Block) (uint64, error)
	GetBlock(uint64) (*block.Block, error)
	GetLatestBlock() (*block.Block, error)
	Dump() ([]*block.Block, error)
	DumpFromID(uint64) ([]*block.Block, error)
	Close() error
}
