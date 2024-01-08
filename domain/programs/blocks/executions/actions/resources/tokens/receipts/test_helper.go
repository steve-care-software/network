package receipts

import (
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/receipts/commands"
	"steve.care/network/domain/receipts/commands/results"
)

// NewReceiptWithReceiptForTests creates a new receipt with receipt for tests
func NewReceiptWithReceiptForTests(input receipts.Receipt) Receipt {
	ins, err := NewBuilder().Create().WithReceipt(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewReceiptWithCommandForTests creates a new receipt with command for tests
func NewReceiptWithCommandForTests(input commands.Command) Receipt {
	ins, err := NewBuilder().Create().WithCommand(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewReceiptWithResultForTests creates a new receipt with result for tests
func NewReceiptWithResultForTests(input results.Result) Receipt {
	ins, err := NewBuilder().Create().WithResult(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewReceiptWithSuccessForTests creates a new receipt with success for tests
func NewReceiptWithSuccessForTests(input results.Success) Receipt {
	ins, err := NewBuilder().Create().WithSuccess(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewReceiptWithFailureForTests creates a new receipt with failure for tests
func NewReceiptWithFailureForTests(input results.Failure) Receipt {
	ins, err := NewBuilder().Create().WithFailure(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewReceiptWithLinkForTests creates a new receipt with link for tests
func NewReceiptWithLinkForTests(input commands.Link) Receipt {
	ins, err := NewBuilder().Create().WithLink(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
