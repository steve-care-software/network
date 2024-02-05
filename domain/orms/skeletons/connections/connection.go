package connections

type connection struct {
	name string
	from []string
	to   []string
}

func createConnection(
	name string,
	from []string,
	to []string,
) Connection {
	out := connection{
		name: name,
		from: from,
		to:   to,
	}

	return &out
}

// Name returns the name
func (obj *connection) Name() string {
	return obj.name
}

// From returns the from
func (obj *connection) From() []string {
	return obj.from
}

// To returns the to
func (obj *connection) To() []string {
	return obj.to
}
