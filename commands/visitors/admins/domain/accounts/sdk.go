package accounts

import (
	"steve.care/network/libraries/credentials"
	"steve.care/network/libraries/hash"
)

// Builder represents an account builder
type Builder interface {
	Create() Builder
	WithUsername(username string) Builder
	WithIdentities(identities Identities) Builder
	WithCreator(creator Account) Builder
	Now() (Account, error)
}

// Account represents an admin account
type Account interface {
	Hash() hash.Hash
	Username() string
	HasIdentities() bool
	Identities() Identities
	HasCreator() bool
	Creator() Account
}

// IdentitiesBuilder represents an identities builder
type IdentitiesBuilder interface {
	Create() IdentitiesBuilder
	WithList(list []Identity) IdentitiesBuilder
	Now() (Identities, error)
}

// Identities represnets identities
type Identities interface {
	List() []Identity
}

// IdentityBuilder represents an identity builder
type IdentityBuilder interface {
	Create() IdentityBuilder
	WithUsername(username string) IdentityBuilder
	WithDescription(description string) IdentityBuilder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	Hash() hash.Hash
	Username() string
	HasDescription() bool
	Description() string
}

// Repository represents an account repository
type Repository interface {
	List() ([]string, error)
	Exists(username string) (bool, error)
	Retrieve(credentials credentials.Credentials) (Account, error)
}
