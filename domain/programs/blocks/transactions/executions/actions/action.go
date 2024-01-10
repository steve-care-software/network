package actions

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
)

type action struct {
	hash   hash.Hash
	create resources.Resource
	del    hash.Hash
}

func createActionWithCreate(
	hash hash.Hash,
	create resources.Resource,
) Action {
	return createActionInternally(hash, create, nil)
}

func createActionWithDelete(
	hash hash.Hash,
	del hash.Hash,
) Action {
	return createActionInternally(hash, nil, del)
}

func createActionInternally(
	hash hash.Hash,
	create resources.Resource,
	del hash.Hash,
) Action {
	out := action{
		hash:   hash,
		create: create,
		del:    del,
	}

	return &out
}

// Hash returns the hash
func (obj *action) Hash() hash.Hash {
	return obj.hash
}

// IsCreate returns true if create, false otherwise
func (obj *action) IsCreate() bool {
	return obj.create != nil
}

// Create returns the create, if any
func (obj *action) Create() resources.Resource {
	return obj.create
}

// IsDelete returns true if delete, false otherwise
func (obj *action) IsDelete() bool {
	return obj.del != nil
}

// Delete returns the delete, if any
func (obj *action) Delete() hash.Hash {
	return obj.del
}
