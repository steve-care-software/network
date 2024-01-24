package applications

import (
	"steve.care/network/applications/applications"
)

// Application represents the core application
type Application interface {
	Init() error
	Begin(name string) (applications.Application, error)
	Commit() error
	Rollback() error
	Close() error
}
