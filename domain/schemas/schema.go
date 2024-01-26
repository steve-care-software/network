package schemas

import "steve.care/network/domain/schemas/groups"

type schema struct {
	version  uint
	group    groups.Group
	previous Schema
}

func createSchema(
	version uint,
	group groups.Group,
) Schema {
	return createSchemaInternally(version, group, nil)
}

func createSchemaWithPrevious(
	version uint,
	group groups.Group,
	previous Schema,
) Schema {
	return createSchemaInternally(version, group, previous)
}

func createSchemaInternally(
	version uint,
	group groups.Group,
	previous Schema,
) Schema {
	out := schema{
		version:  version,
		group:    group,
		previous: previous,
	}

	return &out
}

// Version returns the version
func (obj *schema) Version() uint {
	return obj.version
}

// Group returns the group
func (obj *schema) Group() groups.Group {
	return obj.group
}

// HasPrevious returns true if there is a previous, false optherwise
func (obj *schema) HasPrevious() bool {
	return obj.previous != nil
}

// Previous returns the previous, if any
func (obj *schema) Previous() Schema {
	return obj.previous
}
