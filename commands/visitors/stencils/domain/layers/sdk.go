package layers

// Layer represents a layer
type Layer interface {
	Path() []string
	Input() string
	Instructions() Instructions
	Output() Output
}

// Output represents the output
type Output interface {
	Variable() string
	Kind() Kind
}

// Kind represents the output kind
type Kind interface {
	IsPrompt() bool
	IsContinue() bool
	HasExecute() bool
	Execute() string
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
}

// Instruction represents an instruction
type Instruction interface {
	IsStop() bool
	IsRaiseError() bool
	RaiseError() uint
	IsCondition() bool
	Condition() Condition
	IsSave() bool
	Save() Layer
	IsAssignment() bool
	Assignment() Assignment
}

// Condition represents a condition
type Condition interface {
	Variable() string
	Instructions() Instructions
}

// Assignment represents an assignment
type Assignment interface {
	Name() string
	Assignable() Assignable
}

// Assignable represents an assignable
type Assignable interface {
	IsBytes() bool
	Bytes() Bytes
	IsIdentity() bool
	Identity() Identity
}

// Identity represents the assignable identity
type Identity interface {
	IsSigner() bool
	Signer() Signer
	IsEncryptor() bool
	Encryptor() Encryptor
}

// Bytes represents the bytes assignable
type Bytes interface {
	IsJoin() bool
	Join() BytesReferences
	IsCompare() bool
	Compare() BytesReferences
}

// Encryptor represents encryptor
type Encryptor interface {
	IsDecrypt() bool
	Decrypt() BytesReference
	IsEncrypt() bool
	Encrypt() BytesReference
	IsPublicKey() bool
}

// Signer represents the signer identity assignable
type Signer interface {
	IsSign() bool
	Sign() BytesReference
	IsVote() bool
	Vote() Vote
	IsGenerateSignerPublicKeys() bool
	GenerateSignerPublicKeys() uint
	IsHashPublicKeys() bool
	HashPublicKeys() string
	IsVoteVerify() bool
	VoteVerify() VoteVerify
	IsSignatureVerify() bool
	SignatureVerify() SignatureVerify
	IsBytes() bool
	Bytes() string
	IsPublicKey() bool
}

// SignatureVerify represents a signature verify
type SignatureVerify interface {
	Signature() string
	Message() BytesReference
}

// VoteVerify represents a vote verify
type VoteVerify interface {
	Vote() string
	Message() BytesReference
	HashedRing() string
}

// Vote represents a vote
type Vote interface {
	Ring() string
	Message() BytesReference
}

// BytesReferences represents bytes values
type BytesReferences interface {
	List() []BytesReference
}

// BytesReference a bytes value
type BytesReference interface {
	IsVariable() bool
	Variable() string
	IsBytes() bool
	Bytes() []byte
}

// Repository represents a layer repository
type Repository interface {
	List(basePath []string) ([]string, error)
	Dir(basePath []string) ([]string, error)
	Exists(path []string) (bool, error)
	Retrieve(path []string) (Layer, error)
}

// Service represents a layer service
type Service interface {
	Insert(context uint, layer Layer, path []string) error
	Delete(context uint, path []string) error
}
