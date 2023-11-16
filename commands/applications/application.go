package applications

import (
	"steve.care/network/commands/domain/commands"
	"steve.care/network/commands/visitors/applications"
	"steve.care/network/commands/visitors/domain/programs"
	"steve.care/network/commands/visitors/domain/stacks"
)

type application struct {
	visitorApp     applications.Application
	programAdapter programs.Adapter
	commandBuilder commands.CommandBuilder
}

func createApplication(
	visitorApp applications.Application,
	programAdapter programs.Adapter,
	commandBuilder commands.CommandBuilder,
) Application {
	out := application{
		visitorApp:     visitorApp,
		programAdapter: programAdapter,
		commandBuilder: commandBuilder,
	}

	return &out
}

// ExecuteBytes execute bytes
func (app *application) ExecuteBytes(bytes []byte, stack stacks.Stack) (commands.Command, error) {
	program, err := app.programAdapter.ToInstance(bytes)
	if err != nil {
		return nil, err
	}

	return app.Execute(program, stack)
}

// Execute executes the application
func (app *application) Execute(program programs.Program, stack stacks.Stack) (commands.Command, error) {
	result, err := app.visitorApp.Execute(program, stack)
	if err != nil {
		return nil, err
	}

	return app.commandBuilder.Create().
		WithRequest(program).
		WithResponse(result).
		Now()
}
