package unit

import (
	"fmt"
	"github.com/theghostmac/pluto/internal/cryptography"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	// Generate a private key
	privateKey, err := cryptography.GeneratePrivateKey()
	if err != nil {
		t.Fatalf("Error generating private key: %v", err)
	}

	// Generate Public Key from Private Key
	publicKey := privateKey.GeneratePublicKey()

	// Generate Address from public Key
	address := publicKey.Address()
	fmt.Println(address)
}
