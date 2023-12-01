package edwards25519

import (
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"steve.care/network/domain/encryptors"
)

var curve = edwards25519.NewBlakeSHA256Ed25519()

// NewEncryptor creates a new encryptor
func NewEncryptor() encryptors.Encryptor {
	return createEncryptor()
}
