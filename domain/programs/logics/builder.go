package logics

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/programs/logics/suites"
	"steve.care/network/domain/programs/logics/threads"
)

type builder struct {
	hashAdapter hash.Adapter
	entry       layers.Layer
	library     libraries.Library
	suites      suites.Suites
	threads     threads.Threads
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		entry:       nil,
		library:     nil,
		suites:      nil,
		threads:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithEntry adds an entry to the builder
func (app *builder) WithEntry(entry layers.Layer) Builder {
	app.entry = entry
	return app
}

// WithLibrary adds a library to the builder
func (app *builder) WithLibrary(library libraries.Library) Builder {
	app.library = library
	return app
}

// WithSuites adds a suites to the builder
func (app *builder) WithSuites(suites suites.Suites) Builder {
	app.suites = suites
	return app
}

// WithThreads adds a threads to the builder
func (app *builder) WithThreads(threads threads.Threads) Builder {
	app.threads = threads
	return app
}

// Now builds a new Logic instance
func (app *builder) Now() (Logic, error) {
	if app.entry == nil {
		return nil, errors.New("the entry is mandatory in order to build a Logic instance")
	}

	if app.library == nil {
		return nil, errors.New("the library is mandatory in order to build a Logic instance")
	}

	data := [][]byte{
		app.entry.Hash().Bytes(),
		app.library.Hash().Bytes(),
	}

	if app.suites != nil {
		data = append(data, app.suites.Hash().Bytes())
	}

	if app.threads != nil {
		data = append(data, app.threads.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.suites != nil && app.threads != nil {
		return createLogicWithSuitesAndThreads(*pHash, app.entry, app.library, app.suites, app.threads), nil
	}

	if app.suites != nil {
		return createLogicWithSuites(*pHash, app.entry, app.library, app.suites), nil
	}

	if app.threads != nil {
		return createLogicWithThreads(*pHash, app.entry, app.library, app.threads), nil
	}

	return createLogic(*pHash, app.entry, app.library), nil
}
