package authenticates

import (
	commands_application "steve.care/network/applications/applications/authenticates/commands"
	layers_application "steve.care/network/applications/applications/authenticates/layers"
	links_application "steve.care/network/applications/applications/authenticates/links"
	receipts_application "steve.care/network/applications/applications/authenticates/receipts"
)

type application struct {
	receiptApp  receipts_application.Application
	layerApp    layers_application.Application
	linkApp     links_application.Application
	commandsApp commands_application.Application
}

func createApplication(
	receiptApp receipts_application.Application,
	layerApp layers_application.Application,
	linkApp links_application.Application,
	commandsApp commands_application.Application,
) Application {
	out := application{
		receiptApp:  receiptApp,
		layerApp:    layerApp,
		linkApp:     linkApp,
		commandsApp: commandsApp,
	}

	return &out
}

// Receipts returns the receipt application
func (app *application) Receipts() receipts_application.Application {
	return app.receiptApp
}

// Layers returns the layer application
func (app *application) Layers() layers_application.Application {
	return app.layerApp
}

// Links returns the link application
func (app *application) Links() links_application.Application {
	return app.linkApp
}

// Commands returns the command application
func (app *application) Commands() commands_application.Application {
	return app.commandsApp
}
