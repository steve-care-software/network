package fields

type fields struct {
	list []Field
}

func createFields(
	list []Field,
) Fields {
	out := fields{
		list: list,
	}

	return &out
}

// List returns the list of fields
func (obj *fields) List() []Field {
	return obj.list
}
