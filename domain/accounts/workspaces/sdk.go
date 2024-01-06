package workspaces

import (
	"steve.care/network/domain/dashboards"
	"steve.care/network/domain/hash"
)

// Workspace represents a workspace
type Workspace interface {
	Hash() hash.Hash
	Dashboards() dashboards.Dashboards
	Root() dashboards.Dashboard
}
