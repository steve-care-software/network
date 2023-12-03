package values

type value struct {
	isNil      bool
	pIntValue  *int
	pRealValue *float64
	textValue  string
	bytes      []byte
}

func createValueWithNil() Value {
	return createValueInternally(
		true,
		nil,
		nil,
		"",
		nil,
	)
}

func createValueWithInteger(
	pIntValue *int,
) Value {
	return createValueInternally(
		false,
		pIntValue,
		nil,
		"",
		nil,
	)
}

func createValueWithReal(
	pRealValue *float64,
) Value {
	return createValueInternally(
		false,
		nil,
		pRealValue,
		"",
		nil,
	)
}

func createValueWithText(
	textValue string,
) Value {
	return createValueInternally(
		false,
		nil,
		nil,
		textValue,
		nil,
	)
}

func createValueWithBytes(
	bytes []byte,
) Value {
	return createValueInternally(
		false,
		nil,
		nil,
		"",
		bytes,
	)
}

func createValueInternally(
	isNil bool,
	pIntValue *int,
	pRealValue *float64,
	textValue string,
	bytes []byte,
) Value {
	out := value{
		isNil:      isNil,
		pIntValue:  pIntValue,
		pRealValue: pRealValue,
		textValue:  textValue,
		bytes:      bytes,
	}

	return &out
}

// IsNil returns true if nil, false otherwise
func (obj *value) IsNil() bool {
	return obj.isNil
}

// IsInteger returns true if integer, false otherwise
func (obj *value) IsInteger() bool {
	return obj.pIntValue != nil
}

// Integer returns the integer, if any
func (obj *value) Integer() *int {
	return obj.pIntValue
}

// IsReal returns true if real, false otherwise
func (obj *value) IsReal() bool {
	return obj.pRealValue != nil
}

// Real returns the real, if any
func (obj *value) Real() *float64 {
	return obj.pRealValue
}

// IsText returns true if text, false otherwise
func (obj *value) IsText() bool {
	return obj.textValue != ""
}

// Text returns the text, if any
func (obj *value) Text() string {
	return obj.textValue
}

// IsBytes returns true if bytes, false otherwise
func (obj *value) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes, if any
func (obj *value) Bytes() []byte {
	return obj.bytes
}
