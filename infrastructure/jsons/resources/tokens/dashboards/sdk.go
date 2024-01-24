package dashboards

// Dashboard represents a dashboard
type Dashboard struct {
	Title   string   `json:"title"`
	Widgets []Widget `json:"widgets"`
}

// Widget represents a widget
type Widget struct {
	Title    string    `json:"title"`
	Program  []byte    `json:"program"`
	Input    []byte    `json:"input"`
	Viewport *Viewport `json:"viewport"`
}

// Viewport represents a viewport
type Viewport struct {
	Row    uint `json:"row"`
	Height uint `json:"height"`
}
