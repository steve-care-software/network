package fields

type field struct {
	name     string
	kind     uint8
	canBeNil bool
}

func createField(
	name string,
	kind uint8,
	canBeNil bool,
) Field {
	out := field{
		name:     name,
		kind:     kind,
		canBeNil: canBeNil,
	}

	return &out
}

// Name returns the name
func (obj *field) Name() string {
	return obj.name
}

// Kind returns the kind
func (obj *field) Kind() uint8 {
	return obj.kind
}

// CanBeNil returns true if canBeNil, false otherwise
func (obj *field) CanBeNil() bool {
	return obj.canBeNil
}
