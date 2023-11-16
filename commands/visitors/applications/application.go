package applications

import (
	"steve.care/network/commands/visitors/admins/domain/accounts"
	"steve.care/network/commands/visitors/domain/programs"
	"steve.care/network/commands/visitors/domain/stacks"
)

type application struct {
	stackBuilder           stacks.Builder
	stackFrameFactory      stacks.FrameFactory
	stackFrameBuilder      stacks.FrameBuilder
	stackAssignmentBuilder stacks.AssignmentBuilder
	stackAssignableBuilder stacks.AssignableBuilder
	accountRepository      accounts.Repository
}

func createApplication(
	stackBuilder stacks.Builder,
	stackFrameFactory stacks.FrameFactory,
	stackFrameBuilder stacks.FrameBuilder,
	stackAssignmentBuilder stacks.AssignmentBuilder,
	stackAssignableBuilder stacks.AssignableBuilder,
	accountRepository accounts.Repository,
) Application {
	out := application{
		stackBuilder:           stackBuilder,
		stackFrameFactory:      stackFrameFactory,
		stackFrameBuilder:      stackFrameBuilder,
		stackAssignmentBuilder: stackAssignmentBuilder,
		stackAssignableBuilder: stackAssignableBuilder,
		accountRepository:      accountRepository,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(program programs.Program, stack stacks.Stack) (stacks.Stack, error) {
	stackFrames := stack.List()
	stackFrames = append(stackFrames, app.stackFrameFactory.Create())
	retStack, err := app.stackBuilder.Create().
		WithList(stackFrames).
		Now()

	if err != nil {
		return nil, err
	}

	instructions := program.Instructions()
	return app.instructions(instructions, retStack)
}

func (app *application) instructions(instructions programs.Instructions, stack stacks.Stack) (stacks.Stack, error) {
	lastStack := stack
	list := instructions.List()
	for _, oneInstruction := range list {
		retStack, err := app.instruction(oneInstruction, lastStack)
		if err != nil {
			return nil, err
		}

		lastStack = retStack
	}

	return lastStack, nil
}

func (app *application) instruction(instruction programs.Instruction, stack stacks.Stack) (stacks.Stack, error) {
	assignment := instruction.Assignment()
	return app.assignment(assignment, stack)
}

func (app *application) assignment(assignment programs.Assignment, stack stacks.Stack) (stacks.Stack, error) {
	assignable := assignment.Assignable()
	stackAssignable, err := app.assignable(assignable, stack)
	if err != nil {
		return nil, err
	}

	name := assignment.Name()
	stackAssignment, err := app.stackAssignmentBuilder.Create().
		WithName(name).
		WithAssignable(stackAssignable).
		Now()

	if err != nil {
		return nil, err
	}

	last := stack.Last()
	lastAssignments := last.List()
	lastAssignments = append(lastAssignments, stackAssignment)
	updatedFrame, err := app.stackFrameBuilder.Create().
		WihtList(lastAssignments).
		Now()

	if err != nil {
		return nil, err
	}

	bodyFrames := stack.Body()
	bodyFrames = append(bodyFrames, updatedFrame)
	return app.stackBuilder.Create().
		WithList(bodyFrames).
		Now()
}

func (app *application) assignable(assignable programs.Assignable, stack stacks.Stack) (stacks.Assignable, error) {
	if assignable.IsListNames() {
		names, err := app.accountRepository.List()
		if err != nil {
			return nil, err
		}

		return app.stackAssignableBuilder.Create().
			WithStringList(names).
			Now()
	}

	if assignable.IsAuthorize() {
		credentials := assignable.Authorize()
		account, err := app.accountRepository.Retrieve(credentials)
		if err != nil {
			return app.stackAssignableBuilder.Create().
				WithError(stacks.CouldNotAuthorizeError).
				Now()
		}

		return app.stackAssignableBuilder.Create().
			WithAuthorize(account).
			Now()
	}

	credentials := assignable.Create()
	username := credentials.Username()
	exists, err := app.accountRepository.Exists(username)
	if err != nil {
		return nil, err
	}

	if exists {
		return app.stackAssignableBuilder.Create().
			WithError(stacks.AccountNameAlreadyExists).
			Now()
	}

	return app.stackAssignableBuilder.Create().
		WithCreate(credentials).
		Now()
}
