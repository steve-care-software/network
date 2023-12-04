package queries

import (
	"steve.care/network/domain/databases/criterias/entries/resources"
	"steve.care/network/domain/hash"
)

// Query represents the query resource
type Query interface {
	Hash() hash.Hash
	IsCriteria() bool
	Criteria() resources.Resource
}
