package kinds

type kind struct {
	isNil   bool
	pInt    *int
	pFloat  *float64
	pString *string
	pBytes  []byte
}

func createKindWithNil() Kind {
	return createKindInternally(true, nil, nil, nil, nil)
}

func createKindWithInteger(
	pInt *int,
) Kind {
	return createKindInternally(false, pInt, nil, nil, nil)
}

func createKindWithFloat(
	pFloat *float64,
) Kind {
	return createKindInternally(false, nil, pFloat, nil, nil)
}

func createKindWithString(
	pString *string,
) Kind {
	return createKindInternally(false, nil, nil, pString, nil)
}

func createKindWithBytes(
	pBytes []byte,
) Kind {
	return createKindInternally(false, nil, nil, nil, pBytes)
}

func createKindInternally(
	isNil bool,
	pInt *int,
	pFloat *float64,
	pString *string,
	pBytes []byte,
) Kind {
	out := kind{
		isNil:   isNil,
		pInt:    pInt,
		pFloat:  pFloat,
		pString: pString,
		pBytes:  pBytes,
	}

	return &out
}

// IsNil returns true if nil, false otherwise
func (obj *kind) IsNil() bool {
	return obj.isNil
}

// IsInteger returns true if integer, false otherwise
func (obj *kind) IsInteger() bool {
	return obj.pInt != nil
}

// Integer returns the integer, if any
func (obj *kind) Integer() *int {
	return obj.pInt
}

// IsFloat returns true if float, false otherwise
func (obj *kind) IsFloat() bool {
	return obj.pFloat != nil
}

// Float returns the float, if any
func (obj *kind) Float() *float64 {
	return obj.pFloat
}

// IsString returns true if string, false otherwise
func (obj *kind) IsString() bool {
	return obj.pString != nil
}

// String returns the string, if any
func (obj *kind) String() *string {
	return obj.pString
}

// IsBytes returns true if bytes, false otherwise
func (obj *kind) IsBytes() bool {
	return obj.pBytes != nil
}

// Bytes returns the bytes, if any
func (obj *kind) Bytes() []byte {
	return obj.pBytes
}
