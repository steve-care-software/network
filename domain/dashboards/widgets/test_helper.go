package widgets

import (
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/hash"
)

// NewWidgetsForTests creates a new widgets for tests
func NewWidgetsForTests(list []Widget) Widgets {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewWidgetWithViewportForTests creates a new widget with viewport for tests
func NewWidgetWithViewportForTests(title string, program hash.Hash, input []byte, viewport viewports.Viewport) Widget {
	ins, err := NewWidgetBuilder().Create().WithTitle(title).WithProgram(program).WithInput(input).WithViewport(viewport).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewWidgetForTests creates a new widget for tests
func NewWidgetForTests(title string, program hash.Hash, input []byte) Widget {
	ins, err := NewWidgetBuilder().Create().WithTitle(title).WithProgram(program).WithInput(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
