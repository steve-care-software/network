package edwards25519

import (
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"steve.care/network/applications/applications/encryptors"
)

var curve = edwards25519.NewBlakeSHA256Ed25519()

// NewApplication creates a new application
func NewApplication() encryptors.Application {
	return createApplication()
}
