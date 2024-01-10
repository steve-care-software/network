package threads

import (
	"errors"
	"fmt"
	"time"

	"steve.care/network/applications/applications/programs/logics"
	"steve.care/network/domain/programs"
	"steve.care/network/domain/programs/logics/threads/executions"
)

type application struct {
	logicApplication logics.Application
	repository       executions.Repository
	service          executions.Service
	builder          executions.Builder
}

func createApplication(
	logicApplication logics.Application,
	repository executions.Repository,
	service executions.Service,
	builder executions.Builder,
) Application {
	out := application{
		logicApplication: logicApplication,
		repository:       repository,
		service:          service,
		builder:          builder,
	}

	return &out
}

// Execute executes the threads of a program
func (app *application) Execute(program programs.Program) error {
	if !program.HasLogic() {
		str := fmt.Sprintf("the program (hash: %s) does NOT have logic and therefore cannot execute threads", program.Hash().String())
		return errors.New(str)
	}

	logic := program.Logic()
	if !logic.HasThreads() {
		str := fmt.Sprintf("the program (hash: %s) contains a logic (hash: %s) that does NOT have threads", program.Hash().String(), logic.Threads().Hash().String())
		return errors.New(str)
	}

	library := logic.Library()
	threadsList := logic.Threads().List()
	for _, oneThread := range threadsList {
		input := oneThread.Input()
		layer := oneThread.Entry()
		beginsOn := time.Now().UTC()
		receipt, err := app.logicApplication.Execute(
			input,
			layer,
			library,
			nil,
		)

		if err != nil {
			return err
		}

		endsOn := time.Now().UTC()
		execution, err := app.builder.Create().
			WithThread(oneThread).
			WithReceipt(receipt).
			BeginsOn(beginsOn).
			EndsOn(endsOn).
			Now()

		if err != nil {
			return err
		}

		err = app.service.Insert(execution)
		if err != nil {
			return err
		}
	}

	return nil
}
