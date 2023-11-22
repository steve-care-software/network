package results

import "hash"

const (
	// InputNotFoundError represents an input not found error
	InputNotFoundError (uint) = iota

	// InputNotBytesError represents an input not bytes error
	InputNotBytesError

	// OutputNotFoundError represents an output not found error
	OutputNotFoundError

	// OutputNotBytesError represents an output not bytes error
	OutputNotBytesError
)

// Builder represents the result builder
type Builder interface {
	Create() Builder
	WithSuccess(success Success) Builder
	WithFailure(failure Failure) Builder
	Now() (Result, error)
}

// Result represents result
type Result interface {
	Hash() hash.Hash
	IsSuccess() bool
	Success() []byte
	IsFailure() bool
	Failure() Failure
}

// SuccessBuilder represents the success builder
type SuccessBuilder interface {
	Create() SuccessBuilder
	WithBytes(bytes []byte) SuccessBuilder
	WithAction(action Action) SuccessBuilder
	Now() (Success, error)
}

// Success represents success result
type Success interface {
	Hash() hash.Hash
	Bytes() []byte
	Action() Action
}

// FailureBuilder represents the failure builder
type FailureBuilder interface {
	Create() FailureBuilder
	WithIndex(index uint) FailureBuilder
	WithCode(code uint) FailureBuilder
	IsRaisedInLayer() FailureBuilder
	Now() (Failure, error)
}

// Failure represents failure result
type Failure interface {
	Hash() hash.Hash
	Code() uint
	IsRaisedInLayer() bool
	HasIndex() bool
	Index() uint
}

// ActionBuilder represents the action builder
type ActionBuilder interface {
	Create() ActionBuilder
	IsPrompt() ActionBuilder
	IsContinue() ActionBuilder
	Now() (Action, error)
}

// Action represents the action
type Action interface {
	Hash() hash.Hash
	IsPrompt() bool
	ISContinue() bool
}
