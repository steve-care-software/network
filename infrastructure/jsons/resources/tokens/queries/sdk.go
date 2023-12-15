package queries

// Query represents a query
type Query struct {
	Entity    string    `json:"entity"`
	Condition Condition `json:"condition"`
	Fields    []string  `json:"fields"`
}

// Condition represents a condition
type Condition struct {
	Pointer  Pointer  `json:"pointer"`
	Operator Operator `json:"operator"`
	Element  Element  `json:"element"`
}

// Pointer represents a pointer
type Pointer struct {
	Entity string `json:"entity"`
	Field  string `json:"field"`
}

// Element represents an element
type Element struct {
	Condition *Condition `json:"condition"`
	Resource  *Resource  `json:"resource"`
}

// Resource represents the resource
type Resource struct {
	Field *Pointer    `json:"field"`
	Value interface{} `json:"value"`
}

// Operator represents an operator
type Operator struct {
	Equal      bool                `json:"equal"`
	Relational *RelationalOperator `json:"relational"`
	Integer    *IntegerOperator    `json:"integer"`
}

// RelationalOperator represents a relational operator
type RelationalOperator struct {
	And bool `json:"and"`
	Or  bool `json:"or"`
}

// IntegerOperator represents an integer operator
type IntegerOperator struct {
	SmallerThan bool `json:"smaller_than"`
	BiggerThan  bool `json:"bigger_than"`
	Equal       bool `json:"equal"`
}
