package sqllites

import (
	"steve.care/network/domain/schemas"
	"steve.care/network/domain/schemas/groups"
	group_methods "steve.care/network/domain/schemas/groups/methods"
	"steve.care/network/domain/schemas/groups/resources"
	"steve.care/network/domain/schemas/groups/resources/fields"
	field_methods "steve.care/network/domain/schemas/groups/resources/fields/methods"
	field_types "steve.care/network/domain/schemas/groups/resources/fields/types"
	"steve.care/network/domain/schemas/groups/resources/fields/types/dependencies"
	"steve.care/network/domain/schemas/groups/resources/methods"
	resource_methods "steve.care/network/domain/schemas/groups/resources/methods"
)

type schemaFactory struct {
	builder                schemas.Builder
	groupBuilder           groups.Builder
	methodChainsBuilder    groups.MethodChainsBuilder
	methodChainBuilder     groups.MethodChainBuilder
	methodBuilder          group_methods.Builder
	elementBuilder         groups.ElementBuilder
	resourceBuilder        resources.Builder
	builderMethodsBuilder  resource_methods.Builder
	connectionsBuilder     resources.ConnectionsBuilder
	connectionBuilder      resources.ConnectionBuilder
	pointerBuilder         resources.PointerBuilder
	fieldsBuilder          fields.Builder
	fieldBuilder           fields.FieldBuilder
	fieldMethodsBuilder    field_methods.Builder
	fieldTypeBuilder       field_types.Builder
	fieldDependencyBuilder dependencies.Builder
	keyFieldName           string
	keyFieldMethodNames    []string
}

func createSchemaFactory(
	builder schemas.Builder,
	groupBuilder groups.Builder,
	methodChainsBuilder groups.MethodChainsBuilder,
	methodChainBuilder groups.MethodChainBuilder,
	methodBuilder group_methods.Builder,
	elementBuilder groups.ElementBuilder,
	resourceBuilder resources.Builder,
	builderMethodsBuilder methods.Builder,
	connectionsBuilder resources.ConnectionsBuilder,
	connectionBuilder resources.ConnectionBuilder,
	pointerBuilder resources.PointerBuilder,
	fieldsBuilder fields.Builder,
	fieldBuilder fields.FieldBuilder,
	fieldMethodsBuilder field_methods.Builder,
	fieldTypeBuilder field_types.Builder,
	fieldDependencyBuilder dependencies.Builder,
	keyFieldName string,
	keyFieldMethodNames []string,
) schemas.Factory {
	out := schemaFactory{
		builder:                builder,
		groupBuilder:           groupBuilder,
		methodChainsBuilder:    methodChainsBuilder,
		methodChainBuilder:     methodChainBuilder,
		methodBuilder:          methodBuilder,
		elementBuilder:         elementBuilder,
		resourceBuilder:        resourceBuilder,
		builderMethodsBuilder:  builderMethodsBuilder,
		connectionsBuilder:     connectionsBuilder,
		connectionBuilder:      connectionBuilder,
		pointerBuilder:         pointerBuilder,
		fieldsBuilder:          fieldsBuilder,
		fieldBuilder:           fieldBuilder,
		fieldMethodsBuilder:    fieldMethodsBuilder,
		fieldTypeBuilder:       fieldTypeBuilder,
		fieldDependencyBuilder: fieldDependencyBuilder,
		keyFieldName:           keyFieldName,
		keyFieldMethodNames:    keyFieldMethodNames,
	}

	return &out
}

// Create creates a schema
func (app *schemaFactory) Create() (schemas.Schema, error) {
	key := app.field(
		app.keyFieldName,
		app.fieldMethods(
			app.keyFieldMethodNames,
			"Now",
		),
		app.fieldTypeWithKind(field_types.KindBytes),
		false,
	)

	return app.schema(
		app.group(
			"resources",
			app.groupMethods(
				app.builderMethods(
					"Create",
					"Now",
					app.fieldMethods(
						[]string{"Dashboard"},
						"WithDashboard",
					),
				),
			),
			app.chains([]groups.MethodChain{
				app.chain(
					"IsDashboard",
					[]string{"Dashboard"},
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
		app.groupMethods(
			app.builderMethods(
				"Create",
				"Now",
				app.fieldMethods(
					[]string{"Viewport"},
					"WithViewport",
				),
			),
		),
		app.chains([]groups.MethodChain{
			app.chain(
				"IsViewport",
				[]string{"Viewport"},
				app.elementWithResource(
					app.resource(
						"viewport",
						key,
						app.fields([]fields.Field{
							app.field(
								"row",
								app.fieldMethods(
									[]string{"Row"},
									"WithRow",
								),
								app.fieldTypeWithKind(field_types.KindInteger),
								false,
							),
							app.field(
								"height",
								app.fieldMethods(
									[]string{"Height"},
									"WithHeight",
								),
								app.fieldTypeWithKind(field_types.KindInteger),
								false,
							),
						}),
						app.builderMethods(
							"Create",
							"Now",
							app.fieldMethods(
								[]string{"Viewport"},
								"WithViewport",
							),
						),
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
	methods group_methods.Methods,
	chains groups.MethodChains,
) groups.Group {
	ins, err := app.groupBuilder.Create().
		WithName(name).
		WithMethods(methods).
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
	retriever []string,
	element groups.Element,
) groups.MethodChain {
	ins, err := app.methodChainBuilder.Create().
		WithCondition(condition).
		WithRetriever(retriever).
		WithElement(element).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) groupMethods(
	builder resource_methods.Methods,
) group_methods.Methods {
	ins, err := app.methodBuilder.Create().
		WithBuilder(builder).
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
	builder methods.Methods,
) resources.Resource {
	ins, err := app.resourceBuilder.Create().
		WithName(name).
		WithKey(key).
		WithFields(fields).
		WithBuilder(builder).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) builderMethods(
	initialize string,
	trigger string,
	field field_methods.Methods,
) resource_methods.Methods {
	ins, err := app.builderMethodsBuilder.Create().
		WithInitialize(initialize).
		WithTrigger(trigger).
		WithField(field).
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
	methods field_methods.Methods,
	typ field_types.Type,
	canBeNil bool,
) fields.Field {
	builder := app.fieldBuilder.Create().
		WithName(name).
		WithMethods(methods).
		WithType(typ)

	if canBeNil {
		builder.CanBeNil()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) fieldTypeWithDependency(
	dependency dependencies.Dependency,
) field_types.Type {
	ins, err := app.fieldTypeBuilder.Create().
		WithDependency(dependency).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) fieldTypeWithKind(
	kind uint8,
) field_types.Type {
	ins, err := app.fieldTypeBuilder.Create().
		WithKind(kind).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) fieldMethods(
	retriever []string,
	builder string,
) field_methods.Methods {
	ins, err := app.fieldMethodsBuilder.Create().
		WithRetriever(retriever).
		WithBuilder(builder).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) fieldDependency(
	groups []string,
	resource string,
) dependencies.Dependency {
	ins, err := app.fieldDependencyBuilder.Create().
		WithGroups(groups).
		WithResource(resource).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
