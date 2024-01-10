package peers

import (
	"net/url"

	applications_blocks "steve.care/network/applications/applications/programs/blocks"
	"steve.care/network/domain/programs"
)

// BlockApplicationBuilder represents a block application builder
type BlockApplicationBuilder interface {
	Create() BlockApplicationBuilder
	WithURL(url *url.URL) BlockApplicationBuilder
	Now() (applications_blocks.Application, error)
}

// Application represents the peers application
type Application interface {
	Execute(program programs.Program) error
}
