package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Hash [32]uint8

// HashFromBytes is used to hash the previous block.
func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		panic(fmt.Sprintf("Invalid input length. Expected byte slice length of 32, but got %d.", len(b)))
	}

	var value Hash
	copy(value[:], b)
	return value
}

func (h Hash) IsZero() bool {
	for i := 0; i < 32; i++ {
		if h[i] != 0 {
			return false
		}
	}
	return true
}

func (h Hash) ToSlice() []byte {
	b := make([]byte, 32)
	for i := 0; i < 32; i++ {
		b[i] = h[i]
	}
	return b
}

func (h Hash) String() string {
	return hex.EncodeToString(h.ToSlice())
}

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	_, err := rand.Read(token)
	if err != nil {
		return nil
	}
	return token
}

func RandomHash() Hash {
	return HashFromBytes(RandomBytes(32))
}
