package schemas

import "steve.care/network/domain/schemas/roots"

type schema struct {
	version  uint
	root     roots.Root
	previous Schema
}

func createSchema(
	version uint,
	root roots.Root,
) Schema {
	return createSchemaInternally(version, root, nil)
}

func createSchemaWithPrevious(
	version uint,
	root roots.Root,
	previous Schema,
) Schema {
	return createSchemaInternally(version, root, previous)
}

func createSchemaInternally(
	version uint,
	root roots.Root,
	previous Schema,
) Schema {
	out := schema{
		version:  version,
		root:     root,
		previous: previous,
	}

	return &out
}

// Version returns the version
func (obj *schema) Version() uint {
	return obj.version
}

// Root returns the root
func (obj *schema) Root() roots.Root {
	return obj.root
}

// HasPrevious returns true if there is a previous, false optherwise
func (obj *schema) HasPrevious() bool {
	return obj.previous != nil
}

// Previous returns the previous, if any
func (obj *schema) Previous() Schema {
	return obj.previous
}
