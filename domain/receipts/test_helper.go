package receipts

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/receipts/commands"
)

// NewReceiptForTests creates a new receipt for tests
func NewReceiptForTests(commands commands.Commands, sig signers.Signature) Receipt {
	ins, err := NewBuilder().Create().WithCommands(commands).WithSignature(sig).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
