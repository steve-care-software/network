package links

import "steve.care/network/domain/hash"

type originResource struct {
	hash        hash.Hash
	layer       hash.Hash
	isMandatory bool
}

func createOriginResource(
	hash hash.Hash,
	layer hash.Hash,
	isMandatory bool,
) OriginResource {
	out := originResource{
		hash:        hash,
		layer:       layer,
		isMandatory: isMandatory,
	}

	return &out
}

// Hash returns the hash
func (obj *originResource) Hash() hash.Hash {
	return obj.hash
}

// Layer returns the layer
func (obj *originResource) Layer() hash.Hash {
	return obj.layer
}

// IsMandatory returns true if mandatory, false otherwise
func (obj *originResource) IsMandatory() bool {
	return obj.isMandatory
}
