package utils

import "fmt"

type Hash [32]uint8

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		panic(fmt.Sprintf("Invalid input length. Expected byte slice length of 32, but got %d.", len(b)))
	}

	var value Hash
	copy(value[:], b)
	return value
}
