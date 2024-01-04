package layers

import "steve.care/network/domain/hash"

type assignableResource struct {
	hash            hash.Hash
	compile         BytesReference
	decompile       string
	amountByQuery   BytesReference
	retrieveByQuery BytesReference
	retrieveByHash  BytesReference
	isAmount        bool
}

func createAssignableResourceWithCompile(
	hash hash.Hash,
	compile BytesReference,
) AssignableResource {
	return createAssignableResourceInternally(
		hash,
		compile,
		"",
		nil,
		nil,
		nil,
		false,
	)
}

func createAssignableResourceWithDecompile(
	hash hash.Hash,
	decompile string,
) AssignableResource {
	return createAssignableResourceInternally(
		hash,
		nil,
		decompile,
		nil,
		nil,
		nil,
		false,
	)
}

func createAssignableResourceWithAmountByQuery(
	hash hash.Hash,
	amountByQuery BytesReference,
) AssignableResource {
	return createAssignableResourceInternally(
		hash,
		nil,
		"",
		amountByQuery,
		nil,
		nil,
		false,
	)
}

func createAssignableResourceWithRetrieveByQuery(
	hash hash.Hash,
	retrieveByQuery BytesReference,
) AssignableResource {
	return createAssignableResourceInternally(
		hash,
		nil,
		"",
		nil,
		retrieveByQuery,
		nil,
		false,
	)
}

func createAssignableResourceWithRetrieveByHash(
	hash hash.Hash,
	retrieveByHash BytesReference,
) AssignableResource {
	return createAssignableResourceInternally(
		hash,
		nil,
		"",
		nil,
		nil,
		retrieveByHash,
		false,
	)
}

func createAssignableResourceWithAmount(
	hash hash.Hash,
) AssignableResource {
	return createAssignableResourceInternally(
		hash,
		nil,
		"",
		nil,
		nil,
		nil,
		true,
	)
}

func createAssignableResourceInternally(
	hash hash.Hash,
	compile BytesReference,
	decompile string,
	amountByQuery BytesReference,
	retrieveByQuery BytesReference,
	retrieveByHash BytesReference,
	isAmount bool,
) AssignableResource {
	out := assignableResource{
		hash:            hash,
		compile:         compile,
		decompile:       decompile,
		amountByQuery:   amountByQuery,
		retrieveByQuery: retrieveByQuery,
		retrieveByHash:  retrieveByHash,
		isAmount:        isAmount,
	}

	return &out
}

// Hash returns the hash
func (obj *assignableResource) Hash() hash.Hash {
	return obj.hash
}

// IsCompile returns true if compile, false otherwise
func (obj *assignableResource) IsCompile() bool {
	return obj.compile != nil
}

// Compile returns the compile, if any
func (obj *assignableResource) Compile() BytesReference {
	return obj.compile
}

// IsDecompile returns true if decompile, false otherwise
func (obj *assignableResource) IsDecompile() bool {
	return obj.decompile != ""
}

// Decompile returns the decompile, if any
func (obj *assignableResource) Decompile() string {
	return obj.decompile
}

// IsAmountByQuery returns true if amountByQuery, false otherwise
func (obj *assignableResource) IsAmountByQuery() bool {
	return obj.amountByQuery != nil
}

// AmountByQuery returns the amountByQuery, if any
func (obj *assignableResource) AmountByQuery() BytesReference {
	return obj.amountByQuery
}

// IsRetrieveByQuery returns true if retrieveByQuery, false otherwise
func (obj *assignableResource) IsRetrieveByQuery() bool {
	return obj.retrieveByQuery != nil
}

// RetrieveByQuery returns the retrieveByQuery, if any
func (obj *assignableResource) RetrieveByQuery() BytesReference {
	return obj.retrieveByQuery
}

// IsRetrieveByHash returns true if retrieveByHash, false otherwise
func (obj *assignableResource) IsRetrieveByHash() bool {
	return obj.retrieveByHash != nil
}

// RetrieveByHash returns the retrieveByHash, if any
func (obj *assignableResource) RetrieveByHash() BytesReference {
	return obj.retrieveByHash
}

// IsAmount returns true if amount, false otherwise
func (obj *assignableResource) IsAmount() bool {
	return obj.isAmount
}
