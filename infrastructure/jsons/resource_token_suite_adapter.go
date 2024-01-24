package jsons

import (
	resources_suites "steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/suites"
	"steve.care/network/domain/programs/logics/suites"
	"steve.care/network/domain/programs/logics/suites/expectations"
	"steve.care/network/domain/programs/logics/suites/expectations/outputs"
	structs_tokens "steve.care/network/infrastructure/jsons/resources/tokens"
	structs_suites "steve.care/network/infrastructure/jsons/resources/tokens/suites"
)

type resourceTokenSuiteAdapter struct {
	layerAdapter       *resourceTokenLayerAdapter
	linkAdapter        *resourceTokenLinkAdapter
	builder            resources_suites.Builder
	suiteBuilder       suites.SuiteBuilder
	expectationBuilder expectations.Builder
	outputBuilder      outputs.Builder
}

func createResourceTokenSuiteAdapter(
	layerAdapter *resourceTokenLayerAdapter,
	linkAdapter *resourceTokenLinkAdapter,
	builder resources_suites.Builder,
	suiteBuilder suites.SuiteBuilder,
	expectationBuilder expectations.Builder,
	outputBuilder outputs.Builder,
) *resourceTokenSuiteAdapter {
	out := resourceTokenSuiteAdapter{
		layerAdapter:       layerAdapter,
		linkAdapter:        linkAdapter,
		builder:            builder,
		suiteBuilder:       suiteBuilder,
		expectationBuilder: expectationBuilder,
		outputBuilder:      outputBuilder,
	}

	return &out
}

func (app *resourceTokenSuiteAdapter) toStruct(ins resources_suites.Suite) structs_tokens.Suite {
	output := structs_tokens.Suite{}
	if ins.IsSuite() {
		suite := app.suiteToStruct(ins.Suite())
		output.Suite = &suite
	}

	if ins.IsExpectation() {
		expectation := app.expectationToStruct(ins.Expectation())
		output.Expectation = &expectation
	}

	return output
}

func (app *resourceTokenSuiteAdapter) toInstance(ins structs_tokens.Suite) (resources_suites.Suite, error) {
	builder := app.builder.Create()
	if ins.Suite != nil {
		suite, err := app.structToSuite(*ins.Suite)
		if err != nil {
			return nil, err
		}

		builder.WithSuite(suite)
	}

	if ins.Expectation != nil {
		expectation, err := app.structToExpectation(*ins.Expectation)
		if err != nil {
			return nil, err
		}

		builder.WithExpectation(expectation)
	}

	return builder.Now()
}

func (app *resourceTokenSuiteAdapter) suiteToStruct(
	ins suites.Suite,
) structs_suites.Suite {
	origin := app.linkAdapter.originToStruct(ins.Origin())
	expectation := app.expectationToStruct(ins.Expectation())
	return structs_suites.Suite{
		Origin:      origin,
		Input:       ins.Input(),
		Expectation: expectation,
	}
}

func (app *resourceTokenSuiteAdapter) structToSuite(
	ins structs_suites.Suite,
) (suites.Suite, error) {
	origin, err := app.linkAdapter.structToOrigin(ins.Origin)
	if err != nil {
		return nil, err
	}

	expectation, err := app.structToExpectation(ins.Expectation)
	if err != nil {
		return nil, err
	}

	return app.suiteBuilder.Create().
		WithOrigin(origin).
		WithInput(ins.Input).
		WithExpectation(expectation).
		Now()
}

func (app *resourceTokenSuiteAdapter) expectationToStruct(
	ins expectations.Expectation,
) structs_suites.Expectation {
	output := structs_suites.Expectation{}
	if ins.IsSuccess() {
		success := app.outputToStruct(ins.Success())
		output.Success = &success
	}

	if ins.IsMistake() {
		mistake := app.linkAdapter.conditionToStruct(ins.Mistake())
		output.Mistake = &mistake
	}

	return output
}

func (app *resourceTokenSuiteAdapter) structToExpectation(
	ins structs_suites.Expectation,
) (expectations.Expectation, error) {
	builder := app.expectationBuilder.Create()
	if ins.Success != nil {
		output, err := app.structToOutput(*ins.Success)
		if err != nil {
			return nil, err
		}

		builder.WithSuccess(output)
	}

	if ins.Mistake != nil {
		condition, err := app.linkAdapter.structToCondition(*ins.Mistake)
		if err != nil {
			return nil, err
		}

		builder.WithMistake(condition)
	}

	return builder.Now()
}

func (app *resourceTokenSuiteAdapter) outputToStruct(
	ins outputs.Output,
) structs_suites.Output {
	kind := app.layerAdapter.kindToStruct(ins.Kind())
	return structs_suites.Output{
		Kind:  kind,
		Value: ins.Value(),
	}
}

func (app *resourceTokenSuiteAdapter) structToOutput(
	ins structs_suites.Output,
) (outputs.Output, error) {
	kind, err := app.layerAdapter.structToKind(ins.Kind)
	if err != nil {
		return nil, err
	}

	return app.outputBuilder.Create().
		WithKind(kind).
		WithValue(ins.Value).
		Now()
}
