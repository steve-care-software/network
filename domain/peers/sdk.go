package peers

import (
	"net/url"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs"
)

// Builder represents the peers builder
type Builder interface {
	Create() Builder
	WithProgram(program programs.Program) Builder
	WithPeers(peers []*url.URL) Builder
	Now() (Peers, error)
}

// Peers represents peers
type Peers interface {
	Program() programs.Program
	Peers() []*url.URL
}

// Repository represents the peers repository
type Repository interface {
	Retrieve(program hash.Hash) (Peers, error)
}
