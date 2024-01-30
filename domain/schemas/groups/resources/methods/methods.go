package methods

import field_methods "steve.care/network/domain/schemas/groups/resources/fields/methods"

type methods struct {
	initialize string
	trigger    string
	field      field_methods.Methods
}

func createMethods(
	initialize string,
	trigger string,
	field field_methods.Methods,
) Methods {
	out := methods{
		initialize: initialize,
		trigger:    trigger,
		field:      field,
	}

	return &out
}

// Initialize returns the initialize
func (obj *methods) Initialize() string {
	return obj.initialize
}

// Trigger returns the trigger
func (obj *methods) Trigger() string {
	return obj.trigger
}

// Field returns the field
func (obj *methods) Field() field_methods.Methods {
	return obj.field
}
