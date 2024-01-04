package conditions

// NewIntegerOperatorWithIsEqualForTests creates a new integer operator with IsEqual for tests
func NewIntegerOperatorWithIsEqualForTests() IntegerOperator {
	ins, err := NewIntegerOperatorBuilder().Create().IsEqual().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIntegerOperatorWithIsBiggerThanForTests creates a new integer operator with IsBiggerThan for tests
func NewIntegerOperatorWithIsBiggerThanForTests() IntegerOperator {
	ins, err := NewIntegerOperatorBuilder().Create().IsBiggerThan().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIntegerOperatorWithIsSmallerThanForTests creates a new integer operator with IsSmallerThan for tests
func NewIntegerOperatorWithIsSmallerThanForTests() IntegerOperator {
	ins, err := NewIntegerOperatorBuilder().Create().IsSmallerThan().Now()
	if err != nil {
		panic(err)
	}

	return ins
}
