package sqllites

import "steve.care/network/applications/databases"

// NewApplication creates a new application
func NewApplication(
	basePath string,
) databases.Application {
	return createApplication(
		basePath,
	)
}
