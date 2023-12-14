package resources

import "steve.care/network/domain/resources/tokens"

// Resource represents a resource
type Resource struct {
	Token     tokens.Token `json:"token"`
	Signature string       `json:"signature"`
}
