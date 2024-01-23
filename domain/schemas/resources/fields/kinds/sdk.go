package kinds

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the kind builder
type Builder interface {
	Create() Builder
	WithInteger(intValue int) Builder
	WithFloat(floatValue float64) Builder
	WithString(stringValue string) Builder
	WithBytes(bytes []byte) Builder
	IsNil() Builder
	Now() (Kind, error)
}

// Kind represents a field kind
type Kind interface {
	IsNil() bool
	IsInteger() bool
	Integer() *int
	IsFloat() bool
	Float() *float64
	IsString() bool
	String() *string
	IsBytes() bool
	Bytes() []byte
}
