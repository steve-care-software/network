package resources

import "steve.care/network/infrastructure/jsons/resources/tokens"

// Resource represents a resource
type Resource struct {
	Token     tokens.Token `json:"token"`
	Signature []byte       `json:"signature"`
}
