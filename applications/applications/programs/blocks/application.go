package blocks

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"steve.care/network/domain/accounts"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks"
	"steve.care/network/domain/programs/blocks/transactions"
)

type application struct {
	hashAdapter            hash.Adapter
	blockRepository        blocks.Repository
	blockService           blocks.Service
	blockBuilder           blocks.Builder
	blockContentBuilder    blocks.ContentBuilder
	transactionsRepository transactions.Repository
	transactionsService    transactions.Service
	transactionsBuilder    transactions.Builder
	authenticated          accounts.Account
	difficulty             blocks.Difficulty
}

func createApplication(
	hashAdapter hash.Adapter,
	blockRepository blocks.Repository,
	blockService blocks.Service,
	blockBuilder blocks.Builder,
	blockContentBuilder blocks.ContentBuilder,
	transactionsRepository transactions.Repository,
	transactionsService transactions.Service,
	transactionsBuilder transactions.Builder,
	authenticated accounts.Account,
	difficulty blocks.Difficulty,
) Application {
	out := application{
		hashAdapter:            hashAdapter,
		blockRepository:        blockRepository,
		blockService:           blockService,
		blockBuilder:           blockBuilder,
		blockContentBuilder:    blockContentBuilder,
		transactionsRepository: transactionsRepository,
		transactionsService:    transactionsService,
		transactionsBuilder:    transactionsBuilder,
		authenticated:          authenticated,
		difficulty:             difficulty,
	}

	return &out
}

// Exists returns true if the block exists, false otherwise
func (app *application) Exists(hash hash.Hash) bool {
	_, err := app.Retrieve(hash)
	return err != nil
}

// Retrieve retrieves a block
func (app *application) Retrieve(hash hash.Hash) (blocks.Block, error) {
	return app.blockRepository.Retrieve(hash)
}

// RetrieveHeadByProgram retrieves a head block by program hash
func (app *application) RetrieveHeadByProgram(program hash.Hash) (blocks.Block, error) {
	return app.blockRepository.RetrieveHeadByProgram(program)
}

// Transact adds a transaction
func (app *application) Transact(trx transactions.Transactions) error {
	list := trx.List()
	for _, oneTrx := range list {
		err := app.transactionsService.Insert(oneTrx)
		if err != nil {
			return err
		}
	}

	return nil
}

// Queue returns the queue
func (app *application) Queue() (transactions.Transactions, error) {
	trxList := []transactions.Transaction{}
	hashList, err := app.transactionsRepository.Queue()
	if err != nil {
		return nil, err
	}

	for _, oneHash := range hashList {
		trx, err := app.transactionsRepository.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		trxList = append(trxList, trx)
	}

	return app.transactionsBuilder.Create().
		WithList(trxList).
		Now()
}

// Mine mines a block for a program
func (app *application) Mine(program hash.Hash) error {
	// fetch the head block:
	headBlock, err := app.blockRepository.RetrieveHeadByProgram(program)
	if err != nil {
		return err
	}

	// fetch the trx queue:
	queue, err := app.Queue()
	if err != nil {
		return err
	}

	// build the block content:
	parent := headBlock.Hash()
	blockContent, err := app.blockContentBuilder.Create().
		WithTransactions(queue).
		WithParent(parent).
		Now()

	if err != nil {
		return err
	}

	// build the hash prefix:
	prevTrxAmount := uint(len(headBlock.Content().Transactions().List()))
	_, prefix, err := app.difficulty.Fetch(prevTrxAmount)
	if err != nil {
		return err
	}

	// create the message from the content hash:
	currentValue := big.NewInt(0)
	var minedHash hash.Hash

	for {

		// compute the hash:
		pHash, err := blocks.Compute(
			blockContent.Hash().Bytes(),
			currentValue.Bytes(),
		)

		if err != nil {
			return err
		}

		if bytes.HasPrefix(pHash.Bytes(), prefix) {
			minedHash = *pHash
			break
		}

		currentValue = currentValue.Add(currentValue, big.NewInt(1))
	}

	// build the block:
	newBlock, err := app.blockBuilder.Create().
		WithContent(blockContent).
		Now()

	if err != nil {
		return err
	}

	blockMinedHash := newBlock.MinedHash()
	if !bytes.Equal(blockMinedHash.Bytes(), minedHash.Bytes()) {
		str := fmt.Sprintf("there was an error was computing the mined hash while creating the block, the mining value was expected to be %s, %s returned", minedHash.String(), blockMinedHash.String())
		return errors.New(str)
	}

	// insert the block:
	return app.Insert(newBlock)
}

// Rewind rewinds a block
func (app *application) Rewind(head hash.Hash) error {
	// retrieve the block:
	block, err := app.Retrieve(head)
	if err != nil {
		return err
	}

	content := block.Content()
	if content.HasParent() {
		str := fmt.Sprintf("the provided block (hash: %s) cannot be rewinded because it does NOT contain a parent", block.Hash().String())
		return errors.New(str)
	}

	return app.blockService.Delete(block.Hash())
}

// Insert inserts a block
func (app *application) Insert(blocks blocks.Block) error {
	return app.blockService.Insert(blocks)
}
