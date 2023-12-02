package actions

import (
	"steve.care/network/domain/databases/transactions/actions/deletes"
	"steve.care/network/domain/databases/transactions/actions/inserts"
	"steve.care/network/domain/databases/transactions/actions/updates"
)

// Action represents an action
type Action interface {
	IsDelete() bool
	Delete() deletes.Delete
	IsUpdate() bool
	Update() updates.Update
	IsInsert() bool
	Insert() inserts.Insert
}
