package resources

type connection struct {
	field     string
	reference Pointer
}

func createConnection(
	field string,
	reference Pointer,
) Connection {
	out := connection{
		field:     field,
		reference: reference,
	}

	return &out
}

// Field returns the field
func (obj *connection) Field() string {
	return obj.field
}

// Reference returns the reference
func (obj *connection) Reference() Pointer {
	return obj.reference
}
