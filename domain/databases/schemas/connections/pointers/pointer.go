package pointers

type pointer struct {
	entity string
	field  string
}

func createPointer(
	entity string,
	field string,
) Pointer {
	out := pointer{
		entity: entity,
		field:  field,
	}

	return &out
}

// Entity returns the entity
func (obj *pointer) Entity() string {
	return obj.entity
}

// Field returns the field
func (obj *pointer) Field() string {
	return obj.field
}
