package executions

import (
	"time"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/threads"
	"steve.care/network/domain/receipts"
)

type execution struct {
	hash     hash.Hash
	thread   threads.Thread
	receipt  receipts.Receipt
	beginsOn time.Time
	endsOn   time.Time
}

func createExecution(
	hash hash.Hash,
	thread threads.Thread,
	receipt receipts.Receipt,
	beginsOn time.Time,
	endsOn time.Time,
) Execution {
	out := execution{
		hash:     hash,
		thread:   thread,
		receipt:  receipt,
		beginsOn: beginsOn,
		endsOn:   endsOn,
	}

	return &out
}

// Hash returns the hash
func (obj *execution) Hash() hash.Hash {
	return obj.hash
}

// Thread returns the thread
func (obj *execution) Thread() threads.Thread {
	return obj.thread
}

// Receipt returns the receipt
func (obj *execution) Receipt() receipts.Receipt {
	return obj.receipt
}

// BeginsOn returns the beginsOn
func (obj *execution) BeginsOn() time.Time {
	return obj.beginsOn
}

// EndsOn returns the endsOn
func (obj *execution) EndsOn() time.Time {
	return obj.endsOn
}
