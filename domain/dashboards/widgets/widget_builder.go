package widgets

import (
	"errors"

	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/hash"
)

type widgetBuilder struct {
	hashAdapter hash.Adapter
	title       string
	program     hash.Hash
	input       []byte
	viewport    viewports.Viewport
}

func createWidgetBuilder(
	hashAdapter hash.Adapter,
) WidgetBuilder {
	out := widgetBuilder{
		hashAdapter: hashAdapter,
		title:       "",
		program:     nil,
		input:       nil,
		viewport:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *widgetBuilder) Create() WidgetBuilder {
	return createWidgetBuilder(
		app.hashAdapter,
	)
}

// WithTitle adds a title to the builder
func (app *widgetBuilder) WithTitle(title string) WidgetBuilder {
	app.title = title
	return app
}

// WithProgram adds a program to the builder
func (app *widgetBuilder) WithProgram(program hash.Hash) WidgetBuilder {
	app.program = program
	return app
}

// WithInput adds an input to the builder
func (app *widgetBuilder) WithInput(input []byte) WidgetBuilder {
	app.input = input
	return app
}

// WithViewport adds a viewport to the builder
func (app *widgetBuilder) WithViewport(viewport viewports.Viewport) WidgetBuilder {
	app.viewport = viewport
	return app
}

// Now builds a new Widget instance
func (app *widgetBuilder) Now() (Widget, error) {
	if app.title == "" {
		return nil, errors.New("the title is mandatory in order to build a Widget instance")
	}

	if app.program == nil {
		return nil, errors.New("the program is mandatory in order to build a Widget instance")
	}

	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.input == nil {
		return nil, errors.New("the input is mandatory in order to build a Widget instance")
	}

	data := [][]byte{
		[]byte(app.title),
		app.program.Bytes(),
		app.input,
	}

	if app.viewport != nil {
		data = append(data, app.viewport.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.viewport != nil {
		return createWidgetWithViewport(*pHash, app.title, app.program, app.input, app.viewport), nil
	}

	return createWidget(*pHash, app.title, app.program, app.input), nil
}
