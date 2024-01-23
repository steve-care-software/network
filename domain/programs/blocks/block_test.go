package blocks

import (
	"bytes"
	"math/big"
	"reflect"
	"testing"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions"
	"steve.care/network/domain/programs/blocks/transactions/executions"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	token_layers "steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

func TestBlock_Success(t *testing.T) {
	token := tokens.NewTokenWithLayerForTests(
		token_layers.NewLayerWithLayerForTests(
			layers.NewLayerForTests(
				layers.NewInstructionsForTests([]layers.Instruction{
					layers.NewInstructionWithStopForTests(),
				}),
				layers.NewOutputForTests(
					"myVariable",
					layers.NewKindWithPromptForTests(),
				),
				"myInput",
			),
		),
	)

	msg := token.Hash().Bytes()
	signature, err := signers.NewFactory().Create().Sign(msg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	executions := executions.NewExecutionsForTests([]executions.Execution{
		executions.NewExecutionForTests(
			actions.NewActionsForTests([]actions.Action{
				actions.NewActionWithCreateForTests(
					resources.NewResourceForTests(token, signature),
				),
			}),
		),
	})

	executionsMsg := executions.Hash().Bytes()
	execSignature, err := signers.NewFactory().Create().Sign(executionsMsg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	content := NewContentForTests(transactions.NewTransactionsForTests([]transactions.Transaction{
		transactions.NewTransactionForTests(executions, execSignature),
	}))

	var pMinedHash *hash.Hash
	prefix := []byte{0, 0}
	number := big.NewInt(0)
	for {
		pHash, err := Compute(content.Hash().Bytes(), number.Bytes())
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}

		retBytes := pHash.Bytes()
		if bytes.HasPrefix(retBytes, prefix) {
			pMinedHash = pHash
			break
		}

		// increment:
		number = number.Add(number, big.NewInt(1))
	}

	result := number.Bytes()
	ins := NewBlockForTests(content, result)

	retContent := ins.Content()
	if !reflect.DeepEqual(content, retContent) {
		t.Errorf("the returned content is invalid")
		return
	}

	retResult := ins.Result()
	if !reflect.DeepEqual(result, retResult) {
		t.Errorf("the returned result is invalid")
		return
	}

	retMinedHash := ins.MinedHash()
	if !reflect.DeepEqual(*pMinedHash, retMinedHash) {
		t.Errorf("the returned mined hash is invalid")
		return
	}

	retDifficulty := ins.Difficulty()
	expectedDifficulty := uint(len(prefix))
	if retDifficulty != expectedDifficulty {
		t.Errorf("the difficulty was expected to be %d, %d returned", expectedDifficulty, retDifficulty)
		return
	}
}

func TestBlock_withoutContent_returnsError(t *testing.T) {
	result := []byte("this is some result")
	_, err := NewBuilder().Create().WithResult(result).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBlock_withoutResult_returnsError(t *testing.T) {
	token := tokens.NewTokenWithLayerForTests(
		token_layers.NewLayerWithLayerForTests(
			layers.NewLayerForTests(
				layers.NewInstructionsForTests([]layers.Instruction{
					layers.NewInstructionWithStopForTests(),
				}),
				layers.NewOutputForTests(
					"myVariable",
					layers.NewKindWithPromptForTests(),
				),
				"myInput",
			),
		),
	)

	msg := token.Hash().Bytes()
	signature, err := signers.NewFactory().Create().Sign(msg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	executions := executions.NewExecutionsForTests([]executions.Execution{
		executions.NewExecutionForTests(
			actions.NewActionsForTests([]actions.Action{
				actions.NewActionWithCreateForTests(
					resources.NewResourceForTests(token, signature),
				),
			}),
		),
	})

	executionsMsg := executions.Hash().Bytes()
	execSignature, err := signers.NewFactory().Create().Sign(executionsMsg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	content := NewContentForTests(transactions.NewTransactionsForTests([]transactions.Transaction{
		transactions.NewTransactionForTests(executions, execSignature),
	}))

	_, err = NewBuilder().Create().WithContent(content).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBlock_withEmptyResult_returnsError(t *testing.T) {
	token := tokens.NewTokenWithLayerForTests(
		token_layers.NewLayerWithLayerForTests(
			layers.NewLayerForTests(
				layers.NewInstructionsForTests([]layers.Instruction{
					layers.NewInstructionWithStopForTests(),
				}),
				layers.NewOutputForTests(
					"myVariable",
					layers.NewKindWithPromptForTests(),
				),
				"myInput",
			),
		),
	)

	msg := token.Hash().Bytes()
	signature, err := signers.NewFactory().Create().Sign(msg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	executions := executions.NewExecutionsForTests([]executions.Execution{
		executions.NewExecutionForTests(
			actions.NewActionsForTests([]actions.Action{
				actions.NewActionWithCreateForTests(
					resources.NewResourceForTests(token, signature),
				),
			}),
		),
	})

	executionsMsg := executions.Hash().Bytes()
	execSignature, err := signers.NewFactory().Create().Sign(executionsMsg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	content := NewContentForTests(transactions.NewTransactionsForTests([]transactions.Transaction{
		transactions.NewTransactionForTests(executions, execSignature),
	}))

	_, err = NewBuilder().Create().WithContent(content).WithResult([]byte{}).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
