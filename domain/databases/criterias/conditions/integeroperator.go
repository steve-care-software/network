package conditions

type integerOperator struct {
	isSmallerThan bool
	isBiggerThan  bool
	isEqual       bool
}

func createIntegerOperator() IntegerOperator {
	return createIntegerOperatorInternally(
		false,
		false,
		false,
	)
}

func createIntegerOperatorWithSmallerThan() IntegerOperator {
	return createIntegerOperatorInternally(
		true,
		false,
		false,
	)
}

func createIntegerOperatorWithSmallerThanAndEqual() IntegerOperator {
	return createIntegerOperatorInternally(
		true,
		false,
		true,
	)
}

func createIntegerOperatorWithBiggerThan() IntegerOperator {
	return createIntegerOperatorInternally(
		false,
		true,
		false,
	)
}

func createIntegerOperatorWithBiggerThanAndEqual() IntegerOperator {
	return createIntegerOperatorInternally(
		false,
		true,
		true,
	)
}

func createIntegerOperatorWithEqual() IntegerOperator {
	return createIntegerOperatorInternally(
		false,
		false,
		true,
	)
}

func createIntegerOperatorInternally(isSmallerThan bool, isBiggerThan bool, isEqual bool) IntegerOperator {
	out := integerOperator{
		isSmallerThan: isSmallerThan,
		isBiggerThan:  isBiggerThan,
		isEqual:       isEqual,
	}

	return &out
}

// IsSmallerThan returns true if smaller than, false otherwise
func (obj *integerOperator) IsSmallerThan() bool {
	return obj.isSmallerThan
}

// IsBiggerThan returns true if bigger than, false otherwise
func (obj *integerOperator) IsBiggerThan() bool {
	return obj.isBiggerThan
}

// IsEqual returns true if equal, false otherwise
func (obj *integerOperator) IsEqual() bool {
	return obj.isEqual
}
