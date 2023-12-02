package entities

import "steve.care/network/domain/databases/conditions"

// Entity represents a request entity
type Entity interface {
	Container() string
	Fields() []string
	HasCondition() bool
	Condition() conditions.Condition
}
