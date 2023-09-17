package unit

import (
	"crypto/elliptic"
	"github.com/theghostmac/pluto/internal/cryptography"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	// Generate a private key
	privateKey := cryptography.GeneratePrivateKey()
	// Check if the private key is nil
	if privateKey.privateKey == nil {
		t.Errorf("GeneratePrivateKey() returned a nil private key")
	}
}

func TestPrivateKeyPublicKey(t *testing.T) {
	// Generate a private key
	privateKey := cryptography.GeneratePrivateKey()
	// Get the corresponding public key
	publicKey := privateKey.PublicKey()

	// Check if the public key is nil.
	if publicKey.publicKey == nil {
		t.Errorf("PrivateKey.PublicKey() returned a nil public key")
	}

	// Check if the public key's curve matches the expected curve (P-256)
	if publicKey.Curve != elliptic.P256() {
		t.Errorf("PublicKey's curve is not P-256")
	}

	// Check if the public key's X and Y coordinates are not nil
	if publicKey.publicKey.X == nil || publicKey.publicKey.Y == nil {
		t.Errorf("PublicKey's X or Y coordinates are nil")
	}
}
