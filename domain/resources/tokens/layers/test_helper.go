package layers

import "steve.care/network/domain/receipts/commands/layers"

// NewLayerWithBytesReferenceForTests creates a new layer with bytes reference for tests
func NewLayerWithBytesReferenceForTests(input layers.BytesReference) Layer {
	ins, err := NewBuilder().Create().WithBytesReference(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
