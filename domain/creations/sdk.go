package creations

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/resources"
)

// Builder represents the creation builder
type Builder interface {
	Create() Builder
	CreatedBy(receipt receipts.Receipt) Builder
	WithResource(resource resources.Service) Builder
	Now() (Creation, error)
}

// Creation represents the creation
type Creation interface {
	Hash() hash.Hash
	CreatedBy() receipts.Receipt
	Resource() resources.Resource
}
