package cryptography

import "crypto/ecdsa"

type PrivateKey struct {
	privateKey *ecdsa.PrivateKey
}

type PublicKey struct {
	publicKey *ecdsa.PublicKey
}

type Signature struct {
}
