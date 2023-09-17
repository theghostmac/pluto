package unit_test

import (
	"bytes"
	"fmt"
	"github.com/theghostmac/pluto/internal/core/blockchain"
	"github.com/theghostmac/pluto/internal/core/transactions"
	"github.com/theghostmac/pluto/internal/core/utils"
	"reflect"
	"testing"
	"time"
)

func TestHeaderEncodeDecodeBinary(t *testing.T) {
	// Create a sample Header with values.
	header := blockchain.Header{
		Version:           1,
		PreviousBlockHash: utils.RandomHash(),
		Timestamp:         time.Now().UnixNano(),
		Height:            10,
		Nonce:             67890,
	}

	// Encode the Header to binary
	var encodedHeaderBuffer bytes.Buffer
	if err := header.EncodeBinary(&encodedHeaderBuffer); err != nil {
		t.Fatalf("Error encoding header: %v", err)
	}
	// Decode the encoded binary data back into the new Header.
	decodedHeader := blockchain.Header{}
	if err := decodedHeader.DecodeBinary(&encodedHeaderBuffer); err != nil {
		t.Fatalf("Error decoding header: %v", err)
	}

	// Compare original header with decoded header.
	if !headerEqual(header, decodedHeader) {
		t.Fatalf("Decoded header does not match the original header.")
	}
}

func headerEqual(h1, h2 blockchain.Header) bool {
	return h1.Version == h2.Version &&
		h1.PreviousBlockHash == h2.PreviousBlockHash &&
		h1.Timestamp == h2.Timestamp &&
		h1.Height == h2.Height &&
		h1.Nonce == h2.Nonce
}

func TestBlockEncodeDecodeBinary(t *testing.T) {
	// Create a sample Block with header and Transaction.
	block := blockchain.Block{
		Header: blockchain.Header{
			Version:           1,
			PreviousBlockHash: utils.Hash{},
			Timestamp:         time.Now().UnixNano(),
			Height:            15,
			Nonce:             67890,
		},
		Transactions: []transactions.Transactions{},
	}

	// Encode the Block to binary
	var encodedBlockBuffer bytes.Buffer
	if err := block.EncodeBinary(&encodedBlockBuffer); err != nil {
		t.Fatalf("Error encoding Block: %v", err)
	}

	// Create a new Block to decode into
	var decodedBlock blockchain.Block
	// Decode the binary data into the new Block
	if err := decodedBlock.DecodeBinary(&encodedBlockBuffer); err != nil {
		t.Fatalf("Error decoding Block: %v", err)
	}

	// Compare the original Block with the decoded Block
	if !reflect.DeepEqual(block, decodedBlock) {
		t.Errorf("Decoded Block does not match original Block.")
	}
	fmt.Print(block)
}
