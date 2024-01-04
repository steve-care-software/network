package roots

import "steve.care/network/domain/hash"

// Root represents a blockchain root
type Root interface {
	Hash() hash.Hash
	HasOwners() bool
	Owners() []hash.Hash
}
