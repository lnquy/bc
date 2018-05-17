package block

import (
	"time"
	"crypto/sha256"
	"encoding/binary"
	"github.com/lnquy/bc"
	"fmt"
)

type Block struct {
	ID           uint64
	Timestamp    uint64
	PreviousHash []byte
	Nonce        uint64
	Hash         []byte
	Data         []byte
}

func (b *Block) String() string {
	return fmt.Sprintf("ID: %d\nTimestamp: %d\nPreviousHash: %x\nNonce: %d\nHash: %x\nData: %s\n",
		b.ID, b.Timestamp, b.PreviousHash, b.Nonce, b.Hash, b.Data)
}

func GenesisBlock() *Block {
	gb := &Block{
		ID:           0,
		Timestamp:    uint64(time.Now().Unix()),
		PreviousHash: []byte{},
		Nonce:        0,
		Hash:         []byte{},
		Data:         []byte("I'm the mother of all. Please TRUST me!"),
	}
	mineBlock(gb)
	return gb
}

func NewBlock(latestID uint64, prevHash []byte, data []byte) *Block {
	b := &Block{
		ID:           latestID + 1,
		Timestamp:    uint64(time.Now().Unix()),
		PreviousHash: prevHash,
		Nonce:        0,
		Data:         data,
	}
	mineBlock(b)
	return b
}

func mineBlock(block *Block) {
	for {
		h := sha256.Sum256(getRawBlock(block))
		if !isValidHash(h) {
			block.Nonce++
			continue
		}

		block.Hash = h[:]
		return
	}
}

// ID - Timestamp - PreviousHash - Nonce - Data
func getRawBlock(block *Block) []byte {
	d := make([]byte, 0)
	bUint64 := make([]byte, 8)

	// ID
	binary.LittleEndian.PutUint64(bUint64, block.ID)
	d = append(d, bUint64...)
	// Timestamp
	binary.LittleEndian.PutUint64(bUint64, block.Timestamp)
	d = append(d, bUint64...)
	// Previous hash
	d = append(d, block.PreviousHash...)
	// Nonce
	binary.LittleEndian.PutUint64(bUint64, block.Nonce)
	d = append(d, bUint64...)
	// Data
	d = append(d, block.Data...)

	return d
}

func isValidHash(hash [32]byte) bool {
	zeros := 0
	for _, b := range hash {
		if b&0x0F != 0x00 {
			goto exit
		}
		zeros++
		if b&0xF0 != 0x00 {
			goto exit
		}
		zeros++
	}

exit:
	return zeros >= bc.BLOCKCHAIN_POW_DIFFICULTY
}
