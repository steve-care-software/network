package lists

import "steve.care/network/domain/hash"

// List represents a list
type List interface {
	Container() string
	Hashes() []hash.Hash
}
