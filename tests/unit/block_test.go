package unit_test

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/theghostmac/pluto/internal/core/blockchain"
	"github.com/theghostmac/pluto/internal/core/utils"
	"github.com/theghostmac/pluto/tests/unit"
	"reflect"
	"testing"
	"time"
)

// Note to self
// Use already-made random block generator function: block := unit.RandomBlock()

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
	block := unit.RandomBlock()

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

		// Print details about the mismatched fields
		if block.Header.Version != decodedBlock.Header.Version {
			t.Errorf("Version: Original=%d, Decoded=%d", block.Header.Version, decodedBlock.Header.Version)
		}

		if block.Header.PreviousBlockHash != decodedBlock.Header.PreviousBlockHash {
			t.Errorf("PreviousBlockHash: Original=%s, Decoded=%s", block.Header.PreviousBlockHash, decodedBlock.Header.PreviousBlockHash)
		}

		if block.Header.Timestamp != decodedBlock.Header.Timestamp {
			t.Errorf("Timestamp: Original=%d, Decoded=%d", block.Header.Timestamp, decodedBlock.Header.Timestamp)
		}

		if block.Header.Height != decodedBlock.Header.Height {
			t.Errorf("Height: Original=%d, Decoded=%d", block.Header.Height, decodedBlock.Header.Height)
		}
	}

	fmt.Print("The block is ", block)
}

func TestHeaderEncodeBinary(t *testing.T) {
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
}

func TestHeaderDecodeBinary(t *testing.T) {
	encodedData := []byte{
		0x01, 0x00, 0x00, 0x00, // Version (uint32, little-endian)
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, // PreviousBlockHash (20 bytes)
		0x18, 0x00, 0x00, 0x00, // Timestamp (int64, little-endian)
		0x0A, 0x00, 0x00, 0x00, // Height (uint32, little-endian)
		0x32, 0x54, 0x76, 0x98, 0xBA, 0xDC, 0xFE, 0x00, // Nonce (uint64, little-endian)
	}

	// Create a new Header to decode into
	var decodedHeader blockchain.Header

	// Decode the binary data into the new Header
	if err := decodedHeader.DecodeBinary(bytes.NewReader(encodedData)); err != nil {
		t.Fatalf("Error decoding header: %v", err)
	}
}

func TestBlockHash(t *testing.T) {
	block := unit.RandomBlock()
	hash := block.BlockHasher()
	fmt.Println(hash)
	assert.False(t, hash.IsZero(), "Block has should not be zero")
}
