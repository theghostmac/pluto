package cryptography

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/theghostmac/pluto/internal/core/utils"
)

func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	panic(err)

	return PrivateKey{
		privateKey: key,
	}
}

func (pk PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		publicKey: &pk.privateKey.PublicKey,
	}
}

func (pk PublicKey) ToSlice() []bytes {
	return elliptic.MarshalCompressed(pk.publicKey, pk.publicKey.X, pk.publicKey.Y)
}

func (pk PublicKey) Address() utils.Address {
	hash := sha256.Sum256(pk.ToSlice())
	return utils.Address{} // TODO
}
