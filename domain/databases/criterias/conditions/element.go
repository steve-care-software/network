package conditions

type element struct {
	condition Condition
	resource  Resource
}

func createElementWithCondition(
	condition Condition,
) Element {
	return createElementInternally(condition, nil)
}

func createElementWithResource(
	resource Resource,
) Element {
	return createElementInternally(nil, resource)
}

func createElementInternally(
	condition Condition,
	resource Resource,
) Element {
	out := element{
		condition: condition,
		resource:  resource,
	}

	return &out
}

// IsCondition returns true if there is a condition, false otherwise
func (obj *element) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *element) Condition() Condition {
	return obj.condition
}

// IsResource returns true if there is a resource, false otherwise
func (obj *element) IsResource() bool {
	return obj.resource != nil
}

// Resource returns the resource, if any
func (obj *element) Resource() Resource {
	return obj.resource
}
