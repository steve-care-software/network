package resources

type pointer struct {
	resource Resource
	field    string
}

func createPointer(
	resource Resource,
	field string,
) Pointer {
	out := pointer{
		resource: resource,
		field:    field,
	}

	return &out
}

// Resource returns the resource
func (obj *pointer) Resource() Resource {
	return obj.resource
}

// Field returns the field
func (obj *pointer) Field() string {
	return obj.field
}
