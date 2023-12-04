package kinds

type kind struct {
	isNil     bool
	isInteger bool
	isReal    bool
	isText    bool
	isBlob    bool
}

func createKindWithNil() Kind {
	return createKindInternally(
		true,
		false,
		false,
		false,
		false,
	)
}

func createKindWithInteger() Kind {
	return createKindInternally(
		false,
		true,
		false,
		false,
		false,
	)
}

func createKindWithReal() Kind {
	return createKindInternally(
		false,
		false,
		true,
		false,
		false,
	)
}

func createKindWithText() Kind {
	return createKindInternally(
		false,
		false,
		false,
		true,
		false,
	)
}

func createKindWithBlob() Kind {
	return createKindInternally(
		false,
		false,
		false,
		false,
		true,
	)
}

func createKindInternally(
	isNil bool,
	isInteger bool,
	isReal bool,
	isText bool,
	isBlob bool,
) Kind {
	out := kind{
		isNil:     false,
		isInteger: false,
		isReal:    false,
		isText:    false,
		isBlob:    false,
	}

	return &out
}

// IsNil returns true if nil, false otherwise
func (obj *kind) IsNil() bool {
	return obj.isNil
}

// IsInteger returns true if integer, false otherwise
func (obj *kind) IsInteger() bool {
	return obj.isInteger
}

// IsReal returns true if real, false otherwise
func (obj *kind) IsReal() bool {
	return obj.isReal
}

// IsText returns true if text, false otherwise
func (obj *kind) IsText() bool {
	return obj.isText
}

// IsBlob returns true if blob, false otherwise
func (obj *kind) IsBlob() bool {
	return obj.isBlob
}
