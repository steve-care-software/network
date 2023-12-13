package entries

import "steve.care/network/domain/databases/criterias"

type entry struct {
	criteria criterias.Criteria
	fields   []string
}

func createEntry(
	criteria criterias.Criteria,
	fields []string,
) Entry {
	out := entry{
		criteria: criteria,
		fields:   fields,
	}

	return &out
}

// Criteria returns the criteria
func (obj *entry) Criteria() criterias.Criteria {
	return obj.criteria
}

// Fields returns the fields
func (obj *entry) Fields() []string {
	return obj.fields
}
