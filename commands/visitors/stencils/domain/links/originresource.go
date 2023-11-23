package links

import "steve.care/network/libraries/hash"

type originResource struct {
	hash        hash.Hash
	container   []string
	isMandatory bool
}

func createOriginResource(
	hash hash.Hash,
	container []string,
	isMandatory bool,
) OriginResource {
	out := originResource{
		hash:        hash,
		container:   container,
		isMandatory: isMandatory,
	}

	return &out
}

// Hash returns the hash
func (obj *originResource) Hash() hash.Hash {
	return obj.hash
}

// Container returns the container
func (obj *originResource) Container() []string {
	return obj.container
}

// IsMandatory returns true if mandatory, false otherwise
func (obj *originResource) IsMandatory() bool {
	return obj.isMandatory
}
