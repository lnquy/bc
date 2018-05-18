package bigcache

import (
	"encoding/json"
	"fmt"

	"github.com/allegro/bigcache"
	"github.com/lnquy/blockchain/block"
	"github.com/lnquy/blockchain/ledger"
)

type cache struct {
	bigCache *bigcache.BigCache
}

func NewCache() (ledger.Ledger, error) {
	bc, err := bigcache.NewBigCache(bigcache.DefaultConfig(0))
	if err != nil {
		return nil, err
	}
	return &cache{
		bigCache: bc,
	}, nil
}

func (c *cache) AddBlock(bl *block.Block) (uint64, error) {
	b, err := json.Marshal(bl)
	if err != nil {
		return 0, err
	}

	if err = c.bigCache.Set(fmt.Sprintf("%d", bl.ID), b); err != nil {
		return 0, err
	}
	return bl.ID, nil
}

func (c *cache) GetBlock(id uint64) (*block.Block, error) {
	b, err := c.bigCache.Get(fmt.Sprintf("%d", id))
	if err != nil {
		return nil, err
	}
	bl := block.Block{}
	if err = json.Unmarshal(b, &bl); err != nil {
		return nil, err
	}
	return &bl, nil
}

func (c *cache) GetLatestBlock() (*block.Block, error) {
	if c.bigCache.Len() <= 0 {
		return nil, fmt.Errorf("the blockchain is empty")
	}
	b, err := c.bigCache.Get(fmt.Sprintf("%d", c.bigCache.Len()-1))
	if err != nil {
		return nil, err
	}
	bl := block.Block{}
	if err = json.Unmarshal(b, &bl); err != nil {
		return nil, err
	}
	return &bl, nil
}

func (c *cache) Dump() ([]*block.Block, error) {
	chainLength := uint64(c.bigCache.Len())
	blocks := make([]*block.Block, chainLength)

	for i := uint64(0); i < chainLength; i++ {
		bl, err := c.GetBlock(i)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, bl)
	}
	return blocks, nil
}

func (c *cache) DumpFromID(id uint64) ([]*block.Block, error) {
	chainLength := uint64(c.bigCache.Len())
	if id >= chainLength {
		return nil, fmt.Errorf("invalid id. Latest block ID is: %d", chainLength-1)
	}

	blocks := make([]*block.Block, chainLength-id)
	for i := id; i < chainLength; i++ {
		bl, err := c.GetBlock(i)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, bl)
	}
	return blocks, nil
}

func (c *cache) Close() error {
	return c.bigCache.Reset()
}
