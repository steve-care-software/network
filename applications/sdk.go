package applications

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/results"
	"steve.care/network/domain/stacks"
)

// Application represents a stencil application
type Application interface {
	Execute(hash hash.Hash, stack stacks.Stack) (results.Result, error)
}
