package applications

import "steve.care/network/applications/applications"

// Application represents the core application
type Application interface {
	Begin() (applications.Application, error)
	Commit(applications.Application) error
	Cancel(applications.Application) error
	Rollback(applications.Application) error
}
