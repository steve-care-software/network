package conditions

import "steve.care/network/domain/databases/criterias/values"

type resource struct {
	field Pointer
	value values.Value
}

func createResourceWithField(
	field Pointer,
) Resource {
	return createResourceInternally(field, nil)
}

func createResourceWithValue(
	value values.Value,
) Resource {
	return createResourceInternally(nil, value)
}

func createResourceInternally(
	field Pointer,
	value values.Value,
) Resource {
	out := resource{
		field: field,
		value: value,
	}

	return &out
}

// IsField returns true if there is a field, false otherwise
func (obj *resource) IsField() bool {
	return obj.field != nil
}

// Field returns the field, if any
func (obj *resource) Field() Pointer {
	return obj.field
}

// IsValue returns true if there is a value, false otherwise
func (obj *resource) IsValue() bool {
	return obj.value != nil
}

// Value returns the value, if any
func (obj *resource) Value() values.Value {
	return obj.value
}
