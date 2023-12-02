package requests

import (
	"steve.care/network/domain/databases/queries/requests/entities"
	"steve.care/network/domain/databases/queries/requests/lists"
)

// Request represents a query request
type Request interface {
	IsEntity() bool
	Entity() entities.Entity
	IsList() bool
	List() lists.List
}
