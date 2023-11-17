package applications

import (
	"errors"
	"fmt"

	"steve.care/network/commands/visitors/admins/identities/dashboards/stencils/layers/domain/layers"
	"steve.care/network/commands/visitors/admins/identities/dashboards/stencils/layers/domain/programs"
	"steve.care/network/commands/visitors/admins/identities/dashboards/stencils/layers/domain/stacks"
)

type application struct {
	layerRepository          layers.Repository
	layerService             layers.Service
	programAdapter           programs.Adapter
	stackBuilder             stacks.Builder
	stackMemoryBuilder       stacks.MemoryBuilder
	stackFramesBuilder       stacks.FramesBuilder
	stackFrameBuilder        stacks.FrameBuilder
	stackInstructionsBuilder stacks.InstructionsBuilder
	stackInstructionBuilder  stacks.InstructionBuilder
	stackMoveBuilder         stacks.MoveBuilder
	stackAssignmentsBuilder  stacks.AssignmentsBuilder
	stackAssignmentBuilder   stacks.AssignmentBuilder
	stackAssignableBuilder   stacks.AssignableBuilder
	stackResourceBuilder     stacks.ResourceBuilder
}

func createApplication(
	layerRepository layers.Repository,
	layerService layers.Service,
	programAdapter programs.Adapter,
	stackBuilder stacks.Builder,
	stackMemoryBuilder stacks.MemoryBuilder,
	stackFramesBuilder stacks.FramesBuilder,
	stackFrameBuilder stacks.FrameBuilder,
	stackInstructionsBuilder stacks.InstructionsBuilder,
	stackInstructionBuilder stacks.InstructionBuilder,
	stackMoveBuilder stacks.MoveBuilder,
	stackAssignmentsBuilder stacks.AssignmentsBuilder,
	stackAssignmentBuilder stacks.AssignmentBuilder,
	stackAssignableBuilder stacks.AssignableBuilder,
	stackResourceBuilder stacks.ResourceBuilder,
) Application {
	out := application{
		layerRepository:          layerRepository,
		layerService:             layerService,
		programAdapter:           programAdapter,
		stackBuilder:             stackBuilder,
		stackMemoryBuilder:       stackMemoryBuilder,
		stackFramesBuilder:       stackFramesBuilder,
		stackFrameBuilder:        stackFrameBuilder,
		stackInstructionsBuilder: stackInstructionsBuilder,
		stackInstructionBuilder:  stackInstructionBuilder,
		stackMoveBuilder:         stackMoveBuilder,
		stackAssignmentsBuilder:  stackAssignmentsBuilder,
		stackAssignmentBuilder:   stackAssignmentBuilder,
		stackAssignableBuilder:   stackAssignableBuilder,
		stackResourceBuilder:     stackResourceBuilder,
	}

	return &out
}

// ExecuteBytes executes the bytes
func (app *application) ExecuteBytes(bytes []byte, stack stacks.Stack) (stacks.Stack, error) {
	program, err := app.programAdapter.ToInstance(bytes)
	if err != nil {
		return nil, err
	}

	return app.Execute(program, stack)
}

// Execute executes the program
func (app *application) Execute(program programs.Program, stack stacks.Stack) (stacks.Stack, error) {
	instructions := program.Instructions()
	return app.executeInstructions(instructions, stack)
}

func (app *application) executeInstructions(instructions programs.Instructions, stack stacks.Stack) (stacks.Stack, error) {
	updatedStack := stack
	list := instructions.List()
	for _, oneInstruction := range list {
		retStack, err := app.executeInstruction(oneInstruction, updatedStack)
		if err != nil {
			return nil, err
		}

		updatedStack = retStack
	}

	return updatedStack, nil
}

func (app *application) executeInstruction(instruction programs.Instruction, stack stacks.Stack) (stacks.Stack, error) {
	currentFrameInstructionList := []stacks.Instruction{}
	last := stack.Last()
	if last.HasInstructions() {
		currentFrameInstructionList = last.Instructions().List()
	}

	currentFrameAssignments := []stacks.Assignment{}
	if last.HasAssignments() {
		currentFrameAssignments = last.Assignments().List()
	}

	if instruction.IsAssignment() {
		assignment := instruction.Assignment()
		stackAssignment, err := app.executeAssignment(assignment)
		if err != nil {
			return nil, err
		}

		currentFrameAssignments = append(currentFrameAssignments, stackAssignment)
	}

	if instruction.IsDelete() {
		path := instruction.Delete()
		stackInstruction, err := app.executeDelete(path)
		if err != nil {
			return nil, err
		}

		currentFrameInstructionList = append(currentFrameInstructionList, stackInstruction)
	}

	if instruction.IsMove() {
		move := instruction.Move()
		stackInstruction, err := app.executeMove(move)
		if err != nil {
			return nil, err
		}

		currentFrameInstructionList = append(currentFrameInstructionList, stackInstruction)
	}

	if instruction.IsSave() {
		save := instruction.Save()
		stackInstruction, err := app.executeSave(save)
		if err != nil {
			return nil, err
		}

		currentFrameInstructionList = append(currentFrameInstructionList, stackInstruction)
	}

	updatedFrameBuilder := app.stackFrameBuilder.Create()
	if len(currentFrameInstructionList) > 0 {
		updatedStackInstructions, err := app.stackInstructionsBuilder.Create().
			WithInstructions(currentFrameInstructionList).
			Now()

		if err != nil {
			return nil, err
		}

		updatedFrameBuilder.WithInstructions(updatedStackInstructions)
	}

	if len(currentFrameAssignments) > 0 {
		updatedStackAssignments, err := app.stackAssignmentsBuilder.Create().
			WithList(currentFrameAssignments).
			Now()

		if err != nil {
			return nil, err
		}

		updatedFrameBuilder.WithAssignments(updatedStackAssignments)
	}

	updatedFrame, err := updatedFrameBuilder.Now()
	if err != nil {
		return nil, err
	}

	body := stack.Body().List()
	framesList := append(body, updatedFrame)
	updatedFrames, err := app.stackFramesBuilder.Create().
		WithList(framesList).
		Now()

	if err != nil {
		return nil, err
	}

	memory := stack.Memory()
	return app.stackBuilder.Create().
		WithFrames(updatedFrames).
		WithMemory(memory).
		Now()
}

func (app *application) executeDelete(path []string) (stacks.Instruction, error) {
	exists, err := app.layerRepository.Exists(path)
	if err != nil {
		return nil, err
	}

	if !exists {
		return app.stackInstructionBuilder.Create().
			WithErrorCode(stacks.LayerDoesNotExistsError).
			Now()
	}

	return app.stackInstructionBuilder.Create().
		WithDelete(path).
		Now()
}

func (app *application) executeMove(move programs.Move) (stacks.Instruction, error) {
	from := move.From()
	fromExists, err := app.layerRepository.Exists(from)
	if err != nil {
		return nil, err
	}

	if !fromExists {
		return app.stackInstructionBuilder.Create().
			WithErrorCode(stacks.LayerDoesNotExistsError).
			Now()
	}

	to := move.To()
	toExists, err := app.layerRepository.Exists(to)
	if err != nil {
		return nil, err
	}

	if toExists {
		return app.stackInstructionBuilder.Create().
			WithErrorCode(stacks.LayerAlreadyExistsError).
			Now()
	}

	stackMove, err := app.stackMoveBuilder.Create().
		From(from).
		To(to).
		Now()

	if err != nil {
		return nil, err
	}

	return app.stackInstructionBuilder.Create().
		WithMove(stackMove).
		Now()
}

func (app *application) executeSave(save programs.Save) (stacks.Instruction, error) {
	path := save.Path()
	layer := save.Layer()
	exists, err := app.layerRepository.Exists(path)
	if err != nil {
		return nil, err
	}

	if exists {
		return app.stackInstructionBuilder.Create().
			WithErrorCode(stacks.LayerAlreadyExistsError).
			Now()
	}

	stackResource, err := app.stackResourceBuilder.Create().
		WithPath(path).
		WithLayer(layer).
		Now()

	if err != nil {
		return nil, err
	}

	return app.stackInstructionBuilder.Create().
		WithSave(stackResource).
		Now()
}

func (app *application) executeAssignment(assignment programs.Assignment) (stacks.Assignment, error) {
	assignable := assignment.Assignable()
	stackAssigable, err := app.executeAssignable(assignable)
	if err != nil {
		return nil, err
	}

	name := assignment.Name()
	return app.stackAssignmentBuilder.Create().
		WithName(name).
		WithAssignable(stackAssigable).
		Now()
}

func (app *application) executeAssignable(assignable programs.Assignable) (stacks.Assignable, error) {
	builder := app.stackAssignableBuilder.Create()
	if assignable.IsExists() {
		path := assignable.Exists()
		exists, err := app.layerRepository.Exists(path)
		if err != nil {
			return nil, err
		}

		builder.WithBool(exists)
	}

	if assignable.IsList() {
		path := assignable.List()
		list, err := app.layerRepository.List(path)
		if err != nil {
			return nil, err
		}

		builder.WithStringList(list)
	}

	if assignable.IsDir() {
		path := assignable.Dir()
		dir, err := app.layerRepository.Dir(path)
		if err != nil {
			return nil, err
		}

		builder.WithStringList(dir)
	}

	if assignable.IsRetrieve() {
		path := assignable.Retrieve()
		layer, err := app.layerRepository.Retrieve(path)
		if err != nil {
			return nil, err
		}

		builder.WithLayer(layer)
	}

	return builder.Now()
}

// Process process the stack
func (app *application) Process(context uint, stack stacks.Stack) error {
	if !stack.HasInstructions() {
		return nil
	}

	instructions := stack.Instructions()
	return app.processInstructions(context, instructions)
}

func (app *application) processInstructions(context uint, instructions stacks.Instructions) error {
	list := instructions.List()
	for idx, oneInstruction := range list {
		err := app.processInstruction(context, oneInstruction, idx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *application) processInstruction(context uint, instruction stacks.Instruction, index int) error {
	if instruction.IsSave() {
		save := instruction.Save()
		path := save.Path()
		layer := save.Layer()
		return app.layerService.Insert(
			context,
			layer,
			path,
		)
	}

	if instruction.IsDelete() {
		path := instruction.Delete()
		return app.layerService.Delete(
			context,
			path,
		)
	}

	if instruction.IsMove() {
		move := instruction.Move()
		from := move.From()
		to := move.To()
		retLayer, err := app.layerRepository.Retrieve(from)
		if err != nil {
			return err
		}

		err = app.layerService.Insert(
			context,
			retLayer,
			to,
		)

		if err != nil {
			return err
		}

		return app.layerService.Delete(
			context,
			from,
		)
	}

	pCode := instruction.ErrorCode()
	switch *pCode {
	case stacks.LayerAlreadyExistsError:
		str := fmt.Sprintf("the layer already exists, code: %d", *pCode)
		return errors.New(str)
	case stacks.LayerDoesNotExistsError:
		str := fmt.Sprintf("the layer does not exists, code: %d", *pCode)
		return errors.New(str)
	default:
		str := fmt.Sprintf("the error code (%d) is invalid", *pCode)
		return errors.New(str)
	}
}
