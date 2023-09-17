package utils

import "fmt"

type Address [20]uint8

func (a Address) ToSlice() []byte {
	b := make([]byte, 20)
	for i := 0; i < 20; i++ {
		b[i] = a[i]
	}
	return b
}

func AddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		panic(fmt.Sprintf("Invalid input length. Expected byte slice length of 20, but got %d.", len(b)))
	}

	var value Address
	copy(value[:], b)
	return value
}
