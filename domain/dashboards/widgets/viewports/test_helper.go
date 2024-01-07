package viewports

// NewViewportForTests creates a new viewport for tests
func NewViewportForTests(row uint, height uint) Viewport {
	ins, err := NewBuilder().Create().WithRow(row).WithHeight(height).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
