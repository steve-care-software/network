package jsons

import (
	resources_suites "steve.care/network/domain/programs/blockchains/blocks/executions/actions/resources/tokens/suites"
	"steve.care/network/domain/programs/logics/suites"
	"steve.care/network/domain/programs/logics/suites/expectations"
	structs_tokens "steve.care/network/infrastructure/jsons/resources/tokens"
	structs_suites "steve.care/network/infrastructure/jsons/resources/tokens/suites"
)

type resourceTokenSuiteAdapter struct {
	layerAdapter       *resourceTokenLayerAdapter
	linkAdapter        *resourceTokenLinkAdapter
	builder            resources_suites.Builder
	suiteBuilder       suites.Builder
	expectationBuilder expectations.Builder
}

func createResourceTokenSuiteAdapter(
	layerAdapter *resourceTokenLayerAdapter,
	linkAdapter *resourceTokenLinkAdapter,
	builder resources_suites.Builder,
	suiteBuilder suites.Builder,
	expectationBuilder expectations.Builder,
) *resourceTokenSuiteAdapter {
	out := resourceTokenSuiteAdapter{
		layerAdapter:       layerAdapter,
		linkAdapter:        linkAdapter,
		builder:            builder,
		suiteBuilder:       suiteBuilder,
		expectationBuilder: expectationBuilder,
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
	input := app.layerAdapter.layerToStruct(ins.Input())
	expectation := app.expectationToStruct(ins.Expectation())
	return structs_suites.Suite{
		Origin:      origin,
		Input:       input,
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

	input, err := app.layerAdapter.structToLayer(ins.Input)
	if err != nil {
		return nil, err
	}

	expectation, err := app.structToExpectation(ins.Expectation)
	if err != nil {
		return nil, err
	}

	return app.suiteBuilder.Create().
		WithOrigin(origin).
		WithInput(input).
		WithExpectation(expectation).
		Now()
}

func (app *resourceTokenSuiteAdapter) expectationToStruct(
	ins expectations.Expectation,
) structs_suites.Expectation {
	output := structs_suites.Expectation{}
	if ins.IsOutput() {
		layer := app.layerAdapter.layerToStruct(ins.Output())
		output.Output = &layer
	}

	if ins.IsCondition() {
		condition := app.linkAdapter.conditionToStruct(ins.Condition())
		output.Condition = &condition
	}

	return output
}

func (app *resourceTokenSuiteAdapter) structToExpectation(
	ins structs_suites.Expectation,
) (expectations.Expectation, error) {
	builder := app.expectationBuilder.Create()
	if ins.Output != nil {
		output, err := app.layerAdapter.structToLayer(*ins.Output)
		if err != nil {
			return nil, err
		}

		builder.WithOutput(output)
	}

	if ins.Condition != nil {
		condition, err := app.linkAdapter.structToCondition(*ins.Condition)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	return builder.Now()
}
