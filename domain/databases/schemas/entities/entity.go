package entities

import "steve.care/network/domain/databases/schemas/entities/fields"

type entity struct {
	name   string
	head   string
	fields fields.Fields
}

func createEntity(
	name string,
	head string,
	fields fields.Fields,
) Entity {
	out := entity{
		name:   name,
		head:   head,
		fields: fields,
	}

	return &out
}

// Name returns the name
func (obj *entity) Name() string {
	return obj.name
}

// Head returns the head
func (obj *entity) Head() string {
	return obj.head
}

// Fields returns the fields
func (obj *entity) Fields() fields.Fields {
	return obj.fields
}
