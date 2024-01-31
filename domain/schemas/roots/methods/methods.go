package methods

type methods struct {
	initialize string
	trigger    string
}

func createMethods(
	initialize string,
	trigger string,
) Methods {
	out := methods{
		initialize: initialize,
		trigger:    trigger,
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
