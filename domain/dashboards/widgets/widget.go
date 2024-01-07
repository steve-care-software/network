package widgets

import (
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/hash"
)

type widget struct {
	hash     hash.Hash
	title    string
	program  hash.Hash
	input    []byte
	viewport viewports.Viewport
}

func createWidget(
	hash hash.Hash,
	title string,
	program hash.Hash,
	input []byte,
) Widget {
	return createWidgetInternally(hash, title, program, input, nil)
}

func createWidgetWithViewport(
	hash hash.Hash,
	title string,
	program hash.Hash,
	input []byte,
	viewport viewports.Viewport,
) Widget {
	return createWidgetInternally(hash, title, program, input, viewport)
}

func createWidgetInternally(
	hash hash.Hash,
	title string,
	program hash.Hash,
	input []byte,
	viewport viewports.Viewport,
) Widget {
	out := widget{
		hash:     hash,
		title:    title,
		program:  program,
		input:    input,
		viewport: viewport,
	}

	return &out
}

// Hash returns the hash
func (obj *widget) Hash() hash.Hash {
	return obj.hash
}

// Title returns the title
func (obj *widget) Title() string {
	return obj.title
}

// Program returns the program
func (obj *widget) Program() hash.Hash {
	return obj.program
}

// Input returns the input
func (obj *widget) Input() []byte {
	return obj.input
}

// HasViewport returns true if there is a viewport, false otherwise
func (obj *widget) HasViewport() bool {
	return obj.viewport != nil
}

// Viewport returns the viewport, if any
func (obj *widget) Viewport() viewports.Viewport {
	return obj.viewport
}
