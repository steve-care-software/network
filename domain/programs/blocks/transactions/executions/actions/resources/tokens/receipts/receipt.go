package receipts

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/receipts/commands"
	"steve.care/network/domain/receipts/commands/results"
)

type receipt struct {
	receipt receipts.Receipt
	command commands.Command
	result  results.Result
	success results.Success
	failure results.Failure
	link    commands.Link
}

func createReceiptWithReceipt(
	receiptIns receipts.Receipt,
) Receipt {
	return createReceiptInternally(
		receiptIns,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createReceiptWithCommand(
	command commands.Command,
) Receipt {
	return createReceiptInternally(
		nil,
		command,
		nil,
		nil,
		nil,
		nil,
	)
}

func createReceiptWithResult(
	result results.Result,
) Receipt {
	return createReceiptInternally(
		nil,
		nil,
		result,
		nil,
		nil,
		nil,
	)
}

func createReceiptWithSuccess(
	success results.Success,
) Receipt {
	return createReceiptInternally(
		nil,
		nil,
		nil,
		success,
		nil,
		nil,
	)
}

func createReceiptWithFailure(
	failure results.Failure,
) Receipt {
	return createReceiptInternally(
		nil,
		nil,
		nil,
		nil,
		failure,
		nil,
	)
}

func createReceiptWithLink(
	link commands.Link,
) Receipt {
	return createReceiptInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		link,
	)
}

func createReceiptInternally(
	receiptIns receipts.Receipt,
	command commands.Command,
	result results.Result,
	success results.Success,
	failure results.Failure,
	link commands.Link,
) Receipt {
	out := receipt{
		receipt: receiptIns,
		command: command,
		result:  result,
		success: success,
		failure: failure,
		link:    link,
	}

	return &out
}

// Hash returns the hash
func (obj *receipt) Hash() hash.Hash {
	if obj.IsReceipt() {
		return obj.receipt.Hash()
	}

	if obj.IsCommand() {
		return obj.command.Hash()
	}

	if obj.IsSuccess() {
		return obj.success.Hash()
	}

	if obj.IsFailure() {
		return obj.failure.Hash()
	}

	return obj.link.Hash()
}

// IsReceipt returns true if there is a receipt, false otherwise
func (obj *receipt) IsReceipt() bool {
	return obj.receipt != nil
}

// Receipt returns the receipt, if any
func (obj *receipt) Receipt() receipts.Receipt {
	return obj.receipt
}

// IsCommand returns true if there is a command, false otherwise
func (obj *receipt) IsCommand() bool {
	return obj.command != nil
}

// Command returns the command, if any
func (obj *receipt) Command() commands.Command {
	return obj.command
}

// IsResult returns true if there is a result, false otherwise
func (obj *receipt) IsResult() bool {
	return obj.result != nil
}

// Result returns the result, if any
func (obj *receipt) Result() results.Result {
	return obj.result
}

// IsSuccess returns true if there is a success, false otherwise
func (obj *receipt) IsSuccess() bool {
	return obj.success != nil
}

// Success returns the success, if any
func (obj *receipt) Success() results.Success {
	return obj.success
}

// IsFailure returns true if there is a failure, false otherwise
func (obj *receipt) IsFailure() bool {
	return obj.failure != nil
}

// Failure returns the failure, if any
func (obj *receipt) Failure() results.Failure {
	return obj.failure
}

// IsLink returns true if there is a link, false otherwise
func (obj *receipt) IsLink() bool {
	return obj.link != nil
}

// Link returns the link, if any
func (obj *receipt) Link() commands.Link {
	return obj.link
}
