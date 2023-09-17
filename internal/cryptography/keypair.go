package cryptography

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/theghostmac/pluto/internal/core/utils"
)

func GeneratePrivateKey() (PrivateKey, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return PrivateKey{}, err
	}

	return PrivateKey{
		privateKey: key,
	}, nil
}

func (pk PrivateKey) GeneratePublicKey() PublicKey {
	return PublicKey{
		publicKey: &pk.privateKey.PublicKey,
	}
}

func (pk PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(pk.publicKey, pk.publicKey.X, pk.publicKey.Y)
}

func (pk PublicKey) Address() utils.Address {
	hash := sha256.Sum256(pk.ToSlice())
	return utils.AddressFromBytes(hash[len(hash)-20:])
}
