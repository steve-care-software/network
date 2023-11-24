package receipts

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands"
)

type receipt struct {
	hash      hash.Hash
	commands  commands.Commands
	signature signers.Signature
}

func createReceipt(
	hash hash.Hash,
	commands commands.Commands,
	signature signers.Signature,
) Receipt {
	out := receipt{
		hash:      hash,
		commands:  commands,
		signature: signature,
	}

	return &out
}

// Hash returns the hash
func (obj *receipt) Hash() hash.Hash {
	return obj.hash
}

// Commands returns the commands
func (obj *receipt) Commands() commands.Commands {
	return obj.commands
}

// Signature returns the signature
func (obj *receipt) Signature() signers.Signature {
	return obj.signature
}
