package tokens

import (
	"reflect"
	"testing"
	"time"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/dashboards"
	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/hash"
	token_dashboards "steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/dashboards"
	token_layers "steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/layers"
	token_links "steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/links"
	token_queries "steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/queries"
	token_receipts "steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/receipts"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
	"steve.care/network/domain/queries"
	"steve.care/network/domain/queries/conditions"
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/receipts/commands"
	"steve.care/network/domain/receipts/commands/results"
)

func TestToken_withLayer_Success(t *testing.T) {
	layer := token_layers.NewLayerWithLayerForTests(
		layers.NewLayerForTests(
			layers.NewInstructionsForTests([]layers.Instruction{
				layers.NewInstructionWithStopForTests(),
			}),
			layers.NewOutputForTests(
				"myVariable",
				layers.NewKindWithPromptForTests(),
			),
		),
	)

	ins := NewTokenWithLayerForTests(layer)

	now := time.Now().UTC()
	retCreatedOn := ins.CreatedOn()
	if now.Before(retCreatedOn) {
		t.Errorf("the time is invalid")
		return
	}

	content := ins.Content()
	if !content.IsLayer() {
		t.Errorf("the token was expected to contain a layer")
		return
	}

	if content.IsLink() {
		t.Errorf("the token was expected to NOT contain a link")
		return
	}

	if content.IsSuite() {
		t.Errorf("the token was expected to NOT contain a suite")
		return
	}

	if content.IsReceipt() {
		t.Errorf("the token was expected to NOT contain a receipt")
		return
	}

	if content.IsQuery() {
		t.Errorf("the token was expected to NOT contain a query")
		return
	}

	if content.IsDashboard() {
		t.Errorf("the token was expected to NOT contain a dashboard")
		return
	}

	retLayer := content.Layer()
	if !reflect.DeepEqual(layer, retLayer) {
		t.Errorf("the returned layer is invalid")
		return
	}
}

func TestToken_withLink_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))

	link := token_links.NewLinkWithLinkForTests(
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
	)

	ins := NewTokenWithLinkForTests(link)

	now := time.Now().UTC()
	retCreatedOn := ins.CreatedOn()
	if now.Before(retCreatedOn) {
		t.Errorf("the time is invalid")
		return
	}

	content := ins.Content()
	if content.IsLayer() {
		t.Errorf("the token was expected to NOT contain a layer")
		return
	}

	if !content.IsLink() {
		t.Errorf("the token was expected to contain a link")
		return
	}

	if content.IsSuite() {
		t.Errorf("the token was expected to NOT contain a suite")
		return
	}

	if content.IsReceipt() {
		t.Errorf("the token was expected to NOT contain a receipt")
		return
	}

	if content.IsQuery() {
		t.Errorf("the token was expected to NOT contain a query")
		return
	}

	if content.IsDashboard() {
		t.Errorf("the token was expected to NOT contain a dashboard")
		return
	}

	retLink := content.Link()
	if !reflect.DeepEqual(link, retLink) {
		t.Errorf("the returned link is invalid")
		return
	}
}

func TestToken_withReceipt_Success(t *testing.T) {
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

	receipt := token_receipts.NewReceiptWithReceiptForTests(
		receipts.NewReceiptForTests(commands, signature),
	)

	ins := NewTokenWithReceiptForTests(receipt)

	now := time.Now().UTC()
	retCreatedOn := ins.CreatedOn()
	if now.Before(retCreatedOn) {
		t.Errorf("the time is invalid")
		return
	}

	content := ins.Content()
	if content.IsLayer() {
		t.Errorf("the token was expected to NOT contain a layer")
		return
	}

	if content.IsLink() {
		t.Errorf("the token was expected to NOT contain a link")
		return
	}

	if content.IsSuite() {
		t.Errorf("the token was expected to NOT contain a suite")
		return
	}

	if !content.IsReceipt() {
		t.Errorf("the token was expected to contain a receipt")
		return
	}

	if content.IsQuery() {
		t.Errorf("the token was expected to NOT contain a query")
		return
	}

	if content.IsDashboard() {
		t.Errorf("the token was expected to NOT contain a dashboard")
		return
	}

	retReceipt := content.Receipt()
	if !reflect.DeepEqual(receipt, retReceipt) {
		t.Errorf("the returned receipt is invalid")
		return
	}
}

func TestToken_withQuery_Success(t *testing.T) {
	query := token_queries.NewQueryWithQueryForTests(
		queries.NewQueryForTests(
			"myEntity",
			conditions.NewConditionForTests(
				conditions.NewPointerForTests("myEntity", "myField"),
				conditions.NewOperatorWithEqualForTests(),
				conditions.NewElementWithResourceForTests(
					conditions.NewResourceWithValueForTests(45),
				),
			),
		),
	)

	ins := NewTokenWithQueryForTests(query)

	now := time.Now().UTC()
	retCreatedOn := ins.CreatedOn()
	if now.Before(retCreatedOn) {
		t.Errorf("the time is invalid")
		return
	}

	content := ins.Content()
	if content.IsLayer() {
		t.Errorf("the token was expected to NOT contain a layer")
		return
	}

	if content.IsLink() {
		t.Errorf("the token was expected to NOT contain a link")
		return
	}

	if content.IsSuite() {
		t.Errorf("the token was expected to NOT contain a suite")
		return
	}

	if content.IsReceipt() {
		t.Errorf("the token was expected to NOT contain a receipt")
		return
	}

	if !content.IsQuery() {
		t.Errorf("the token was expected to contain a query")
		return
	}

	if content.IsDashboard() {
		t.Errorf("the token was expected to NOT contain a dashboard")
		return
	}

	retQuery := content.Query()
	if !reflect.DeepEqual(query, retQuery) {
		t.Errorf("the returned query is invalid")
		return
	}
}

func TestToken_withDashboard_Success(t *testing.T) {
	pProgram, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	dashboard := token_dashboards.NewDashboardWithDashboardForTests(
		dashboards.NewDashboardForTests(
			"this is a title",
			widgets.NewWidgetsForTests([]widgets.Widget{
				widgets.NewWidgetForTests(
					"this is a title",
					*pProgram,
					[]byte("this is an input"),
				),
			}),
		),
	)

	ins := NewTokenWithDashboardForTests(dashboard)

	now := time.Now().UTC()
	retCreatedOn := ins.CreatedOn()
	if now.Before(retCreatedOn) {
		t.Errorf("the time is invalid")
		return
	}

	content := ins.Content()
	if content.IsLayer() {
		t.Errorf("the token was expected to NOT contain a layer")
		return
	}

	if content.IsLink() {
		t.Errorf("the token was expected to NOT contain a link")
		return
	}

	if content.IsSuite() {
		t.Errorf("the token was expected to NOT contain a suite")
		return
	}

	if content.IsReceipt() {
		t.Errorf("the token was expected to NOT contain a receipt")
		return
	}

	if content.IsQuery() {
		t.Errorf("the token was expected to NOT contain a query")
		return
	}

	if !content.IsDashboard() {
		t.Errorf("the token was expected to contain a dashboard")
		return
	}

	retDashboard := content.Dashboard()
	if !reflect.DeepEqual(dashboard, retDashboard) {
		t.Errorf("the returned dashboard is invalid")
		return
	}
}

func TestToken_withoutParam_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
