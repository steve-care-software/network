package programs

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks"
	"steve.care/network/domain/programs/logics"
)

type builder struct {
	hashAdapter hash.Adapter
	description string
	head        blocks.Block
	logic       logics.Logic
	metadata    MetaData
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		description: "",
		head:        nil,
		logic:       nil,
		metadata:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
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

// WithMetaData adds a metadata to the builder
func (app *builder) WithMetaData(metadata MetaData) Builder {
	app.metadata = metadata
	return app
}

// Now builds a new Program instance
func (app *builder) Now() (Program, error) {
	if app.description == "" {
		return nil, errors.New("the description is mandatory in order to build a Program instance")
	}

	data := [][]byte{
		[]byte(app.description),
	}

	if app.head != nil {
		data = append(data, app.head.Hash().Bytes())
	}

	if app.logic != nil {
		data = append(data, app.logic.Hash().Bytes())
	}

	if app.metadata != nil {
		data = append(data, app.metadata.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.head != nil && app.logic != nil && app.metadata != nil {
		return createProgramWithHeadAndLogicAndMetaData(*pHash, app.description, app.head, app.logic, app.metadata), nil
	}

	if app.head != nil && app.logic != nil {
		return createProgramWithHeadAndLogic(*pHash, app.description, app.head, app.logic), nil
	}

	if app.head != nil && app.metadata != nil {
		return createProgramWithHeadAndMetaData(*pHash, app.description, app.head, app.metadata), nil
	}

	if app.logic != nil && app.metadata != nil {
		return createProgramWithLogicAndMetaData(*pHash, app.description, app.logic, app.metadata), nil
	}

	if app.head != nil {
		return createProgramWithHead(*pHash, app.description, app.head), nil
	}

	if app.logic != nil {
		return createProgramWithLogic(*pHash, app.description, app.logic), nil
	}

	if app.metadata != nil {
		return createProgramWithMetaData(*pHash, app.description, app.metadata), nil
	}

	return createProgram(*pHash, app.description), nil
}
