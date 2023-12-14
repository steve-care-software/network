package layers

import "steve.care/network/domain/hash"

type layer struct {
	hash         hash.Hash
	instructions Instructions
	output       Output
	input        string
}

func createLayer(
	hash hash.Hash,
	instructions Instructions,
	output Output,
) Layer {
	return createLayerInternally(
		hash,
		instructions,
		output,
		"",
	)
}

func createLayerWithInput(
	hash hash.Hash,
	instructions Instructions,
	output Output,
	input string,
) Layer {
	return createLayerInternally(
		hash,
		instructions,
		output,
		input,
	)
}

func createLayerInternally(
	hash hash.Hash,
	instructions Instructions,
	output Output,
	input string,
) Layer {
	out := layer{
		hash:         hash,
		instructions: instructions,
		output:       output,
		input:        input,
	}

	return &out
}

// Hash returns the hash
func (obj *layer) Hash() hash.Hash {
	return obj.hash
}

// Instructions returns the instructions
func (obj *layer) Instructions() Instructions {
	return obj.instructions
}

// Output returns the output
func (obj *layer) Output() Output {
	return obj.output
}

// HasInput returns true if there is an input, false otherwise
func (obj *layer) HasInput() bool {
	return obj.input != ""
}

// Input returns the input
func (obj *layer) Input() string {
	return obj.input
}
