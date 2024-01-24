package suites

import (
	"steve.care/network/infrastructure/jsons/resources/tokens/layers"
	"steve.care/network/infrastructure/jsons/resources/tokens/links"
)

// Suite represents a suite
type Suite struct {
	Origin      links.Origin `json:"origin"`
	Input       []byte       `json:"input"`
	Expectation Expectation  `json:"expectation"`
}

// Expectation represents an expectation
type Expectation struct {
	Success *Output          `json:"success"`
	Mistake *links.Condition `json:"mistake"`
}

// Output represents an output
type Output struct {
	Kind  layers.Kind `json:"kind"`
	Value []byte      `json:"value"`
}
