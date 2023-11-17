package applications

import (
	"steve.care/network/commands/visitors/admins/domain/accounts"
	"steve.care/network/commands/visitors/admins/domain/programs"
	"steve.care/network/commands/visitors/admins/domain/stacks"
)

type application struct {
	programAdapter            programs.Adapter
	stackBuilder              stacks.Builder
	stackFrameFactory         stacks.FrameFactory
	stackFrameBuilder         stacks.FrameBuilder
	stackAssignmentBuilder    stacks.AssignmentBuilder
	stackAssignableBuilder    stacks.AssignableBuilder
	stackCreateAccountBuilder stacks.CreateAccountBuilder
	accountRepository         accounts.Repository
	accountBuilder            accounts.Builder
}

func createApplication(
	programAdapter programs.Adapter,
	stackBuilder stacks.Builder,
	stackFrameFactory stacks.FrameFactory,
	stackFrameBuilder stacks.FrameBuilder,
	stackAssignmentBuilder stacks.AssignmentBuilder,
	stackAssignableBuilder stacks.AssignableBuilder,
	stackCreateAccountBuilder stacks.CreateAccountBuilder,
	accountRepository accounts.Repository,
	accountBuilder accounts.Builder,
) Application {
	out := application{
		programAdapter:            programAdapter,
		stackBuilder:              stackBuilder,
		stackFrameFactory:         stackFrameFactory,
		stackFrameBuilder:         stackFrameBuilder,
		stackAssignmentBuilder:    stackAssignmentBuilder,
		stackAssignableBuilder:    stackAssignableBuilder,
		stackCreateAccountBuilder: stackCreateAccountBuilder,
		accountRepository:         accountRepository,
		accountBuilder:            accountBuilder,
	}

	return &out
}

// ExecuteBytes execute the application using the passed bytes
func (app *application) ExecuteBytes(bytes []byte, stack stacks.Stack) (stacks.Stack, error) {
	program, err := app.programAdapter.ToInstance(bytes)
	if err != nil {
		return nil, err
	}

	return app.Execute(program, stack)
}

// Execute executes the application
func (app *application) Execute(program programs.Program, stack stacks.Stack) (stacks.Stack, error) {
	/*stackFrames := stack.List()
	stackFrames = append(stackFrames, app.stackFrameFactory.Create())
	retStack, err := app.stackBuilder.Create().
		WithList(stackFrames).
		Now()

	if err != nil {
		return nil, err
	}

	instructions := program.Instructions()
	return app.instructions(instructions, retStack)*/
	return nil, nil
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
	/*last := stack.Last()
	lastAssignments := last.List()
	if instruction.IsAssignment() {
		assignment := instruction.Assignment()
		stackAssignment, err := app.assignment(assignment, stack)
		if err != nil {
			return nil, err
		}

		lastAssignments = append(lastAssignments, stackAssignment)
	}

	if instruction.IsDeleteAuthorized() {

	}

	if instruction.IsCreateAdmin() {
		creator := stack.Memory().Authorized()
		credentials := instruction.CreateAdmin()
		username := credentials.Username()
		exists, err := app.accountRepository.Exists(username)
		if err != nil {
			return nil, err
		}

		if exists {
			assignable, err := app.stackAssignableBuilder.Create().
				WithError(stacks.AccountNameAlreadyExists).
				Now()

			if err != nil {
				return nil, err
			}

		}

		account, err := app.accountBuilder.Create().
			WithCreator(creator).
			WithUsername(username).
			Now()

		if err != nil {
			return nil, err
		}

		password := credentials.Password()
		createAccount, err := app.stackCreateAccountBuilder.Create().
			WithAccount(account).
			WithPassword(password).
			Now()

		if err != nil {
			return nil, err
		}

		builder.WithCreateAccount(createAccount)
	}

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
		Now()*/

	return nil, nil
}

func (app *application) assignment(assignment programs.Assignment, stack stacks.Stack) (stacks.Assignment, error) {
	assignable := assignment.Assignable()
	stackAssignable, err := app.assignable(assignable, stack)
	if err != nil {
		return nil, err
	}

	name := assignment.Name()
	return app.stackAssignmentBuilder.Create().
		WithName(name).
		WithAssignable(stackAssignable).
		Now()
}

func (app *application) assignable(assignable programs.Assignable, stack stacks.Stack) (stacks.Assignable, error) {
	builder := app.stackAssignableBuilder.Create()
	if assignable.IsHasIdentities() {
		account := stack.Memory().Authorized()
		value := account.HasIdentities()
		builder.WithBool(value)
	}

	if assignable.IsListIdentities() {
		account := stack.Memory().Authorized()
		if !account.HasIdentities() {
			return builder.WithError(stacks.AuthorizedAccountDoNotContainIdentitiesError).
				Now()
		}

		identities := account.Identities()
		builder.WithIdentities(identities)
	}

	return builder.Now()
}
