package schemas

import "steve.care/network/domain/schemas/resources"

type schema struct {
	version   uint
	resources resources.Resources
	previous  Schema
}

func createSchema(
	version uint,
	resources resources.Resources,
) Schema {
	return createSchemaInternally(version, resources, nil)
}

func createSchemaWithPrevious(
	version uint,
	resources resources.Resources,
	previous Schema,
) Schema {
	return createSchemaInternally(version, resources, previous)
}

func createSchemaInternally(
	version uint,
	resources resources.Resources,
	previous Schema,
) Schema {
	out := schema{
		version:   version,
		resources: resources,
		previous:  previous,
	}

	return &out
}

// Version returns the version
func (obj *schema) Version() uint {
	return obj.version
}

// Resources returns the resources
func (obj *schema) Resources() resources.Resources {
	return obj.resources
}

// HasPrevious returns true if there is a previous, false optherwise
func (obj *schema) HasPrevious() bool {
	return obj.previous != nil
}

// Previous returns the previous, if any
func (obj *schema) Previous() Schema {
	return obj.previous
}
