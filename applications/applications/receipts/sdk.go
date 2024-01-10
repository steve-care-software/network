package receipts

import (
	"steve.care/network/domain/programs/blocks/transactions/executions"
	"steve.care/network/domain/receipts"
)

// Application represents the receipts application
type Application interface {
	Convert(receipt receipts.Receipt) (executions.Execution, error)
}
