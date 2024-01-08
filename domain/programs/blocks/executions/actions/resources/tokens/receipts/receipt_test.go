package receipts

import (
	"reflect"
	"testing"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/receipts/commands"
	"steve.care/network/domain/receipts/commands/results"
)

func TestReceipt_withReceipt_Success(t *testing.T) {
	commands := commands.NewCommandsForTests([]commands.Command{
		commands.NewCommandForTests(
			[]byte("this is the command input"),
			layers.NewLayerForTests(
				layers.NewInstructionsForTests([]layers.Instruction{
					layers.NewInstructionWithStopForTests(),
				}),
				layers.NewOutputForTests(
					"myVariable",
					layers.NewKindWithContinueForTests(),
				),
			),
			results.NewResultWithSuccessForTests(
				results.NewSuccessForTests(
					[]byte("this is some bytes"),
					layers.NewKindWithPromptForTests(),
				),
			),
		),
	})

	msg := commands.Hash().Bytes()
	signature, err := signers.NewFactory().Create().Sign(msg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	receipt := receipts.NewReceiptForTests(commands, signature)

	ins := NewReceiptWithReceiptForTests(receipt)

	if !ins.IsReceipt() {
		t.Errorf("the receipt was expected to contain a receipt")
		return
	}

	if ins.IsCommand() {
		t.Errorf("the receipt was expected to NOT contain a command")
		return
	}

	if ins.IsResult() {
		t.Errorf("the receipt was expected to NOT contain a result")
		return
	}

	if ins.IsSuccess() {
		t.Errorf("the receipt was expected to NOT contain a success")
		return
	}

	if ins.IsFailure() {
		t.Errorf("the receipt was expected to NOT contain a failure")
		return
	}

	if ins.IsLink() {
		t.Errorf("the receipt was expected to NOT contain a link")
		return
	}

	retReceipt := ins.Receipt()
	if !reflect.DeepEqual(receipt, retReceipt) {
		t.Errorf("the returned receipt is invalid")
		return
	}
}

func TestReceipt_withCommand_Success(t *testing.T) {
	command := commands.NewCommandForTests(
		[]byte("this is the command input"),
		layers.NewLayerForTests(
			layers.NewInstructionsForTests([]layers.Instruction{
				layers.NewInstructionWithStopForTests(),
			}),
			layers.NewOutputForTests(
				"myVariable",
				layers.NewKindWithContinueForTests(),
			),
		),
		results.NewResultWithSuccessForTests(
			results.NewSuccessForTests(
				[]byte("this is some bytes"),
				layers.NewKindWithPromptForTests(),
			),
		),
	)

	ins := NewReceiptWithCommandForTests(command)

	if ins.IsReceipt() {
		t.Errorf("the receipt was expected to NOT contain a receipt")
		return
	}

	if !ins.IsCommand() {
		t.Errorf("the receipt was expected to contain a command")
		return
	}

	if ins.IsResult() {
		t.Errorf("the receipt was expected to NOT contain a result")
		return
	}

	if ins.IsSuccess() {
		t.Errorf("the receipt was expected to NOT contain a success")
		return
	}

	if ins.IsFailure() {
		t.Errorf("the receipt was expected to NOT contain a failure")
		return
	}

	if ins.IsLink() {
		t.Errorf("the receipt was expected to NOT contain a link")
		return
	}

	retCommand := ins.Command()
	if !reflect.DeepEqual(command, retCommand) {
		t.Errorf("the returned command is invalid")
		return
	}
}

func TestReceipt_withResult_Success(t *testing.T) {
	result := results.NewResultWithSuccessForTests(
		results.NewSuccessForTests(
			[]byte("this is some bytes"),
			layers.NewKindWithPromptForTests(),
		),
	)

	ins := NewReceiptWithResultForTests(result)

	if ins.IsReceipt() {
		t.Errorf("the receipt was expected to NOT contain a receipt")
		return
	}

	if ins.IsCommand() {
		t.Errorf("the receipt was expected to NOT contain a command")
		return
	}

	if !ins.IsResult() {
		t.Errorf("the receipt was expected to contain a result")
		return
	}

	if ins.IsSuccess() {
		t.Errorf("the receipt was expected to NOT contain a success")
		return
	}

	if ins.IsFailure() {
		t.Errorf("the receipt was expected to NOT contain a failure")
		return
	}

	if ins.IsLink() {
		t.Errorf("the receipt was expected to NOT contain a link")
		return
	}

	retResult := ins.Result()
	if !reflect.DeepEqual(result, retResult) {
		t.Errorf("the returned result is invalid")
		return
	}
}

func TestReceipt_withSuccess_Success(t *testing.T) {
	success := results.NewSuccessForTests(
		[]byte("this is some bytes"),
		layers.NewKindWithPromptForTests(),
	)

	ins := NewReceiptWithSuccessForTests(success)

	if ins.IsReceipt() {
		t.Errorf("the receipt was expected to NOT contain a receipt")
		return
	}

	if ins.IsCommand() {
		t.Errorf("the receipt was expected to NOT contain a command")
		return
	}

	if ins.IsResult() {
		t.Errorf("the receipt was expected to NOT contain a result")
		return
	}

	if !ins.IsSuccess() {
		t.Errorf("the receipt was expected to contain a success")
		return
	}

	if ins.IsFailure() {
		t.Errorf("the receipt was expected to NOT contain a failure")
		return
	}

	if ins.IsLink() {
		t.Errorf("the receipt was expected to NOT contain a link")
		return
	}

	retSuccess := ins.Success()
	if !reflect.DeepEqual(success, retSuccess) {
		t.Errorf("the returned success is invalid")
		return
	}
}

func TestReceipt_withFailure_Success(t *testing.T) {
	failure := results.NewFailureForTests(uint(58), false)
	ins := NewReceiptWithFailureForTests(failure)

	if ins.IsReceipt() {
		t.Errorf("the receipt was expected to NOT contain a receipt")
		return
	}

	if ins.IsCommand() {
		t.Errorf("the receipt was expected to NOT contain a command")
		return
	}

	if ins.IsResult() {
		t.Errorf("the receipt was expected to NOT contain a result")
		return
	}

	if ins.IsSuccess() {
		t.Errorf("the receipt was expected to NOT contain a success")
		return
	}

	if !ins.IsFailure() {
		t.Errorf("the receipt was expected to contain a failure")
		return
	}

	if ins.IsLink() {
		t.Errorf("the receipt was expected to NOT contain a link")
		return
	}

	retFailure := ins.Failure()
	if !reflect.DeepEqual(failure, retFailure) {
		t.Errorf("the returned failure is invalid")
		return
	}
}

func TestReceipt_withLink_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))

	link := commands.NewLinkForTests(
		[]byte("this is an input"),
		links.NewLinkForTests(
			links.NewOriginForTests(
				links.NewOriginResourceForTests(*pFirstLayer),
				links.NewOperatorWithAndForTests(),
				links.NewOriginValueWithResourceForTests(
					links.NewOriginResourceForTests(*pSecondLayer),
				),
			),
			links.NewElementsForTests([]links.Element{
				links.NewElementForTests(*pLayer),
			}),
		),
		commands.NewCommandForTests(
			[]byte("this is the command input"),
			layers.NewLayerForTests(
				layers.NewInstructionsForTests([]layers.Instruction{
					layers.NewInstructionWithStopForTests(),
				}),
				layers.NewOutputForTests(
					"myVariable",
					layers.NewKindWithContinueForTests(),
				),
			),
			results.NewResultWithSuccessForTests(
				results.NewSuccessForTests(
					[]byte("this is some bytes"),
					layers.NewKindWithPromptForTests(),
				),
			),
		),
	)

	ins := NewReceiptWithLinkForTests(link)

	if ins.IsReceipt() {
		t.Errorf("the receipt was expected to NOT contain a receipt")
		return
	}

	if ins.IsCommand() {
		t.Errorf("the receipt was expected to NOT contain a command")
		return
	}

	if ins.IsResult() {
		t.Errorf("the receipt was expected to NOT contain a result")
		return
	}

	if ins.IsSuccess() {
		t.Errorf("the receipt was expected to NOT contain a success")
		return
	}

	if ins.IsFailure() {
		t.Errorf("the receipt was expected to NOT contain a failure")
		return
	}

	if !ins.IsLink() {
		t.Errorf("the receipt was expected to contain a link")
		return
	}

	retLink := ins.Link()
	if !reflect.DeepEqual(link, retLink) {
		t.Errorf("the returned link is invalid")
		return
	}
}

func TestReceipt_withoutParam_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
