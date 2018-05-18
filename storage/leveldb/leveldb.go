package leveldb

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/lnquy/bc/block"
	"github.com/lnquy/bc/config"
	"github.com/lnquy/bc/storage"
	"github.com/syndtr/goleveldb/leveldb"
)

type lvlDBStorage struct {
	db *leveldb.DB
}

func NewStorage(conf config.LevelDB) (storage.Ledger, error) {
	db, err := leveldb.OpenFile(conf.DBFile, nil)
	if err != nil {
		return nil, err
	}
	return &lvlDBStorage{
		db: db,
	}, nil
}

func (l *lvlDBStorage) AddBlock(bl *block.Block) (uint64, error) {
	b, err := json.Marshal(bl)
	if err != nil {
		return 0, err
	}

	if err = l.db.Put(uint64ToBytes(bl.ID), b, nil); err != nil {
		return 0, err
	}
	return bl.ID, nil
}

func (l *lvlDBStorage)  GetBlock(id uint64) (*block.Block, error) {
	b, err := l.db.Get(uint64ToBytes(id), nil)
	if err != nil {
		return nil, err
	}
	bl := block.Block{}
	if err = json.Unmarshal(b, &bl); err != nil {
		return nil, err
	}
	return &bl, nil
}

func (l *lvlDBStorage) GetLatestBlock() (*block.Block, error) {
	nuBlocks := l.count()
	if nuBlocks <= 0 {
		return nil, fmt.Errorf("the blockchain is empty")
	}
	b, err := l.db.Get(uint64ToBytes(nuBlocks - 1), nil)
	if err != nil {
		return nil, err
	}
	bl := block.Block{}
	if err = json.Unmarshal(b, &bl); err != nil {
		return nil, err
	}
	return &bl, nil
}

func (l *lvlDBStorage) Dump() ([]*block.Block, error) {
	blocks := make([]*block.Block, 0)

	iter := l.db.NewIterator(nil, nil)
	defer iter.Release()
	for iter.Next() {
		bl := block.Block{}
		if err := json.Unmarshal(iter.Value(), &bl); err != nil {
			return nil, err
		}
		blocks = append(blocks, &bl)
	}
	if iter.Error() != nil {
		return nil, iter.Error()
	}
	sort.Sort(block.ByID(blocks))
	return blocks, nil
}

func (l *lvlDBStorage) DumpFromID(id uint64) ([]*block.Block, error) {
	blocks := make([]*block.Block, 0)

	iter := l.db.NewIterator(nil, nil)
	defer iter.Release()
	for ok := iter.Seek(uint64ToBytes(id)); ok; ok = iter.Next() {
		bl := block.Block{}
		if err := json.Unmarshal(iter.Value(), &bl); err != nil {
			return nil, err
		}
		blocks = append(blocks, &bl)
	}
	if iter.Error() != nil {
		return nil, iter.Error()
	}
	sort.Sort(block.ByID(blocks))
	return blocks, nil
}

func (l *lvlDBStorage) Close() error {
	return l.db.Close()
}

func (l *lvlDBStorage) count() uint64 {
	iter := l.db.NewIterator(nil, nil)
	defer iter.Release()
	var c uint64 = 0
	for iter.Next() {
		c++
	}
	if iter.Error() != nil {
		return 0
	}
	return c
}

func uint64ToBytes(v uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, v)
	return b
}
