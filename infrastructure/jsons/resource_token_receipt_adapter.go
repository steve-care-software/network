package jsons

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/receipts/commands"
	"steve.care/network/domain/receipts/commands/results"
	resources_receipts "steve.care/network/domain/resources/tokens/receipts"
	structs_tokens "steve.care/network/infrastructure/jsons/resources/tokens"
	structs_receipts "steve.care/network/infrastructure/jsons/resources/tokens/receipts"
)

type resourceTokenReceiptAdapter struct {
	layerAdapter     *resourceTokenLayerAdapter
	linkAdapter      *resourceTokenLinkAdapter
	signatureAdapter signers.SignatureAdapter
	builder          resources_receipts.Builder
	receiptBuilder   receipts.Builder
	commandsBuilder  commands.Builder
	commandBuilder   commands.CommandBuilder
	resultBuilder    results.Builder
	failureBuilder   results.FailureBuilder
	successBuilder   results.SuccessBuilder
	linkBuilder      commands.LinkBuilder
}

func createResourceTokenReceiptAdapter(
	layerAdapter *resourceTokenLayerAdapter,
	linkAdapter *resourceTokenLinkAdapter,
	signatureAdapter signers.SignatureAdapter,
	builder resources_receipts.Builder,
	receiptBuilder receipts.Builder,
	commandsBuilder commands.Builder,
	commandBuilder commands.CommandBuilder,
	resultBuilder results.Builder,
	failureBuilder results.FailureBuilder,
	successBuilder results.SuccessBuilder,
	linkBuilder commands.LinkBuilder,
) *resourceTokenReceiptAdapter {
	out := resourceTokenReceiptAdapter{
		layerAdapter:     layerAdapter,
		linkAdapter:      linkAdapter,
		signatureAdapter: signatureAdapter,
		builder:          builder,
		receiptBuilder:   receiptBuilder,
		commandsBuilder:  commandsBuilder,
		commandBuilder:   commandBuilder,
		resultBuilder:    resultBuilder,
		failureBuilder:   failureBuilder,
		successBuilder:   successBuilder,
		linkBuilder:      linkBuilder,
	}

	return &out
}

func (app *resourceTokenReceiptAdapter) toStruct(ins resources_receipts.Receipt) (*structs_tokens.Receipt, error) {
	output := structs_tokens.Receipt{}
	if ins.IsReceipt() {
		pReceipt, err := app.receiptToStruct(ins.Receipt())
		if err != nil {
			return nil, err
		}

		output.Receipt = pReceipt
	}

	if ins.IsCommand() {
		command := app.commandToStruct(ins.Command())
		output.Command = &command
	}

	if ins.IsResult() {
		result := app.resultToStruct(ins.Result())
		output.Result = &result
	}

	if ins.IsSuccess() {
		success := app.successToStruct(ins.Success())
		output.Success = &success
	}

	if ins.IsFailure() {
		failure := app.failureToStruct(ins.Failure())
		output.Failure = &failure
	}

	if ins.IsLink() {
		link := app.linkToStruct(ins.Link())
		output.Link = &link
	}

	return &output, nil
}

func (app *resourceTokenReceiptAdapter) toInstance(ins structs_tokens.Receipt) (resources_receipts.Receipt, error) {
	builder := app.builder.Create()
	if ins.Receipt != nil {
		receipt, err := app.structToReceipt(*ins.Receipt)
		if err != nil {
			return nil, err
		}

		builder.WithReceipt(receipt)
	}

	if ins.Command != nil {
		command, err := app.structToCommand(*ins.Command)
		if err != nil {
			return nil, err
		}

		builder.WithCommand(command)
	}

	if ins.Result != nil {
		result, err := app.structToResult(*ins.Result)
		if err != nil {
			return nil, err
		}

		builder.WithResult(result)
	}

	if ins.Success != nil {
		success, err := app.structToSuccess(*ins.Success)
		if err != nil {
			return nil, err
		}

		builder.WithSuccess(success)
	}

	if ins.Failure != nil {
		failure, err := app.structToFailure(*ins.Failure)
		if err != nil {
			return nil, err
		}

		builder.WithFailure(failure)
	}

	if ins.Link != nil {
		link, err := app.structToLink(*ins.Link)
		if err != nil {
			return nil, err
		}

		builder.WithLink(link)
	}

	return builder.Now()
}

func (app *resourceTokenReceiptAdapter) receiptToStruct(
	ins receipts.Receipt,
) (*structs_receipts.Receipt, error) {
	commands := app.commandsToStructs(ins.Commands())
	sigBytes, err := ins.Signature().Bytes()
	if err != nil {
		return nil, err
	}

	return &structs_receipts.Receipt{
		Commands:  commands,
		Signature: sigBytes,
	}, nil
}

func (app *resourceTokenReceiptAdapter) structToReceipt(
	ins structs_receipts.Receipt,
) (receipts.Receipt, error) {
	commands, err := app.structsToCommands(ins.Commands)
	if err != nil {
		return nil, err
	}

	sig, err := app.signatureAdapter.ToSignature(ins.Signature)
	if err != nil {
		return nil, err
	}

	return app.receiptBuilder.Create().
		WithCommands(commands).
		WithSignature(sig).
		Now()
}

func (app *resourceTokenReceiptAdapter) commandsToStructs(
	ins commands.Commands,
) []structs_receipts.Command {
	list := ins.List()
	output := []structs_receipts.Command{}
	for _, oneCommand := range list {
		ins := app.commandToStruct(oneCommand)
		output = append(output, ins)
	}

	return output
}

func (app *resourceTokenReceiptAdapter) structsToCommands(
	list []structs_receipts.Command,
) (commands.Commands, error) {
	output := []commands.Command{}
	for _, oneStruct := range list {
		ins, err := app.structToCommand(oneStruct)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.commandsBuilder.Create().
		WithList(output).
		Now()
}

func (app *resourceTokenReceiptAdapter) commandToStruct(
	ins commands.Command,
) structs_receipts.Command {
	layer := app.layerAdapter.layerToStruct(ins.Layer())
	result := app.resultToStruct(ins.Result())
	output := structs_receipts.Command{
		Input:  ins.Input(),
		Layer:  layer,
		Result: result,
	}

	if ins.HasParent() {
		parent := app.linkToStruct(ins.Parent())
		output.Parent = &parent
	}

	return output
}

func (app *resourceTokenReceiptAdapter) structToCommand(
	ins structs_receipts.Command,
) (commands.Command, error) {
	layer, err := app.layerAdapter.structToLayer(ins.Layer)
	if err != nil {
		return nil, err
	}

	result, err := app.structToResult(ins.Result)
	if err != nil {
		return nil, err
	}

	builder := app.commandBuilder.Create().
		WithInput(ins.Input).
		WithLayer(layer).
		WithResult(result)

	if ins.Parent != nil {
		parent, err := app.structToLink(*ins.Parent)
		if err != nil {
			return nil, err
		}

		builder.WithParent(parent)
	}

	return builder.Now()
}

func (app *resourceTokenReceiptAdapter) resultToStruct(
	ins results.Result,
) structs_receipts.Result {
	output := structs_receipts.Result{}
	if ins.IsSuccess() {
		success := app.successToStruct(ins.Success())
		output.Success = &success
	}

	if ins.IsFailure() {
		failure := app.failureToStruct(ins.Failure())
		output.Failure = &failure
	}

	return output
}

func (app *resourceTokenReceiptAdapter) structToResult(
	ins structs_receipts.Result,
) (results.Result, error) {
	builder := app.resultBuilder.Create()
	if ins.Success != nil {
		success, err := app.structToSuccess(*ins.Success)
		if err != nil {
			return nil, err
		}

		builder.WithSuccess(success)
	}

	if ins.Failure != nil {
		failure, err := app.structToFailure(*ins.Failure)
		if err != nil {
			return nil, err
		}

		builder.WithFailure(failure)
	}

	return builder.Now()
}

func (app *resourceTokenReceiptAdapter) successToStruct(
	ins results.Success,
) structs_receipts.Success {
	kind := app.layerAdapter.kindToStruct(ins.Kind())
	output := structs_receipts.Success{
		Bytes: ins.Bytes(),
		Kind:  kind,
	}

	return output
}

func (app *resourceTokenReceiptAdapter) structToSuccess(
	ins structs_receipts.Success,
) (results.Success, error) {
	kind, err := app.layerAdapter.structToKind(ins.Kind)
	if err != nil {
		return nil, err
	}

	return app.successBuilder.Create().
		WithBytes(ins.Bytes).
		WithKind(kind).
		Now()
}

func (app *resourceTokenReceiptAdapter) failureToStruct(
	ins results.Failure,
) structs_receipts.Failure {
	output := structs_receipts.Failure{
		Code:          ins.Code(),
		RaisedInLayer: ins.IsRaisedInLayer(),
	}

	if ins.HasIndex() {
		output.Index = ins.Index()
	}

	return output
}

func (app *resourceTokenReceiptAdapter) structToFailure(
	ins structs_receipts.Failure,
) (results.Failure, error) {
	builder := app.failureBuilder.Create().
		WithCode(ins.Code)

	if ins.RaisedInLayer {
		builder.IsRaisedInLayer()
	}

	if ins.Index != nil {
		builder.WithIndex(*ins.Index)
	}

	return builder.Now()
}

func (app *resourceTokenReceiptAdapter) linkToStruct(
	ins commands.Link,
) structs_receipts.Link {
	link := app.linkAdapter.linkToStruct(ins.Link())
	command := app.commandToStruct(ins.Command())
	output := structs_receipts.Link{
		Input:   ins.Input(),
		Link:    link,
		Command: command,
	}

	return output
}

func (app *resourceTokenReceiptAdapter) structToLink(
	ins structs_receipts.Link,
) (commands.Link, error) {
	link, err := app.linkAdapter.structToLink(ins.Link)
	if err != nil {
		return nil, err
	}

	command, err := app.structToCommand(ins.Command)
	if err != nil {
		return nil, err
	}

	return app.linkBuilder.Create().
		WithInput(ins.Input).
		WithLink(link).
		WithCommand(command).
		Now()
}
