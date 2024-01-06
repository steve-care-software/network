package permissions

import (
	"steve.care/network/domain/hash"
)

// Builder represents the permission builder
type Builder interface {
	Create() Builder
	WithBlacklist(blacklist []hash.Hash) Builder
	WithWhitelist(whitelist []hash.Hash) Builder
	NOw() (Permission, error)
}

// Permission represents a permission
type Permission interface {
	Hash() hash.Hash
	HasBlacklist() bool
	Blacklist() []hash.Hash
	HasWhitelist() bool
	Whitelist() []hash.Hash
}
