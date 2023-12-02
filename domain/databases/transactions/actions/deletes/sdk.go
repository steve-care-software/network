package deletes

import (
	"steve.care/network/domain/hash"
)

// Delete represents a delete action
type Delete interface {
	Hash() hash.Hash
	Container() string
}
