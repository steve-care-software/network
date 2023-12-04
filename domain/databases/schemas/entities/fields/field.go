package fields

import "steve.care/network/domain/databases/schemas/entities/fields/kinds"

type field struct {
	name     string
	kind     kinds.Kind
	isUnique bool
}

func createField(
	name string,
	kind kinds.Kind,
	isUnique bool,
) Field {
	out := field{
		name:     name,
		kind:     kind,
		isUnique: isUnique,
	}

	return &out
}

// Name returns the name
func (obj *field) Name() string {
	return obj.name
}

// Kind returns the kind
func (obj *field) Kind() kinds.Kind {
	return obj.kind
}

// IsUnique returns true if unique, false otherwise
func (obj *field) IsUnique() bool {
	return obj.isUnique
}
