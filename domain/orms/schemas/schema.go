package schemas

import "steve.care/network/domain/orms/schemas/roots"

type schema struct {
	version  uint
	roots    roots.Roots
	previous Schema
}

func createSchema(
	version uint,
	roots roots.Roots,
) Schema {
	return createSchemaInternally(version, roots, nil)
}

func createSchemaWithPrevious(
	version uint,
	roots roots.Roots,
	previous Schema,
) Schema {
	return createSchemaInternally(version, roots, previous)
}

func createSchemaInternally(
	version uint,
	roots roots.Roots,
	previous Schema,
) Schema {
	out := schema{
		version:  version,
		roots:    roots,
		previous: previous,
	}

	return &out
}

// Version returns the version
func (obj *schema) Version() uint {
	return obj.version
}

// Roots returns the roots
func (obj *schema) Roots() roots.Roots {
	return obj.roots
}

// HasPrevious returns true if there is a previous, false optherwise
func (obj *schema) HasPrevious() bool {
	return obj.previous != nil
}

// Previous returns the previous, if any
func (obj *schema) Previous() Schema {
	return obj.previous
}
