package layers

import "steve.care/network/infrastructure/jsons/resources/tokens/links"

// Layer represents a layer
type Layer struct {
	Input        string        `json:"string"`
	Instructions []Instruction `json:"instructions"`
	Output       Output        `json:"output"`
}

// Output represents an outputs
type Output struct {
	Variable string `json:"variable"`
	Kind     Kind   `json:"kind"`
	Execute  string `json:"execute"`
}

// Kind represents a kind
type Kind struct {
	IsPrompt   bool `json:"is_prompt"`
	IsContinue bool `json:"is_continue"`
}

// Instruction represents the instruction
type Instruction struct {
	Stop       bool              `json:"stop"`
	RaiseError *uint             `json:"raise_error"`
	Condition  *Condition        `json:"condition"`
	Assignment *Assignment       `json:"assignment"`
	Link       *LinkInstruction  `json:"link"`
	Layer      *LayerInstruction `json:"layer"`
}

// LinkInstruction represents the link instruction
type LinkInstruction struct {
	Save   *links.Link `json:"save"`
	Delete string      `json:"delete"`
}

// LayerInstruction represents a layer instruction
type LayerInstruction struct {
	Save   *Layer `json:"save"`
	Delete string `json:"delete"`
}

// Condition represents the condition
type Condition struct {
	Variable     string        `json:"variable"`
	Instructions []Instruction `json:"instructions"`
}

// Assignment represents the assignment
type Assignment struct {
	Name       string     `json:"name"`
	Assignable Assignable `json:"assignable"`
}

// Assignable represents the assignable
type Assignable struct {
	Bytes    *Bytes    `json:"bytes"`
	Identity *Identity `json:"identity"`
}

// Bytes represents the bytes
type Bytes struct {
	Join    []string `json:"join"`
	Compare []string `json:"compare"`
	Hash    string   `json:"hash"`
}

// Identity represents the identity
type Identity struct {
	Signer    *Signer    `json:"signer"`
	Encryptor *Encryptor `json:"encryptor"`
}

// Encryptor represents the encryptor
type Encryptor struct {
	Decrypt     string `json:"decrypt"`
	Encrypt     string `json:"encrypt"`
	IsPublicKey *bool  `json:"is_public_key"`
}

// Signer represents a signer
type Signer struct {
	Sign             string           `json:"sign"`
	Vote             *Vote            `json:"vote"`
	GenSignerPubKeys uint             `json:"generate_signer_public_keys"`
	HashPublicKeys   string           `json:"hash_public_keys"`
	VoteVerify       *VoteVerify      `json:"vote_verify"`
	SignatureVerify  *SignatureVerify `json:"signature_verify"`
	Bytes            string           `json:"bytes"`
	IsPublicKey      bool             `json:"is_public_key"`
}

// SignatureVerify represents a signature verify
type SignatureVerify struct {
	Signature string `json:"signature"`
	Message   string `json:"message"`
}

// VoteVerify represents a vote verify
type VoteVerify struct {
	Vote       string `json:"vote"`
	Message    string `json:"message"`
	HashedRing string `json:"hashed_ring"`
}

// Vote represents a vote
type Vote struct {
	Ring    string `json:"ring"`
	Message string `json:"message"`
}
