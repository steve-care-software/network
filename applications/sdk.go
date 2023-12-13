package applications

import (
	"steve.care/network/applications/applications"
)

// Application represents the core application
type Application interface {
	Init(name string, script string) (applications.Application, error)
	InitInMemory(script string) (applications.Application, error)
	Begin(name string) (applications.Application, error)
	BeginInMemory() (applications.Application, error)
	Commit() error
	Cancel() error
	Rollback() error
	Close() error
}
