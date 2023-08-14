package blockchain

import (
	"encoding/binary"
	"io"
)

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

/* ----> More readable way to create these:
package blockchain

import (
	"encoding/binary"
	"io"
)

func (h *Header) encodeDecodeBinary(wr io.ReadWriter, binaryFunc func(io.ReadWriter, binary.ByteOrder, interface{}) error) error {
	order := binary.LittleEndian
	err := binaryFunc(wr, order, &h.Version)
	if err != nil {
		return err
	}
	err = binaryFunc(wr, order, &h.PreviousBlockHash)
	if err != nil {
		return err
	}
	err = binaryFunc(wr, order, &h.Timestamp)
	if err != nil {
		return err
	}
	err = binaryFunc(wr, order, &h.Height)
	if err != nil {
		return err
	}
	return binaryFunc(wr, order, &h.Nonce)
}

func (h *Header) EncodeBinary(w io.Writer) error {
	return h.encodeDecodeBinary(w, binary.Write)
}

func (h *Header) DecodeBinary(r io.Reader) error {
	return h.encodeDecodeBinary(r, binary.Read)
}


*/
