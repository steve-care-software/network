package jsons

import (
	"encoding/json"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	structs "steve.care/network/infrastructure/jsons/resources"
)

type resourceAdapter struct {
	tokenAdapter     *resourceTokenAdapter
	signatureAdapter signers.SignatureAdapter
	builder          resources.Builder
}

// ToBytes converts a resource to bytes
func (app *resourceAdapter) ToBytes(ins resources.Resource) ([]byte, error) {
	pToken, err := app.tokenAdapter.toStruct(ins.Token())
	if err != nil {
		return nil, err
	}

	signatureBytes, err := ins.Signature().Bytes()
	if err != nil {
		return nil, err
	}

	structIns := structs.Resource{
		Token:     *pToken,
		Signature: signatureBytes,
	}

	return json.Marshal(structIns)
}

// ToInstance converts bytes to resource instance
func (app *resourceAdapter) ToInstance(bytes []byte) (resources.Resource, error) {
	structIns := structs.Resource{}
	err := json.Unmarshal(bytes, &structIns)
	if err != nil {
		return nil, err
	}

	token, err := app.tokenAdapter.toInstance(structIns.Token)
	if err != nil {
		return nil, err
	}

	signature, err := app.signatureAdapter.ToSignature(structIns.Signature)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithToken(token).
		WithSignature(signature).
		Now()
}
