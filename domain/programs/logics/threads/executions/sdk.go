package executions

import (
	"time"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/threads"
	"steve.care/network/domain/receipts"
)

// Builder represents an execution
type Builder interface {
	Create() Builder
	WithThread(thread threads.Thread) Builder
	WithReceipt(receipt receipts.Receipt) Builder
	BeginsOn(beginsOn time.Time) Builder
	EndsOn(endsOn time.Time) Builder
	Now() (Execution, error)
}

// Execution represents a thread execution
type Execution interface {
	Hash() hash.Hash
	Thread() threads.Thread
	Receipt() receipts.Receipt
	BeginsOn() time.Time
	EndsOn() time.Time
}

// Repository represents an execution repository
type Repository interface {
	ListByThread(thread threads.Thread) ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Execution, error)
}

// Service represents an execution service
type Service interface {
	Insert(execution Execution) error
	Delete(hash hash.Hash) error
}
