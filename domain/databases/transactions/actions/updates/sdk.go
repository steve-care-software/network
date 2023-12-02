package updates

import (
	"steve.care/network/domain/databases/conditions"
	"steve.care/network/domain/databases/values"
)

// Update represents an update
type Update interface {
	Container() string
	Values() map[string]values.Value
	Condition() conditions.Condition
}
