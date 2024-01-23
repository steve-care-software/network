package blocks

import (
	"errors"

	"steve.care/network/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	content     Content
	result      []byte
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		content:     nil,
		result:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithContent adds content to the builder
func (app *builder) WithContent(content Content) Builder {
	app.content = content
	return app
}

// WithResult adds result to the builder
func (app *builder) WithResult(result []byte) Builder {
	app.result = result
	return app
}

// Now builds a new Block instance
func (app *builder) Now() (Block, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Block instance")
	}

	if app.result != nil && len(app.result) <= 0 {
		app.result = nil
	}

	if app.result == nil {
		return nil, errors.New("the result is mandatory in order to build a Block instance")
	}

	msg := app.content.Hash().Bytes()
	pMinedHash, err := Compute(msg, app.result)
	if err != nil {
		return nil, err
	}

	difficulty := uint(0)
	minedHashBytes := pMinedHash.Bytes()
	firstByte := minedHashBytes[0]
	for i := 0; i < len(minedHashBytes); i++ {
		if firstByte == minedHashBytes[i] {
			difficulty++
			continue
		}

		break
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.content.Hash().Bytes(),
		app.result,
	})

	if err != nil {
		return nil, err
	}

	return createBlock(
		*pHash,
		app.content,
		difficulty,
		*pMinedHash,
		app.result,
	), nil
}
