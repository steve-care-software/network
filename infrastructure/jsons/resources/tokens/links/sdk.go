package links

// Link represents the link
type Link struct {
	Origin   Origin    `json:"origin"`
	Elements []Element `json:"elements"`
}

// Element represents an element
type Element struct {
	Layer     string     `json:"layer"`
	Condition *Condition `json:"condition"`
}

// Condition represents a condition
type Condition struct {
	Resource ConditionResource `json:"resource"`
	Operator Operator          `json:"operator"`
	Next     ConditionValue    `json:"next"`
}

// ConditionValue represents a condition value
type ConditionValue struct {
	Resource  *ConditionResource `json:"resource"`
	Condition *Condition         `json:"condition"`
}

// ConditionResource represents the condition resource
type ConditionResource struct {
	Code            uint `json:"code"`
	IsRaisedInLayer bool `json:"is_raised_in_layer"`
}

// Origin represents the origin
type Origin struct {
	Resource OriginResource `json:"resource"`
	Operator Operator       `json:"operator"`
	Next     OriginValue    `json:"next"`
}

// OriginValue represents the origin value
type OriginValue struct {
	Resource *OriginResource `json:"resource"`
	Origin   *Origin         `json:"origin"`
}

// OriginResource represents the origin resource
type OriginResource struct {
	Layer       string `json:"layer"`
	IsMandatory bool   `json:"is_mandatory"`
}

// Operator represents the operator
type Operator struct {
	And bool `json:"and"`
	Or  bool `json:"or"`
	Xor bool `json:"xor"`
}
