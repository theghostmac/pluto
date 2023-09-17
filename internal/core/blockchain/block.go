package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"github.com/theghostmac/pluto/internal/core/utils"
	"io"
)

// The EncodeBinary method is used to encode the Header structure into binary format.
// It writes the Version, PreviousBlockHash, Timestamp, Height, and Nonce fields of the Header structure to a Writer in little-endian order.
func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.PreviousBlockHash); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Height); err != nil {
		return err
	}

	return binary.Write(w, binary.LittleEndian, &h.Nonce)
}

// The DecodeBinary method is used to decode the Header structure from binary format.
// It reads the Version, PreviousBlockHash, Timestamp, Height, and Nonce fields of the Header structure from a Reader in little-endian order.
func (h *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.PreviousBlockHash); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
		return err
	}

	return binary.Read(r, binary.LittleEndian, &h.Nonce)
}

// BlockHasher hashes the Block.
func (b *Block) BlockHasher() utils.Hash {
	newBuffer := &bytes.Buffer{}
	err := b.Header.EncodeBinary(newBuffer)
	if err != nil {
		return [32]uint8{}
	}

	if b.Hash.IsZero() {
		b.Hash = utils.Hash(sha256.Sum256(newBuffer.Bytes()))
	}

	return b.Hash
}

// The EncodeBinary method is used to encode the Block structure into binary format.
// It first encodes the Header of the block, and then encodes each transaction in the Transactions field of the Block structure.
func (b *Block) EncodeBinary(w io.Writer) error {
	if err := b.Header.EncodeBinary(w); err != nil {
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.EncodeBinary(w); err != nil {
			return err
		}
	}

	return nil
}

// The DecodeBinary method is used to decode the Block structure from binary format.
// It first decodes the Header of the block, and then decodes each transaction in the Transactions field of the Block structure.
func (b *Block) DecodeBinary(r io.Reader) error {
	if err := b.Header.DecodeBinary(r); err != nil {
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.DecodeBinary(r); err != nil {
			return err
		}
	}

	return nil
}

//// The DecodeBinary method is used to decode the Block structure from binary format.
//// It first decodes the Header of the block, and then decodes each transaction in the Transactions field of the Block structure.
//func (b *Block) DecodeBinary(r io.Reader) error {
//	// Decode the Header
//	if err := b.Header.DecodeBinary(r); err != nil {
//		return err
//	}
//
//	// Decode the Transactions
//	for i := range b.Transactions {
//		if err := b.Transactions[i].DecodeBinary(r); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

func (b *Block) Hasher() utils.Hash {
	return utils.Hash{} // TODO
}
