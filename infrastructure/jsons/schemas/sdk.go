package schemas

// Schema represents a schema
type Schema struct {
	Entities    []Entity     `json:"entities"`
	Connections []Connection `json:"connections"`
}

// Entity represents an entity
type Entity struct {
	Name   string  `json:"name"`
	Head   string  `json:"head"`
	Fields []Field `json:"fields"`
}

// Field represents a field
type Field struct {
	Name     string `json:"name"`
	Kind     Kind   `json:"kind"`
	IsUnique bool   `json:"is_unique"`
}

// Kind represents a field kind
type Kind struct {
	IsNil     bool `json:"is_nil"`
	IsInteger bool `json:"is_integer"`
	IsReal    bool `json:"is_real"`
	IsText    bool `json:"is_text"`
	IsBlob    bool `json:"is_blob"`
}

// Connection represents a connection
type Connection struct {
	From Pointer `json:"from"`
	To   Pointer `json:"to"`
}

// Pointer represents a pointer
type Pointer struct {
	Entity string `json:"entity"`
	Field  string `json:"field"`
}
