package utils

import (
	"crypto/rand"
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
