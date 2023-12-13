package receipts

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/receipts/commands"
)

type receipt struct {
	receipt receipts.Receipt
	command commands.Command
	link    commands.Link
}

func createReceiptWithReceipt(
	receiptIns receipts.Receipt,
) Receipt {
	return createReceiptInternally(
		receiptIns,
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
	)
}

func createReceiptWithLink(
	link commands.Link,
) Receipt {
	return createReceiptInternally(
		nil,
		nil,
		link,
	)
}

func createReceiptInternally(
	receiptIns receipts.Receipt,
	command commands.Command,
	link commands.Link,
) Receipt {
	out := receipt{
		receipt: receiptIns,
		command: command,
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

// IsLink returns true if there is a link, false otherwise
func (obj *receipt) IsLink() bool {
	return obj.link != nil
}

// Link returns the link, if any
func (obj *receipt) Link() commands.Link {
	return obj.link
}
