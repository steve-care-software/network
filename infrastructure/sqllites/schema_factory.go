package sqllites

import (
	"steve.care/network/domain/schemas"
	"steve.care/network/domain/schemas/groups"
	"steve.care/network/domain/schemas/groups/resources"
	"steve.care/network/domain/schemas/groups/resources/fields"
)

type schemaFactory struct {
	builder             schemas.Builder
	groupBuilder        groups.Builder
	methodChainsBuilder groups.MethodChainsBuilder
	methodChainBuilder  groups.MethodChainBuilder
	elementBuilder      groups.ElementBuilder
	resourceBuilder     resources.Builder
	connectionsBuilder  resources.ConnectionsBuilder
	connectionBuilder   resources.ConnectionBuilder
	pointerBuilder      resources.PointerBuilder
	fieldsBuilder       fields.Builder
	fieldBuilder        fields.FieldBuilder
	keyFieldName        string
	keyFieldMethodNames []string
}

func createSchemaFactory(
	builder schemas.Builder,
	groupBuilder groups.Builder,
	methodChainsBuilder groups.MethodChainsBuilder,
	methodChainBuilder groups.MethodChainBuilder,
	elementBuilder groups.ElementBuilder,
	resourceBuilder resources.Builder,
	connectionsBuilder resources.ConnectionsBuilder,
	connectionBuilder resources.ConnectionBuilder,
	pointerBuilder resources.PointerBuilder,
	fieldsBuilder fields.Builder,
	fieldBuilder fields.FieldBuilder,
	keyFieldName string,
	keyFieldMethodNames []string,
) schemas.Factory {
	out := schemaFactory{
		builder:             builder,
		groupBuilder:        groupBuilder,
		methodChainsBuilder: methodChainsBuilder,
		methodChainBuilder:  methodChainBuilder,
		elementBuilder:      elementBuilder,
		resourceBuilder:     resourceBuilder,
		connectionsBuilder:  connectionsBuilder,
		connectionBuilder:   connectionBuilder,
		pointerBuilder:      pointerBuilder,
		fieldsBuilder:       fieldsBuilder,
		fieldBuilder:        fieldBuilder,
		keyFieldName:        keyFieldName,
		keyFieldMethodNames: keyFieldMethodNames,
	}

	return &out
}

// Create creates a schema
func (app *schemaFactory) Create() (schemas.Schema, error) {
	key := app.field(
		app.keyFieldName,
		app.keyFieldMethodNames,
		fields.KindBytes,
		false,
	)

	return app.schema(
		app.group(
			"resources",
			app.chains([]groups.MethodChain{
				app.chain(
					"IsDashboard",
					"Dashboard",
					app.elementWithGroup(
						app.resourcesDashboards(
							key,
						),
					),
				),
			}),
		),
	), nil
}

func (app *schemaFactory) resourcesDashboards(
	key fields.Field,
) groups.Group {
	return app.group(
		"dashboards",
		app.chains([]groups.MethodChain{
			app.chain(
				"IsViewport",
				"Viewport",
				app.elementWithResource(
					app.resource(
						"viewport",
						key,
						app.fields([]fields.Field{
							app.field("row", []string{"Row"}, fields.KindInteger, false),
							app.field("height", []string{"Height"}, fields.KindInteger, false),
						}),
					),
				),
			),
		}),
	)
}

func (app *schemaFactory) schema(
	group groups.Group,
) schemas.Schema {
	ins, err := app.builder.Create().
		WithGroup(group).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) group(
	name string,
	chains groups.MethodChains,
) groups.Group {
	ins, err := app.groupBuilder.Create().
		WithName(name).
		WithChains(chains).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) chains(
	list []groups.MethodChain,
) groups.MethodChains {
	ins, err := app.methodChainsBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) chain(
	condition string,
	value string,
	element groups.Element,
) groups.MethodChain {
	ins, err := app.methodChainBuilder.Create().
		WithCondition(condition).
		WithValue(value).
		WithElement(element).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) elementWithResource(
	resource resources.Resource,
) groups.Element {
	ins, err := app.elementBuilder.Create().
		WithResource(resource).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) elementWithGroup(
	group groups.Group,
) groups.Element {
	ins, err := app.elementBuilder.Create().
		WithGroup(group).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app schemaFactory) resourceWithConnections(
	name string,
	key fields.Field,
	fields fields.Fields,
	connections resources.Connections,
) resources.Resource {
	ins, err := app.resourceBuilder.Create().
		WithName(name).
		WithKey(key).
		WithFields(fields).
		WithConnections(connections).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app schemaFactory) resource(
	name string,
	key fields.Field,
	fields fields.Fields,
) resources.Resource {
	ins, err := app.resourceBuilder.Create().
		WithName(name).
		WithKey(key).
		WithFields(fields).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) connections(
	list []resources.Connection,
) resources.Connections {
	ins, err := app.connectionsBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) connection(
	field string,
	reference resources.Pointer,
) resources.Connection {
	ins, err := app.connectionBuilder.Create().
		WithField(field).
		WithReference(reference).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) pointer(
	resource resources.Resource,
	field string,
) resources.Pointer {
	ins, err := app.pointerBuilder.Create().
		WithResource(resource).
		WithField(field).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) fields(
	list []fields.Field,
) fields.Fields {
	ins, err := app.fieldsBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) field(
	name string,
	methods []string,
	kind uint8,
	canBeNil bool,
) fields.Field {
	builder := app.fieldBuilder.Create().
		WithName(name).
		WithMethods(methods).
		WithKind(kind)

	if canBeNil {
		builder.CanBeNil()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
