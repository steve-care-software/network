package entries

import (
	"errors"

	"steve.care/network/domain/databases/criterias"
)

type builder struct {
	criteria criterias.Criteria
	fields   []string
}

func createBuilder() Builder {
	out := builder{
		criteria: nil,
		fields:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithCriteria adds a criteria to the builder
func (app *builder) WithCriteria(criteria criterias.Criteria) Builder {
	app.criteria = criteria
	return app
}

// WithFields adds a fields to the builder
func (app *builder) WithFields(fields []string) Builder {
	app.fields = fields
	return app
}

// Now builds a new Entry instance
func (app *builder) Now() (Entry, error) {
	if app.criteria == nil {
		return nil, errors.New("the criteria is mandatory in order to build an Entry instance")
	}

	if app.fields != nil && len(app.fields) <= 0 {
		app.fields = nil
	}

	if app.fields == nil {
		return nil, errors.New("the fields is mandatory in order to build an Entry instance")
	}

	return createEntry(
		app.criteria,
		app.fields,
	), nil
}
