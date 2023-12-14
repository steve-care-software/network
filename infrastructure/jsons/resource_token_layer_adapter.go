package jsons

import (
	"steve.care/network/domain/receipts/commands/layers"
	resources_layers "steve.care/network/domain/resources/tokens/layers"
	structs_layers "steve.care/network/infrastructure/jsons/resources/tokens/layers"
)

type resourceTokenLayerAdapter struct {
	bytesReferenceBuilder layers.BytesReferenceBuilder
}

// ToStruct converts a resource layer to struct
func (app *resourceTokenLayerAdapter) ToStruct(ins resources_layers.Layer) structs_layers.Layer {
	return structs_layers.Layer{}
}

// ToInstance converts bytes to resource layer instance
func (app *resourceTokenLayerAdapter) ToInstance(ins structs_layers.Layer) (resources_layers.Layer, error) {
	return nil, nil
}

func (app *resourceTokenLayerAdapter) bytesReferenceToStruct(
	ins layers.BytesReference,
) structs_layers.BytesReference {
	output := structs_layers.BytesReference{}
	if ins.IsVariable() {
		output.Variable = ins.Variable()
	}

	if ins.IsBytes() {
		output.Bytes = ins.Bytes()
	}

	return output
}

func (app *resourceTokenLayerAdapter) structToBytesReference(
	ins structs_layers.BytesReference,
) (layers.BytesReference, error) {
	builder := app.bytesReferenceBuilder.Create()
	if ins.Variable != "" {
		builder.WithVariable(ins.Variable)
	}

	if ins.Bytes != nil && len(ins.Bytes) > 0 {
		builder.WithBytes(ins.Bytes)
	}

	return builder.Now()
}
