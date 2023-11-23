package layers

import (
	"errors"
	"strconv"

	"steve.care/network/domain/hash"
)

type instructionBuilder struct {
	hashAdapter hash.Adapter
	isStop      bool
	raiseError  uint
	condition   Condition
	assignment  Assignment
	link        LinkInstruction
	layer       LayerInstruction
}

func createInstructionBuilder(
	hashAdapter hash.Adapter,
) InstructionBuilder {
	out := instructionBuilder{
		hashAdapter: hashAdapter,
		isStop:      false,
		raiseError:  0,
		condition:   nil,
		assignment:  nil,
		link:        nil,
		layer:       nil,
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

// WithAssignment adds an assignment to the builder
func (app *instructionBuilder) WithAssignment(assignment Assignment) InstructionBuilder {
	app.assignment = assignment
	return app
}

// WithLink adds a link to the builder
func (app *instructionBuilder) WithLink(link LinkInstruction) InstructionBuilder {
	app.link = link
	return app
}

// WithLayer adds a layer to the builder
func (app *instructionBuilder) WithLayer(layer LayerInstruction) InstructionBuilder {
	app.layer = layer
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

	if app.assignment != nil {
		data = append(data, []byte("assignment"))
		data = append(data, app.assignment.Hash().Bytes())
	}

	if app.link != nil {
		data = append(data, []byte("link"))
		data = append(data, app.link.Hash().Bytes())
	}

	if app.layer != nil {
		data = append(data, []byte("layer"))
		data = append(data, app.layer.Hash().Bytes())
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

	if app.link != nil {
		return createInstructionWithLink(*pHash, app.link), nil
	}

	if app.layer != nil {
		return createInstructionWithLayer(*pHash, app.layer), nil
	}

	return createInstructionWithAssignment(*pHash, app.assignment), nil
}
