package methods

type methods struct {
	initialize string
	trigger    string
	builder    string
}

func createMethods(
	initialize string,
	trigger string,
	builder string,
) Methods {
	out := methods{
		initialize: initialize,
		trigger:    trigger,
		builder:    builder,
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

// Builder returns the builder
func (obj *methods) Builder() string {
	return obj.builder
}
