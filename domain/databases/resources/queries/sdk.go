package queries

import (
	"steve.care/network/domain/databases/criterias/conditions"
	"steve.care/network/domain/databases/criterias/entities/resources"
	"steve.care/network/domain/hash"
)

// Query represents the query resource
type Query interface {
	Hash() hash.Hash
	IsResource() bool
	Resource() resources.Resource
	IsCondition() bool
	Condition() conditions.Condition
}
