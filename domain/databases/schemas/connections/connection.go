package connections

import "steve.care/network/domain/databases/schemas/connections/pointers"

type connection struct {
	from pointers.Pointer
	to   pointers.Pointer
}

func createConnection(
	from pointers.Pointer,
	to pointers.Pointer,
) Connection {
	out := connection{
		from: from,
		to:   to,
	}

	return &out
}

// From returns the from pointer
func (obj *connection) From() pointers.Pointer {
	return obj.from
}

// To returns the to pointer
func (obj *connection) To() pointers.Pointer {
	return obj.to
}
