package layers

import "steve.care/network/libraries/hash"

type layer struct {
	hash         hash.Hash
	path         []string
	input        string
	instructions Instructions
	output       Output
}

func createLayer(
	hash hash.Hash,
	path []string,
	input string,
	instructions Instructions,
	output Output,
) Layer {
	out := layer{
		hash:         hash,
		path:         path,
		input:        input,
		instructions: instructions,
		output:       output,
	}

	return &out
}

// Hash returns the hash
func (obj *layer) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *layer) Path() []string {
	return obj.path
}

// Input returns the input
func (obj *layer) Input() string {
	return obj.input
}

// Instructions returns the instructions
func (obj *layer) Instructions() Instructions {
	return obj.instructions
}

// Output returns the output
func (obj *layer) Output() Output {
	return obj.output
}
