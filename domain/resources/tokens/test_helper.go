package tokens

import (
	"time"

	"steve.care/network/domain/resources/tokens/layers"
)

// NewTokenWithLayerForTests creates a new token with layer for tests
func NewTokenWithLayerForTests(input layers.Layer) Token {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithLayer(input).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
