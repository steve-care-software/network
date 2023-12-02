package values

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
