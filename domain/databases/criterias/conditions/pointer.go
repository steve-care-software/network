package conditions

type pointer struct {
	container string
	field     string
}

func createPointer(
	container string,
	field string,
) Pointer {
	out := pointer{
		container: container,
		field:     field,
	}

	return &out
}

// Container returns the container
func (obj *pointer) Container() string {
	return obj.container
}

// Field returns the field
func (obj *pointer) Field() string {
	return obj.field
}
