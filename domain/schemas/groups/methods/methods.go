package methods

import (
	resource_methods "steve.care/network/domain/schemas/groups/resources/methods"
)

type methods struct {
	builder resource_methods.Methods
}

func createMethods(
	builder resource_methods.Methods,
) Methods {
	out := methods{
		builder: builder,
	}

	return &out
}

// Builder returns the builder
func (obj *methods) Builder() resource_methods.Methods {
	return obj.builder
}
