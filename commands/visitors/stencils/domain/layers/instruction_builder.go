package layers

import (
	"errors"
	"strconv"

	"steve.care/network/libraries/hash"
)

type instructionBuilder struct {
	hashAdapter hash.Adapter
	isStop      bool
	raiseError  uint
	condition   Condition
	save        Layer
	assignment  Assignment
}

func createInstructionBuilder(
	hashAdapter hash.Adapter,
) InstructionBuilder {
	out := instructionBuilder{
		hashAdapter: hashAdapter,
		isStop:      false,
		raiseError:  0,
		condition:   nil,
		save:        nil,
		assignment:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder(
		app.hashAdapter,
	)
}

// WithRaiseError raises an error in the builder
func (app *instructionBuilder) WithRaiseError(raiseError uint) InstructionBuilder {
	app.raiseError = raiseError
	return app
}

// WithCondition adds a condition to the builder
func (app *instructionBuilder) WithCondition(condition Condition) InstructionBuilder {
	app.condition = condition
	return app
}

// WthSave saves a layer to the builder
func (app *instructionBuilder) WthSave(save Layer) InstructionBuilder {
	app.save = save
	return app
}

// WithAssignment adds an assignment to the builder
func (app *instructionBuilder) WithAssignment(assignment Assignment) InstructionBuilder {
	app.assignment = assignment
	return app
}

// IsStop flags the builder as a stop
func (app *instructionBuilder) IsStop() InstructionBuilder {
	app.isStop = true
	return app
}

// Now builds a new Instruction instance
func (app *instructionBuilder) Now() (Instruction, error) {
	data := [][]byte{}
	if app.isStop {
		data = append(data, []byte("isStop"))
	}

	if app.raiseError > 0 {
		data = append(data, []byte("raiseError"))
		data = append(data, []byte(strconv.Itoa(int(app.raiseError))))
	}

	if app.condition != nil {
		data = append(data, []byte("condition"))
		data = append(data, app.condition.Hash().Bytes())
	}

	if app.save != nil {
		data = append(data, []byte("save"))
		data = append(data, app.save.Hash().Bytes())
	}

	if app.assignment != nil {
		data = append(data, []byte("assignment"))
		data = append(data, app.assignment.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Instruction is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.isStop {
		return createInstructionWithIsStop(*pHash), nil
	}

	if app.raiseError > 0 {
		return createInstructionWithRaiseError(*pHash, app.raiseError), nil
	}

	if app.condition != nil {
		return createInstructionWithCondition(*pHash, app.condition), nil
	}

	if app.save != nil {
		return createInstructionWithSave(*pHash, app.save), nil
	}

	return createInstructionWithAssignment(*pHash, app.assignment), nil
}
