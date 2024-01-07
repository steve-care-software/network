package results

import "steve.care/network/domain/programs/logics/libraries/layers"

// NewResultWithFailureForTests creates a new result with failure for tests
func NewResultWithFailureForTests(failure Failure) Result {
	ins, err := NewBuilder().Create().WithFailure(failure).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewResultWithSuccessForTests creates a new result with success for tests
func NewResultWithSuccessForTests(success Success) Result {
	ins, err := NewBuilder().Create().WithSuccess(success).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSuccessForTests creates a new success for tests
func NewSuccessForTests(bytes []byte, kind layers.Kind) Success {
	ins, err := NewSuccessBuilder().Create().WithBytes(bytes).WithKind(kind).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewFailureForTests creates a new failure for tests
func NewFailureForTests(code uint, isRaisedInLayer bool) Failure {
	builder := NewFailureBuilder().Create().WithCode(code)
	if isRaisedInLayer {
		builder.IsRaisedInLayer()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
