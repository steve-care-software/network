package cruds

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	Now() (Application, error)
}

// Application represents the program CRUD application
type Application interface {
	Root() (programs.Program, error)
	List(isActive bool) ([]hash.Hash, error)
	Children(path []string) ([]string, error)
	Height(path []string) (*uint, error)
	Revision(path []string, height uint) (hash.Hash, error)
	Retrieve(program hash.Hash) (programs.Program, error)
	Insert(path []string, description string) error
	Rewind(path []string) error
	Delete(path []string) error
}
