package layers

import "steve.care/network/domain/hash"

type assignableResource struct {
	hash            hash.Hash
	compile         string
	decompile       string
	amountByQuery   string
	listByQuery     string
	retrieveByQuery string
	retrieveByHash  string
	isAmount        bool
}

func createAssignableResourceWithCompile(
	hash hash.Hash,
	compile string,
) AssignableResource {
	return createAssignableResourceInternally(
		hash,
		compile,
		"",
		"",
		"",
		"",
		"",
		false,
	)
}

func createAssignableResourceWithDecompile(
	hash hash.Hash,
	decompile string,
) AssignableResource {
	return createAssignableResourceInternally(
		hash,
		"",
		decompile,
		"",
		"",
		"",
		"",
		false,
	)
}

func createAssignableResourceWithAmountByQuery(
	hash hash.Hash,
	amountByQuery string,
) AssignableResource {
	return createAssignableResourceInternally(
		hash,
		"",
		"",
		amountByQuery,
		"",
		"",
		"",
		false,
	)
}

func createAssignableResourceWithListByQuery(
	hash hash.Hash,
	listByQuery string,
) AssignableResource {
	return createAssignableResourceInternally(
		hash,
		"",
		"",
		"",
		listByQuery,
		"",
		"",
		false,
	)
}

func createAssignableResourceWithRetrieveByQuery(
	hash hash.Hash,
	retrieveByQuery string,
) AssignableResource {
	return createAssignableResourceInternally(
		hash,
		"",
		"",
		"",
		"",
		retrieveByQuery,
		"",
		false,
	)
}

func createAssignableResourceWithRetrieveByHash(
	hash hash.Hash,
	retrieveByHash string,
) AssignableResource {
	return createAssignableResourceInternally(
		hash,
		"",
		"",
		"",
		"",
		"",
		retrieveByHash,
		false,
	)
}

func createAssignableResourceWithAmount(
	hash hash.Hash,
) AssignableResource {
	return createAssignableResourceInternally(
		hash,
		"",
		"",
		"",
		"",
		"",
		"",
		true,
	)
}

func createAssignableResourceInternally(
	hash hash.Hash,
	compile string,
	decompile string,
	amountByQuery string,
	listByQuery string,
	retrieveByQuery string,
	retrieveByHash string,
	isAmount bool,
) AssignableResource {
	out := assignableResource{
		hash:            hash,
		compile:         compile,
		decompile:       decompile,
		amountByQuery:   amountByQuery,
		listByQuery:     listByQuery,
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
	return obj.compile != ""
}

// Compile returns the compile, if any
func (obj *assignableResource) Compile() string {
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
	return obj.amountByQuery != ""
}

// AmountByQuery returns the amountByQuery, if any
func (obj *assignableResource) AmountByQuery() string {
	return obj.amountByQuery
}

// IsListByQuery returns true if listByQuery, false otherwise
func (obj *assignableResource) IsListByQuery() bool {
	return obj.listByQuery != ""
}

// ListByQuery returns the listByQuery, if any
func (obj *assignableResource) ListByQuery() string {
	return obj.listByQuery
}

// IsRetrieveByQuery returns true if retrieveByQuery, false otherwise
func (obj *assignableResource) IsRetrieveByQuery() bool {
	return obj.retrieveByQuery != ""
}

// RetrieveByQuery returns the retrieveByQuery, if any
func (obj *assignableResource) RetrieveByQuery() string {
	return obj.retrieveByQuery
}

// IsRetrieveByHash returns true if retrieveByHash, false otherwise
func (obj *assignableResource) IsRetrieveByHash() bool {
	return obj.retrieveByHash != ""
}

// RetrieveByHash returns the retrieveByHash, if any
func (obj *assignableResource) RetrieveByHash() string {
	return obj.retrieveByHash
}

// IsAmount returns true if amount, false otherwise
func (obj *assignableResource) IsAmount() bool {
	return obj.isAmount
}
