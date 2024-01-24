package sqllites

import (
	"steve.care/network/domain/schemas"
	"steve.care/network/domain/schemas/groups"
	"steve.care/network/domain/schemas/groups/resources"
	"steve.care/network/domain/schemas/groups/resources/fields"
)

type schemaFactory struct {
	builder            schemas.Builder
	groupsBuilder      groups.Builder
	groupBuilder       groups.GroupBuilder
	elementsBuilder    groups.ElementsBuilder
	elementBuilder     groups.ElementBuilder
	resourcesBuilder   resources.Builder
	resourceBuilder    resources.ResourceBuilder
	connectionsBuilder resources.ConnectionsBuilder
	connectionBuilder  resources.ConnectionBuilder
	pointerBuilder     resources.PointerBuilder
	fieldsBuilder      fields.Builder
	fieldBuilder       fields.FieldBuilder
	keyFieldName       string
}

func createSchemaFactory(
	builder schemas.Builder,
	groupsBuilder groups.Builder,
	groupBuilder groups.GroupBuilder,
	elementsBuilder groups.ElementsBuilder,
	elementBuilder groups.ElementBuilder,
	resourcesBuilder resources.Builder,
	resourceBuilder resources.ResourceBuilder,
	connectionsBuilder resources.ConnectionsBuilder,
	connectionBuilder resources.ConnectionBuilder,
	pointerBuilder resources.PointerBuilder,
	fieldsBuilder fields.Builder,
	fieldBuilder fields.FieldBuilder,
	keyFieldName string,
) schemas.Factory {
	out := schemaFactory{
		builder:            builder,
		groupsBuilder:      groupsBuilder,
		groupBuilder:       groupBuilder,
		elementsBuilder:    elementsBuilder,
		elementBuilder:     elementBuilder,
		resourcesBuilder:   resourcesBuilder,
		resourceBuilder:    resourceBuilder,
		connectionsBuilder: connectionsBuilder,
		connectionBuilder:  connectionBuilder,
		pointerBuilder:     pointerBuilder,
		fieldsBuilder:      fieldsBuilder,
		fieldBuilder:       fieldBuilder,
		keyFieldName:       keyFieldName,
	}

	return &out
}

// Create creates a schema
func (app *schemaFactory) Create() (schemas.Schema, error) {
	return app.schema(
		app.groups([]groups.Group{
			app.identities(),
			app.group(
				"resources",
				app.elements([]groups.Element{
					app.elementWithGroups(
						app.groups([]groups.Group{
							app.resourcesDashboards(),
						}),
					),
				}),
			),
		}),
	), nil
}

func (app *schemaFactory) identities() groups.Group {
	return app.group(
		"identities",
		app.elements([]groups.Element{
			app.elementWithResources(
				app.resources([]resources.Resource{
					app.resource(
						"account",
						app.field(app.keyFieldName, fields.KindBytes, false),
						app.fields([]fields.Field{
							app.field("username", fields.KindString, false),
							app.field("encryptor", fields.KindBytes, false),
							app.field("signer", fields.KindBytes, false),
						}),
					),
				}),
			),
		}),
	)
}

func (app *schemaFactory) resourcesDashboards() groups.Group {
	return app.group(
		"dashboards",
		app.elements([]groups.Element{
			app.elementWithResources(
				app.resources([]resources.Resource{
					app.resource(
						"viewport",
						app.field(app.keyFieldName, fields.KindBytes, false),
						app.fields([]fields.Field{
							app.field("row", fields.KindInteger, false),
							app.field("height", fields.KindInteger, false),
						}),
					),
				}),
			),
		}),
	)
}

func (app *schemaFactory) schema(
	groups groups.Groups,
) schemas.Schema {
	ins, err := app.builder.Create().
		WithGroups(groups).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) groups(
	list []groups.Group,
) groups.Groups {
	ins, err := app.groupsBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) group(
	name string,
	elements groups.Elements,
) groups.Group {
	ins, err := app.groupBuilder.Create().
		WithName(name).
		WithElements(elements).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) elements(
	list []groups.Element,
) groups.Elements {
	ins, err := app.elementsBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) elementWithResources(
	resources resources.Resources,
) groups.Element {
	ins, err := app.elementBuilder.Create().
		WithResources(resources).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) elementWithGroups(
	groups groups.Groups,
) groups.Element {
	ins, err := app.elementBuilder.Create().
		WithGroups(groups).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *schemaFactory) resources(
	list []resources.Resource,
) resources.Resources {
	ins, err := app.resourcesBuilder.Create().
		WithList(list).
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
	from resources.Pointer,
	to resources.Pointer,
) resources.Connection {
	ins, err := app.connectionBuilder.Create().
		From(from).
		To(to).
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
	kind uint8,
	canBeNil bool,
) fields.Field {
	builder := app.fieldBuilder.Create().
		WithName(name).
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
