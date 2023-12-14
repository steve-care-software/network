package jsons

import "steve.care/network/domain/resources"

type resourceAdapter struct {
}

func createResourceAdapter() resources.Adapter {
	out := resourceAdapter{}
	return &out
}

// ToBytes converts a resource to bytes
func (app *resourceAdapter) ToBytes(ins resources.Resource) ([]byte, error) {
	return nil, nil
}

// ToInstance converts bytes to resource instance
func (app *resourceAdapter) ToInstance(bytes []byte) (resources.Resource, error) {
	return nil, nil
}
