package blockchain

import (
	"github.com/theghostmac/pluto/internal/core/transactions"
	"github.com/theghostmac/pluto/internal/core/utils"
)

// Header represents the header of a blockchain Block.
// It contains metadata and essential information about the Block.
type Header struct {
	// The Version number indicates the format and rules used to interpret the block data.
	// Different versions of the blockchain may have different rules or features,
	// and this field helps identify which version a block adheres to.
	Version           uint32
	PreviousBlockHash utils.Hash
	Timestamp         int64
	Height            uint32
	Nonce             uint64
}

// Block represents a complete blockchain block, and it inherits all the fields from the Header struct.
type Block struct {
	Header
	Transactions []transactions.Transactions

	// Cached version of the Header's hash
	Hash utils.Hash
}
