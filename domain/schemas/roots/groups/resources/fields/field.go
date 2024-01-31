package fields

import (
	"steve.care/network/domain/schemas/roots/groups/resources/fields/methods"
	"steve.care/network/domain/schemas/roots/groups/resources/fields/types"
)

type field struct {
	name     string
	methods  methods.Methods
	typ      types.Type
	canBeNil bool
}

func createField(
	name string,
	methods methods.Methods,
	typ types.Type,
	canBeNil bool,
) Field {
	out := field{
		name:     name,
		methods:  methods,
		typ:      typ,
		canBeNil: canBeNil,
	}

	return &out
}

// Name returns the name
func (obj *field) Name() string {
	return obj.name
}

// Methods returns the methods
func (obj *field) Methods() methods.Methods {
	return obj.methods
}

// Type returns the type
func (obj *field) Type() types.Type {
	return obj.typ
}

// CanBeNil returns true if canBeNil, false otherwise
func (obj *field) CanBeNil() bool {
	return obj.canBeNil
}
