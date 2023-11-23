package databases

import (
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/layers"
	"steve.care/network/domain/links"
)

// Database represents a database
type Database interface {
	Begin() (Service, error)
	Repository() Repository
}

// Service represents a service
type Service interface {
	Account() accounts.Service
	Layer() layers.Service
	Link() links.Service
	Commit() error
	Rollack() error
}

// Repository represents a repository
type Repository interface {
	Account() accounts.Repository
	Layer() layers.Repository
	Link() links.Repository
}
