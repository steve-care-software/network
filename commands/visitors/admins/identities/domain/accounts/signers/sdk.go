package signers

import "steve.care/network/libraries/hash"

// Factory represents a signer factory
type Factory interface {
	Create() Signer
}

// Signer represents a signer
type Signer interface {
	PublicKey() PublicKey
	Sign(msg []byte) (Signature, error)
	Vote(msg []byte) (Vote, error)
	Bytes() []byte
}

// PublicKey represents a public key
type PublicKey interface {
	Equals(pubKey PublicKey) bool
	Bytes() []byte
}

// Vote represents a vote
type Vote interface {
	Hash() hash.Hash
	Ring() []PublicKey
	Verify(msg []byte) bool
	Bytes() []byte
}

// Signature represents a signature
type Signature interface {
	PublicKey(msg []byte) PublicKey
	Verify() bool
	Bytes() []byte
}
