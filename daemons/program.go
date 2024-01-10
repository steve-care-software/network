package daemons

import (
	"steve.care/network/applications/applications/programs"
	"steve.care/network/applications/applications/programs/logics"
	"steve.care/network/applications/applications/programs/logics/threads"
	"steve.care/network/applications/applications/programs/peers"
)

type program struct {
	programApplication programs.Application
	threadApplication  threads.Application
	peersApplication   peers.Application
	isActive           bool
}

func createApplication(
	programApplication programs.Application,
	peersApplication peers.Application,
	logicApplication logics.Application,
) Application {
	out := program{
		programApplication: programApplication,
		peersApplication:   peersApplication,
		isActive:           false,
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
		hashesList, err := app.programApplication.List(true)
		if err != nil {
			return err
		}

		for _, oneHash := range hashesList {
			// retrieve the program from hash:
			program, err := app.programApplication.Retrieve(oneHash)
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
