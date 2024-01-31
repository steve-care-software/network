package methods

type methods struct {
	initialize string
	trigger    string
	element    string
}

func createMethods(
	initialize string,
	trigger string,
) Methods {
	return createMethodsInternally(initialize, trigger, "")
}

func createMethodsWithElement(
	initialize string,
	trigger string,
	element string,
) Methods {
	return createMethodsInternally(initialize, trigger, element)
}

func createMethodsInternally(
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

// HasElement returns true if there is an element method
func (obj *methods) HasElement() bool {
	return obj.element != ""
}

// Element returns the element method
func (obj *methods) Element() string {
	return obj.element
}
