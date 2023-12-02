package kinds

// Kind represents a field kind
type Kind interface {
	IsNil() bool
	IsInteger() bool
	IsReal() bool
	IsText() bool
	IsBytes() bool
}
