package kinds

// Kind represents the kind of field
type Kind interface {
	IsNil() bool
	IsInteger() bool
	IsReal() bool
	IsText() bool
	IsBlob() bool
}
