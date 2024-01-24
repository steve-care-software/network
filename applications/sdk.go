package applications

import (
	"steve.care/network/applications/applications"
)

// Application represents the core application
type Application interface {
	Init(name string) error
	Begin(name string) (applications.Application, error)
	Commit() error
	Rollback() error
	Close() error
}
