package unit

import (
	"github.com/theghostmac/pluto/internal/core/blockchain"
	"github.com/theghostmac/pluto/internal/core/transactions"
	"github.com/theghostmac/pluto/internal/core/utils"
	"time"
)

func RandomBlock() blockchain.Block {
	return blockchain.Block{
		Header: blockchain.Header{
			Version:           1,
			PreviousBlockHash: utils.RandomHash(),
			Timestamp:         time.Now().UnixNano(),
			Height:            10,
		},
		Transactions: []transactions.Transactions{},
		Hash:         utils.Hash{},
	}
}
