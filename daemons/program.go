package daemons

import (
	"steve.care/network/applications/applications/programs/graphs"
	"steve.care/network/applications/applications/programs/logics"
	"steve.care/network/applications/applications/programs/peers"
	"steve.care/network/applications/applications/programs/threads"
)

type program struct {
	graphApplication  graphs.Application
	threadApplication threads.Application
	peersApplication  peers.Application
	isActive          bool
}

func createApplication(
	graphApplication graphs.Application,
	peersApplication peers.Application,
	logicApplication logics.Application,
) Application {
	out := program{
		graphApplication: graphApplication,
		peersApplication: peersApplication,
		isActive:         false,
	}

	return &out
}

// Start starts the application
func (app *program) Start() error {
	app.isActive = true
	for {
		if !app.isActive {
			return nil
		}

		// fetch the active programs hash list:
		isActive := true
		retHashes, err := app.graphApplication.List(&isActive)
		if err != nil {
			return err
		}

		for _, oneHash := range retHashes {
			// retrieve the program from hash:
			program, err := app.graphApplication.Retrieve(oneHash)
			if err != nil {
				return err
			}

			// sync the program with peers:
			err = app.peersApplication.Execute(program)
			if err != nil {
				return err
			}

			// execute the program's threads:
			err = app.threadApplication.Execute(program)
			if err != nil {
				return err
			}
		}

	}
}

// Stop stops the application
func (app *program) Stop() {
	app.isActive = false
}
