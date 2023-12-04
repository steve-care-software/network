package entities

type entities struct {
	list []Entity
}

func createEntities(
	list []Entity,
) Entities {
	out := entities{
		list: list,
	}

	return &out
}

// List returns the list of entities
func (obj *entities) List() []Entity {
	return obj.list
}
