package tokens

import (
	"errors"
	"strconv"
	"time"

	"steve.care/network/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	content     Content
	pCreatedOn  *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		content:     nil,
		pCreatedOn:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithContent adds a content to the builder
func (app *builder) WithContent(content Content) Builder {
	app.content = content
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Token instance
func (app *builder) Now() (Token, error) {
	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Token instance")
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Token instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(app.pCreatedOn.UTC().Nanosecond())),
		app.content.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createToken(
		*pHash,
		app.content,
		*app.pCreatedOn,
	), nil
}
