package values

// Builder represents the value builder
type Builder interface {
	Create() Builder
	WithInteger(intValue int) Builder
	WithReal(realBalue float64) Builder
	WithText(text string) Builder
	WithBytes(bytes []byte) Builder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsNil() bool
	IsInteger() bool
	Integer() int
	IsReal() bool
	Real() float64
	IsText() bool
	Text() string
	IsBytes() bool
	Bytes() []byte
}
