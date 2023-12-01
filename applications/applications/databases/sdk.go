package databases

import "steve.care/network/domain/databases"

// Application repreents the database application
type Application interface {
	OpenInMemory() (databases.Database, error)
	Open(name string) (databases.Database, error)
}
