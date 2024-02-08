package orms

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/orms/skeletons"
)

// Instance represents an instance
type Instance interface {
	Hash() hash.Hash
}

// RepositoryBuilder represents the repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithSkeleton(skeleton skeletons.Skeleton) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents an instance repository
type Repository interface {
	// AmountByQuery returns the amount of instance by criteria
	AmountByQuery(query hash.Hash) (uint, error)

	// ListByQuery lists insatnce hashes by criteria
	ListByQuery(query hash.Hash) ([]hash.Hash, error)

	// RetrieveByQuery retrieves an instance by criteria
	RetrieveByQuery(query hash.Hash) (Instance, error)

	// RetrieveByHash retrieves an instance by hash
	RetrieveByHash(path []string, hash hash.Hash) (Instance, error)
}

// ServiceBuilder represents the service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithSkeleton(skeleton skeletons.Skeleton) ServiceBuilder
	Now() (Service, error)
}

// Service represents a an instance service
type Service interface {
	// Init initializes the service
	Init() error

	// Insert inserts an instance
	Insert(ins Instance, path []string) error

	// Delete deletes an instance
	Delete(path []string, hash hash.Hash) error
}
