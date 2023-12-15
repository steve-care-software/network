package receipts

import (
	"steve.care/network/infrastructure/jsons/resources/tokens/layers"
	"steve.care/network/infrastructure/jsons/resources/tokens/links"
)

// Receipt represents a receipt
type Receipt struct {
	Commands  []Command `json:"commands"`
	Signature []byte    `json:"string"`
}

// Command represents the command
type Command struct {
	Input  []byte       `json:"input"`
	Layer  layers.Layer `json:"layer"`
	Result Result       `json:"result"`
	Parent *Link        `json:"parent"`
}

// Result represents the result
type Result struct {
	Success *Success `json:"success"`
	Failure *Failure `json:"failure"`
}

// Success represents success
type Success struct {
	Bytes []byte      `json:"bytes"`
	Kind  layers.Kind `json:"kind"`
}

// Failure represents a failure
type Failure struct {
	Code          uint  `json:"code"`
	RaisedInLayer bool  `json:"raised_in_layer"`
	Index         *uint `json:"index"`
}

// Link represents a link
type Link struct {
	Input   []byte     `json:"input"`
	Link    links.Link `json:"link"`
	Command Command    `json:"command"`
}
