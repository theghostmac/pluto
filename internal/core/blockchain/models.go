package blockchain

import (
	"github.com/theghostmac/pluto/internal/core/transactions"
	"github.com/theghostmac/pluto/internal/core/utils"
)

type Header struct {
	Version           uint32
	PreviousBlockHash utils.Hash
	Timestamp         int64
	Height            uint32
	Nonce             uint64
}

type Block struct {
	Header
	Transactions []transactions.Transactions
}

/*
# Generic elements of a blockchain
The generic elements of a blockchain include the following:
1. The address
2. Transactions
3. Blocks
A single block is further expounded into different constituent parts.
All of these make up the execution layer of a blockchain.
I have already done the implementation of this in this article.
Let's walk through them in a high level fashion:
1. The Block header - which has the following parts:
	- Previous block header's hash (except the genesis block)
	- Nonce (the contract wallet and the external/user wallet have different nonces)
	- Timestamp (the time of the transaction)
	- Height
	- Merkle root (a bit complex, but we will look at them later).
2. Block body - which has the following part:
	- Transactions.

Headers are important for saving state.
*/
