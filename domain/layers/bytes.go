package layers

import "steve.care/network/domain/hash"

type bytes struct {
	hash      hash.Hash
	join      BytesReferences
	compare   BytesReferences
	hashBytes BytesReference
}

func createBytesWithJoin(
	hash hash.Hash,
	join BytesReferences,
) Bytes {
	return createBytesInternally(hash, join, nil, nil)
}

func createBytesWithCompare(
	hash hash.Hash,
	compare BytesReferences,
) Bytes {
	return createBytesInternally(hash, nil, compare, nil)
}

func createBytesWithHashBytes(
	hash hash.Hash,
	hashBytes BytesReference,
) Bytes {
	return createBytesInternally(hash, nil, nil, hashBytes)
}

func createBytesInternally(
	hash hash.Hash,
	join BytesReferences,
	compare BytesReferences,
	hashBytes BytesReference,
) Bytes {
	out := bytes{
		hash:      hash,
		join:      join,
		compare:   compare,
		hashBytes: hashBytes,
	}

	return &out
}

// Hash returns the hash
func (obj *bytes) Hash() hash.Hash {
	return obj.hash
}

// IsJoin returns true if there is a join, false otherwise
func (obj *bytes) IsJoin() bool {
	return obj.join != nil
}

// Join returns the join, if any
func (obj *bytes) Join() BytesReferences {
	return obj.join
}

// IsCompare returns true if there is a compare, false otherwise
func (obj *bytes) IsCompare() bool {
	return obj.compare != nil
}

// Compare returns the compare, if any
func (obj *bytes) Compare() BytesReferences {
	return obj.compare
}

// IsHashBytes returns true if there is a hashBytes, false otherwise
func (obj *bytes) IsHashBytes() bool {
	return obj.hashBytes != nil
}

// HashBytes returns the hashBytes, if any
func (obj *bytes) HashBytes() BytesReference {
	return obj.hashBytes
}
