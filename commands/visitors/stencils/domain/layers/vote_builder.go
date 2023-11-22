package layers

import (
	"errors"

	"steve.care/network/libraries/hash"
)

type voteBuilder struct {
	hashAdapter hash.Adapter
	ring        string
	message     BytesReference
}

func createVoteBuilder(
	hashAdapter hash.Adapter,
) VoteBuilder {
	out := voteBuilder{
		hashAdapter: hashAdapter,
		ring:        "",
		message:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *voteBuilder) Create() VoteBuilder {
	return createVoteBuilder(
		app.hashAdapter,
	)
}

// WithRing adds a ring to the builder
func (app *voteBuilder) WithRing(ring string) VoteBuilder {
	app.ring = ring
	return app
}

// WithMessage adds a message to the builder
func (app *voteBuilder) WithMessage(message BytesReference) VoteBuilder {
	app.message = message
	return app
}

// Now builds a new Vote instance
func (app *voteBuilder) Now() (Vote, error) {
	if app.ring != "" {
		return nil, errors.New("the ring variable is mandatory in order to build a Vote instance")
	}

	if app.message != nil {
		return nil, errors.New("the message is mandatory in order to build a Vote instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.ring),
		app.message.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createVote(*pHash, app.ring, app.message), nil
}
