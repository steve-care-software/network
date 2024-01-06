package programs

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks"
	"steve.care/network/domain/programs/logics"
)

type builder struct {
	hashAdapter hash.Adapter
	name        string
	description string
	head        blocks.Block
	logic       logics.Logic
	parent      Program
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		name:        "",
		description: "",
		head:        nil,
		logic:       nil,
		parent:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// WithLogic adds a logic to the builder
func (app *builder) WithLogic(logic logics.Logic) Builder {
	app.logic = logic
	return app
}

// WithHead adds a head to the builder
func (app *builder) WithHead(head blocks.Block) Builder {
	app.head = head
	return app
}

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent Program) Builder {
	app.parent = parent
	return app
}

// Now builds a new Program instance
func (app *builder) Now() (Program, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Program instance")
	}

	if app.description == "" {
		return nil, errors.New("the description is mandatory in order to build a Program instance")
	}

	data := [][]byte{
		[]byte(app.name),
		[]byte(app.description),
	}

	if app.head != nil {
		data = append(data, app.head.Hash().Bytes())
	}

	if app.logic != nil {
		data = append(data, app.logic.Hash().Bytes())
	}

	if app.parent != nil {
		data = append(data, app.parent.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.head != nil && app.logic != nil && app.parent != nil {
		return createProgramWithHeadAndLogicAndParent(*pHash, app.name, app.description, app.head, app.logic, app.parent), nil
	}

	if app.head != nil && app.logic != nil {
		return createProgramWithHeadAndLogic(*pHash, app.name, app.description, app.head, app.logic), nil
	}

	if app.head != nil && app.parent != nil {
		return createProgramWithHeadAndParent(*pHash, app.name, app.description, app.head, app.parent), nil
	}

	if app.logic != nil && app.parent != nil {
		return createProgramWithLogicAndParent(*pHash, app.name, app.description, app.logic, app.parent), nil
	}

	if app.head != nil {
		return createProgramWithHead(*pHash, app.name, app.description, app.head), nil
	}

	if app.logic != nil {
		return createProgramWithLogic(*pHash, app.name, app.description, app.logic), nil
	}

	if app.parent != nil {
		return createProgramWithParent(*pHash, app.name, app.description, app.parent), nil
	}

	return createProgram(*pHash, app.name, app.description), nil
}
