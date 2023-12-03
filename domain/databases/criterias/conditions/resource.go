package conditions

type resource struct {
	field Pointer
	value interface{}
}

func createResourceWithField(
	field Pointer,
) Resource {
	return createResourceInternally(field, nil)
}

func createResourceWithValue(
	value interface{},
) Resource {
	return createResourceInternally(nil, value)
}

func createResourceInternally(
	field Pointer,
	value interface{},
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
func (obj *resource) Value() interface{} {
	return obj.value
}
