package methods

type methods struct {
	initialize string
	trigger    string
	element    string
}

func createMethods(
	initialize string,
	trigger string,
	element string,
) Methods {
	out := methods{
		initialize: initialize,
		trigger:    trigger,
		element:    element,
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

// Element returns the element method
func (obj *methods) Element() string {
	return obj.element
}
