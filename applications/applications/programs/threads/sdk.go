package threads

import "steve.care/network/domain/programs"

// Application represents the threads application
type Application interface {
	Execute(program programs.Program) error
}
