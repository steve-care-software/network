package tokens

import (
	"time"

	"steve.care/network/infrastructure/jsons/resources/tokens/layers"
	"steve.care/network/infrastructure/jsons/resources/tokens/links"
	"steve.care/network/infrastructure/jsons/resources/tokens/queries"
	"steve.care/network/infrastructure/jsons/resources/tokens/receipts"
	"steve.care/network/infrastructure/jsons/resources/tokens/suites"
)

// Token represents a token
type Token struct {
	Content   Content   `json:"content"`
	CreatedOn time.Time `json:"created_on"`
}

// Content represents the token content
type Content struct {
	Layer   *Layer   `json:"layer"`
	Link    *Link    `json:"link"`
	Query   *Query   `json:"query"`
	Receipt *Receipt `json:"receipt"`
	Suite   *Suite   `json:"suite"`
}

// Layer represents a layer
type Layer struct {
	Layer            *layers.Layer            `json:"layer"`
	Output           *layers.Output           `json:"output"`
	Kind             *layers.Kind             `json:"kind"`
	Instruction      *layers.Instruction      `json:"instruction"`
	LinkInstruction  *layers.LinkInstruction  `json:"link_instruction"`
	LayerInstruction *layers.LayerInstruction `json:"layer_instruction"`
	Condition        *layers.Condition        `json:"condition"`
	Assignment       *layers.Assignment       `json:"assignment"`
	Assignable       *layers.Assignable       `json:"assignable"`
	Bytes            *layers.Bytes            `json:"bytes"`
	Identity         *layers.Identity         `json:"identity"`
	Encryptor        *layers.Encryptor        `json:"encryptor"`
	Signer           *layers.Signer           `json:"signer"`
	SignatureVerify  *layers.SignatureVerify  `json:"signature_verify"`
	VoteVerify       *layers.VoteVerify       `json:"vote_verify"`
	Vote             *layers.Vote             `json:"vote"`
}

// Link represents a link
type Link struct {
	Link              *links.Link              `json:"link"`
	Element           *links.Element           `json:"element"`
	Condition         *links.Condition         `json:"condition"`
	ConditionValue    *links.ConditionValue    `json:"condition_value"`
	ConditionResource *links.ConditionResource `json:"condition_resource"`
	Origin            *links.Origin            `json:"origin"`
	OriginValue       *links.OriginValue       `json:"origin_value"`
	OriginResource    *links.OriginResource    `json:"origin_resource"`
	Operator          *links.Operator          `json:"operator"`
}

// Query represents a query
type Query struct {
	Query              *queries.Query              `json:"query"`
	Condition          *queries.Condition          `json:"condition"`
	Pointer            *queries.Pointer            `json:"pointer"`
	Element            *queries.Element            `json:"element"`
	Resource           *queries.Resource           `json:"resource"`
	Operator           *queries.Operator           `json:"operator"`
	RelationalOperator *queries.RelationalOperator `json:"relational_operator"`
	IntegerOperator    *queries.IntegerOperator    `json:"integer_operator"`
}

// Receipt represents a receipt
type Receipt struct {
	Receipt *receipts.Receipt `json:"receipt"`
	Command *receipts.Command `json:"command"`
	Result  *receipts.Result  `json:"result"`
	Success *receipts.Success `json:"success"`
	Failure *receipts.Failure `json:"failure"`
	Link    *receipts.Link    `json:"link"`
}

// Suite represents a suite
type Suite struct {
	Suite       *suites.Suite       `json:"suite"`
	Expectation *suites.Expectation `json:"expectation"`
}
