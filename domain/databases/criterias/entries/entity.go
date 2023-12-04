package entries

import "steve.care/network/domain/databases/criterias/entries/resources"

type entry struct {
	resource resources.Resource
	fields   []string
}

func createEntry(
	resource resources.Resource,
	fields []string,
) Entry {
	out := entry{
		resource: resource,
		fields:   fields,
	}

	return &out
}

// Resource returns the resource
func (obj *entry) Resource() resources.Resource {
	return obj.resource
}

// Fields returns the fields
func (obj *entry) Fields() []string {
	return obj.fields
}
