package schemas

import "steve.care/network/domain/schemas/groups"

type schema struct {
	version  uint
	groups   groups.Groups
	previous Schema
}

func createSchema(
	version uint,
	groups groups.Groups,
) Schema {
	return createSchemaInternally(version, groups, nil)
}

func createSchemaWithPrevious(
	version uint,
	groups groups.Groups,
	previous Schema,
) Schema {
	return createSchemaInternally(version, groups, previous)
}

func createSchemaInternally(
	version uint,
	groups groups.Groups,
	previous Schema,
) Schema {
	out := schema{
		version:  version,
		groups:   groups,
		previous: previous,
	}

	return &out
}

// Version returns the version
func (obj *schema) Version() uint {
	return obj.version
}

// Groups returns the groups
func (obj *schema) Groups() groups.Groups {
	return obj.groups
}

// HasPrevious returns true if there is a previous, false optherwise
func (obj *schema) HasPrevious() bool {
	return obj.previous != nil
}

// Previous returns the previous, if any
func (obj *schema) Previous() Schema {
	return obj.previous
}
