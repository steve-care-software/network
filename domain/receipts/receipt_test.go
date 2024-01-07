package receipts

import (
	"reflect"
	"testing"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/receipts/commands"
	"steve.care/network/domain/receipts/commands/results"
)

func TestReceipt_Success(t *testing.T) {
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

	ins := NewReceiptForTests(commands, signature)

	retCommands := ins.Commands()
	if !reflect.DeepEqual(commands, retCommands) {
		t.Errorf("the returned commands is invalid")
		return
	}

	retSignature := ins.Signature()
	if !reflect.DeepEqual(signature, retSignature) {
		t.Errorf("the returned signature is invalid")
		return
	}
}

func TestReceipt_withoutCommands_returnsError(t *testing.T) {
	msg := []byte("this is some msg")
	signature, err := signers.NewFactory().Create().Sign(msg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	_, err = NewBuilder().Create().WithSignature(signature).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestReceipt_withoutSignature_returnsError(t *testing.T) {
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

	_, err := NewBuilder().Create().WithCommands(commands).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
