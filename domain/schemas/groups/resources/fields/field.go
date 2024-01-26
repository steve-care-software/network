package fields

type field struct {
	name     string
	methods  []string
	kind     uint8
	canBeNil bool
}

func createField(
	name string,
	methods []string,
	kind uint8,
	canBeNil bool,
) Field {
	out := field{
		name:     name,
		methods:  methods,
		kind:     kind,
		canBeNil: canBeNil,
	}

	return &out
}

// Name returns the name
func (obj *field) Name() string {
	return obj.name
}

// Methods returns the methods
func (obj *field) Methods() []string {
	return obj.methods
}

// Kind returns the kind
func (obj *field) Kind() uint8 {
	return obj.kind
}

// CanBeNil returns true if canBeNil, false otherwise
func (obj *field) CanBeNil() bool {
	return obj.canBeNil
}
