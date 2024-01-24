package resources

type connection struct {
	from Pointer
	to   Pointer
}

func createConnection(
	from Pointer,
	to Pointer,
) Connection {
	out := connection{
		from: from,
		to:   to,
	}

	return &out
}

// From returns the from pointer
func (obj *connection) From() Pointer {
	return obj.from
}

// To returns the to pointer
func (obj *connection) To() Pointer {
	return obj.to
}
