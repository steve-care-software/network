package conditions

type operator struct {
	isEqual    bool
	relational RelationalOperator
	integer    IntegerOperator
}

func createOperatorWithEqual() Operator {
	return createOperatorInternally(true, nil, nil)
}

func createOperatorWithRelational(
	relational RelationalOperator,
) Operator {
	return createOperatorInternally(false, relational, nil)
}

func createOperatorWithInteger(
	integer IntegerOperator,
) Operator {
	return createOperatorInternally(false, nil, integer)
}

func createOperatorInternally(
	isEqual bool,
	relational RelationalOperator,
	integer IntegerOperator,
) Operator {
	out := operator{
		isEqual:    isEqual,
		relational: relational,
		integer:    integer,
	}

	return &out
}

// IsEqual returns true if equal, false otherwise
func (obj *operator) IsEqual() bool {
	return obj.isEqual
}

// IsRelational returns true if relational, false otherwise
func (obj *operator) IsRelational() bool {
	return obj.relational != nil
}

// Relational returns the relational operaotr, if any
func (obj *operator) Relational() RelationalOperator {
	return obj.relational
}

// IsInteger returns true if integer, false otherwise
func (obj *operator) IsInteger() bool {
	return obj.integer != nil
}

// Integer returns the integer operaotr, if any
func (obj *operator) Integer() IntegerOperator {
	return obj.integer
}
