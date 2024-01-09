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
	Output    *layers.Layer    `json:"output"`
	Condition *links.Condition `json:"condition"`
}
