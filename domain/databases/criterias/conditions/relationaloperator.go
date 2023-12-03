package conditions

type relationalOperator struct {
	isAnd bool
	isOr  bool
}

func createRelationalOperatorWithAnd() RelationalOperator {
	return createRelationalOperatorInternally(true, false)
}

func createRelationalOperatorWithOr() RelationalOperator {
	return createRelationalOperatorInternally(false, true)
}

func createRelationalOperatorInternally(
	isAnd bool,
	isOr bool,
) RelationalOperator {
	out := relationalOperator{
		isAnd: isAnd,
		isOr:  isOr,
	}

	return &out
}

// IsAnd returns true if and, false otherwise
func (obj *relationalOperator) IsAnd() bool {
	return obj.isAnd
}

// IsOr returns true if or, false otherwise
func (obj *relationalOperator) IsOr() bool {
	return obj.isAnd
}
