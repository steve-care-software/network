package programs

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks"
	"steve.care/network/domain/programs/logics"
)

type builder struct {
	hashAdapter hash.Adapter
	space       []string
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
		space:       nil,
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

// WithSpace adds a space to the builder
func (app *builder) WithSpace(space []string) Builder {
	app.space = space
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
	if app.space != nil && len(app.space) <= 0 {
		app.space = nil
	}

	if app.space == nil {
		return nil, errors.New("the space is mandatory in order to build a Program instance")
	}

	validSpace := []string{}
	for _, oneSpace := range app.space {
		if oneSpace == "" {
			continue
		}

		validSpace = append(validSpace, oneSpace)
	}

	if len(validSpace) <= 0 {
		return nil, errors.New("the space elements must NOT be empty in order to build a Program instance")
	}

	if app.description == "" {
		return nil, errors.New("the description is mandatory in order to build a Program instance")
	}

	data := [][]byte{
		[]byte(app.description),
	}

	for _, oneSpace := range app.space {
		data = append(data, []byte(oneSpace))
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
		return createProgramWithHeadAndLogicAndParent(*pHash, validSpace, app.description, app.head, app.logic, app.parent), nil
	}

	if app.head != nil && app.logic != nil {
		return createProgramWithHeadAndLogic(*pHash, validSpace, app.description, app.head, app.logic), nil
	}

	if app.head != nil && app.parent != nil {
		return createProgramWithHeadAndParent(*pHash, validSpace, app.description, app.head, app.parent), nil
	}

	if app.logic != nil && app.parent != nil {
		return createProgramWithLogicAndParent(*pHash, validSpace, app.description, app.logic, app.parent), nil
	}

	if app.head != nil {
		return createProgramWithHead(*pHash, validSpace, app.description, app.head), nil
	}

	if app.logic != nil {
		return createProgramWithLogic(*pHash, validSpace, app.description, app.logic), nil
	}

	if app.parent != nil {
		return createProgramWithParent(*pHash, validSpace, app.description, app.parent), nil
	}

	return createProgram(*pHash, validSpace, app.description), nil
}
