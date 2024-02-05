package sqllites

import (
	"steve.care/network/domain/orms/schemas"
	"steve.care/network/domain/orms/schemas/roots"
	"steve.care/network/domain/orms/schemas/roots/groups"
	"steve.care/network/domain/orms/schemas/roots/groups/methods"
	group_methods "steve.care/network/domain/orms/schemas/roots/groups/methods"
	"steve.care/network/domain/orms/schemas/roots/groups/resources"
	"steve.care/network/domain/orms/schemas/roots/groups/resources/fields"
	field_methods "steve.care/network/domain/orms/schemas/roots/groups/resources/fields/methods"
	field_types "steve.care/network/domain/orms/schemas/roots/groups/resources/fields/types"
	"steve.care/network/domain/orms/schemas/roots/groups/resources/fields/types/dependencies"
	root_methods "steve.care/network/domain/orms/schemas/roots/methods"
)

type schemaFactory struct {
	builder                schemas.Builder
	rootsBuilder           roots.Builder
	rootBuilder            roots.RootBuilder
	rootMethodBuilder      root_methods.Builder
	groupBuilder           groups.Builder
	methodChainsBuilder    groups.MethodChainsBuilder
	methodChainBuilder     groups.MethodChainBuilder
	elementBuilder         groups.ElementBuilder
	resourceBuilder        resources.Builder
	groupMethodsBuilder    group_methods.Builder
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
	rootsBuilder roots.Builder,
	rootBuilder roots.RootBuilder,
	rootMethodBuilder root_methods.Builder,
	groupBuilder groups.Builder,
	methodChainsBuilder groups.MethodChainsBuilder,
	methodChainBuilder groups.MethodChainBuilder,
	elementBuilder groups.ElementBuilder,
	resourceBuilder resources.Builder,
	groupMethodsBuilder methods.Builder,
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
		rootsBuilder:           rootsBuilder,
		rootBuilder:            rootBuilder,
		rootMethodBuilder:      rootMethodBuilder,
		groupBuilder:           groupBuilder,
		methodChainsBuilder:    methodChainsBuilder,
		methodChainBuilder:     methodChainBuilder,
		elementBuilder:         elementBuilder,
		resourceBuilder:        resourceBuilder,
		groupMethodsBuilder:    groupMethodsBuilder,
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
		app.roots([]roots.Root{
			app.root(
				"tokens",
				app.rootMethods(
					"Create",
					"Now",
				),
				app.chains([]groups.MethodChain{
					app.chain(
						"IsDashboard",
						[]string{"Dashboard"},
						app.elementWithGroup(
							app.tokensDashboards(
								key,
							),
						),
					),
				}),
			),
		}),
	), nil
}

func (app *schemaFactory) tokensDashboards(
	key fields.Field,
) groups.Group {
	return app.group(
		"dashboards",
		app.groupMethods(
			"Create",
			"Now",
			"WithDashboard",
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
						app.groupMethods(
							"Create",
							"Now",
							"WithViewport",
						),
					),
				),
			),
			app.chain(
				"IsWidget",
				[]string{"Widget"},
				app.elementWithResource(
					app.resource(
						"widget",
						key,
						app.fields([]fields.Field{
							app.field(
								"title",
								app.fieldMethods(
									[]string{"Title"},
									"WithTitle",
								),
								app.fieldTypeWithKind(field_types.KindString),
								false,
							),
							app.field(
								"program",
								app.fieldMethods(
									[]string{"Program", "Bytes"},
									"WithProgram",
								),
								app.fieldTypeWithKind(field_types.KindBytes),
								false,
							),
							app.field(
								"input",
								app.fieldMethods(
									[]string{"Input"},
									"WithInput",
								),
								app.fieldTypeWithKind(field_types.KindBytes),
								false,
							),
							app.field(
								"viewport",
								app.fieldMethods(
									[]string{"Viewport", "Hash"},
									"WithViewport",
								),
								app.fieldTypeWithDependency(
									app.fieldDependency(
										"Viewport",
										[]string{"tokens", "dashboards"},
										"viewport",
										field_types.KindBytes,
									),
								),
								true,
							),
						}),
						app.groupMethods(
							"Create",
							"Now",
							"WithWidget",
						),
					),
				),
			),
		}),
	)
}

func (app *schemaFactory) schema(
	roots roots.Roots,
) schemas.Schema {
	ins, err := app.builder.Create().
		WithRoots(roots).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) roots(
	list []roots.Root,
) roots.Roots {
	ins, err := app.rootsBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) root(
	name string,
	methods root_methods.Methods,
	chains groups.MethodChains,
) roots.Root {
	ins, err := app.rootBuilder.Create().
		WithName(name).
		WithMethods(methods).
		WithChains(chains).
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

func (app *schemaFactory) groupMethods(
	initialize string,
	trigger string,
	element string,
) group_methods.Methods {
	ins, err := app.groupMethodsBuilder.Create().
		WithInitialize(initialize).
		WithTrigger(trigger).
		WithElement(element).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) rootMethods(
	initialize string,
	trigger string,
) root_methods.Methods {
	ins, err := app.rootMethodBuilder.Create().
		WithInitialize(initialize).
		WithTrigger(trigger).
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
	element string,
) field_methods.Methods {
	ins, err := app.fieldMethodsBuilder.Create().
		WithRetriever(retriever).
		WithElement(element).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) fieldDependency(
	retriever string,
	groups []string,
	resource string,
	kind uint8,
) dependencies.Dependency {
	ins, err := app.fieldDependencyBuilder.Create().
		WithRetriever(retriever).
		WithGroups(groups).
		WithResource(resource).
		WithKind(kind).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
