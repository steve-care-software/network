package peers

import "steve.care/network/domain/programs"

// Application represents the peers application
type Application interface {
	Sync(program programs.Program) error
}
