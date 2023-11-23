package layers

import (
	"steve.care/network/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewOutputBuilder creates a new output builder
func NewOutputBuilder() OutputBuilder {
	hashAdapter := hash.NewAdapter()
	return createOutputBuilder(
		hashAdapter,
	)
}

// NewKindBuilder creates a new kind builder
func NewKindBuilder() KindBuilder {
	hashAdapter := hash.NewAdapter()
	return createKindBuilder(
		hashAdapter,
	)
}

// NewInstructionsBuilder creates a new instructions builder
func NewInstructionsBuilder() InstructionsBuilder {
	hashAdapter := hash.NewAdapter()
	return createInstructionsBuilder(
		hashAdapter,
	)
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	hashAdapter := hash.NewAdapter()
	return createInstructionBuilder(
		hashAdapter,
	)
}

// NewConditionBuildercreates a new condition builder
func NewConditionBuilder() ConditionBuilder {
	hashAdapter := hash.NewAdapter()
	return createConditionBuilder(
		hashAdapter,
	)
}

// NewAssignmentBuilder creates a new assignment builder
func NewAssignmentBuilder() AssignmentBuilder {
	hashAdapter := hash.NewAdapter()
	return createAssignmentBuilder(
		hashAdapter,
	)
}

// NewAssignableBuilder creates a new assignable builder
func NewAssignableBuilder() AssignableBuilder {
	hashAdapter := hash.NewAdapter()
	return createAssignableBuilder(
		hashAdapter,
	)
}

// NewBytesBuilder creates a new bytes builder
func NewBytesBuilder() BytesBuilder {
	hashAdapter := hash.NewAdapter()
	return createBytesBuilder(
		hashAdapter,
	)
}

// NewIdentityBuilder creates a new identity builder
func NewIdentityBuilder() IdentityBuilder {
	hashAdapter := hash.NewAdapter()
	return createIdentityBuilder(
		hashAdapter,
	)
}

// NewEncryptorBuilder creates a new encryptor builder
func NewEncryptorBuilder() EncryptorBuilder {
	hashAdapter := hash.NewAdapter()
	return createEncryptorBuilder(
		hashAdapter,
	)
}

// NewSignerBuilder creates a new signer builder
func NewSignerBuilder() SignerBuilder {
	hashAdapter := hash.NewAdapter()
	return createSignerBuilder(
		hashAdapter,
	)
}

// NewSignatureVerifyBuilder creates a new signature verify builder
func NewSignatureVerifyBuilder() SignatureVerifyBuilder {
	hashAdapter := hash.NewAdapter()
	return createSignatureVerifyBuilder(
		hashAdapter,
	)
}

// NewVoteBuilder creates a new vote builder
func NewVoteBuilder() VoteBuilder {
	hashAdapter := hash.NewAdapter()
	return createVoteBuilder(
		hashAdapter,
	)
}

// NewVoteVerifyBuilder creates a new vote verify builder
func NewVoteVerifyBuilder() VoteVerifyBuilder {
	hashAdapter := hash.NewAdapter()
	return createVoteVerifyBuilder(
		hashAdapter,
	)
}

// NewBytesReferencesBuilder creates a new bytes references builder
func NewBytesReferencesBuilder() BytesReferencesBuilder {
	hashAdapter := hash.NewAdapter()
	return createBytesReferencesBuilder(
		hashAdapter,
	)
}

// NewBytesReferenceBuilder creates a new bytes reference builder
func NewBytesReferenceBuilder() BytesReferenceBuilder {
	hashAdapter := hash.NewAdapter()
	return createBytesReferenceBuilder(
		hashAdapter,
	)
}

// Builder represents a layer builder
type Builder interface {
	Create() Builder
	WithPath(path []string) Builder
	WithInput(input string) Builder
	WithInstructions(instructions Instructions) Builder
	WithOutput(output Output) Builder
	Now() (Layer, error)
}

// Layer represents a layer
type Layer interface {
	Hash() hash.Hash
	Path() []string
	Input() string
	Instructions() Instructions
	Output() Output
}

// OutputBuilder represents an output builder
type OutputBuilder interface {
	Create() OutputBuilder
	WithVariable(variable string) OutputBuilder
	WithKind(kind Kind) OutputBuilder
	WithExecute(execute string) OutputBuilder
	Now() (Output, error)
}

// Output represents the output
type Output interface {
	Hash() hash.Hash
	Variable() string
	Kind() Kind
	HasExecute() bool
	Execute() string
}

// KindBuilder represents a kind builder
type KindBuilder interface {
	Create() KindBuilder
	IsPrompt() KindBuilder
	IsContinue() KindBuilder
	Now() (Kind, error)
}

// Kind represents the output kind
type Kind interface {
	Hash() hash.Hash
	IsPrompt() bool
	IsContinue() bool
}

// InstructionsBuilder represents instructions builder
type InstructionsBuilder interface {
	Create() InstructionsBuilder
	WithList(list []Instruction) InstructionsBuilder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	Hash() hash.Hash
	List() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithRaiseError(raiseError uint) InstructionBuilder
	WithCondition(condition Condition) InstructionBuilder
	WthSave(save Layer) InstructionBuilder
	WithAssignment(assignment Assignment) InstructionBuilder
	IsStop() InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	Hash() hash.Hash
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

// ConditionBuilder represents a condition builder
type ConditionBuilder interface {
	Create() ConditionBuilder
	WithVariable(variable string) ConditionBuilder
	WithInstructions(instructions Instructions) ConditionBuilder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Variable() string
	Instructions() Instructions
}

// AssignmentBuilder represents an assignment builder
type AssignmentBuilder interface {
	Create() AssignmentBuilder
	WithName(name string) AssignmentBuilder
	WithAssignable(assignable Assignable) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Hash() hash.Hash
	Name() string
	Assignable() Assignable
}

// AssignableBuilder represents an assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithBytes(bytes Bytes) AssignableBuilder
	WithIdentity(identity Identity) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	Hash() hash.Hash
	IsBytes() bool
	Bytes() Bytes
	IsIdentity() bool
	Identity() Identity
}

// BytesBuilder represents a bytes builder
type BytesBuilder interface {
	Create() BytesBuilder
	WithJoin(join BytesReferences) BytesBuilder
	WithCompare(compare BytesReferences) BytesBuilder
	Now() (Bytes, error)
}

// Bytes represents the bytes assignable
type Bytes interface {
	Hash() hash.Hash
	IsJoin() bool
	Join() BytesReferences
	IsCompare() bool
	Compare() BytesReferences
}

// IdentityBuilder represents an identity builder
type IdentityBuilder interface {
	Create() IdentityBuilder
	WithSigner(signer Signer) IdentityBuilder
	WithEncryptor(encryptor Encryptor) IdentityBuilder
	Now() (Identity, error)
}

// Identity represents the assignable identity
type Identity interface {
	Hash() hash.Hash
	IsSigner() bool
	Signer() Signer
	IsEncryptor() bool
	Encryptor() Encryptor
}

// EncryptorBuilder represents an encryptor builder
type EncryptorBuilder interface {
	Create() EncryptorBuilder
	WithDecrypt(decrypt BytesReference) EncryptorBuilder
	WithEncrypt(encrypt BytesReference) EncryptorBuilder
	IsPublicKey() EncryptorBuilder
	Now() (Encryptor, error)
}

// Encryptor represents encryptor
type Encryptor interface {
	Hash() hash.Hash
	IsDecrypt() bool
	Decrypt() BytesReference
	IsEncrypt() bool
	Encrypt() BytesReference
	IsPublicKey() bool
}

// SignerBuilder represents a signer builder
type SignerBuilder interface {
	Create() SignerBuilder
	WithSign(sign BytesReference) SignerBuilder
	WithVote(vote Vote) SignerBuilder
	WithGenerateSignerPublicKey(genPubKey uint) SignerBuilder
	WithHashPublicKeys(hashPubKeys string) SignerBuilder
	WithVoteVerify(voteVerify VoteVerify) SignerBuilder
	WithSignatureVerify(sigVerify SignatureVerify) SignerBuilder
	WithBytes(bytes string) SignerBuilder
	IsPublicKey() SignerBuilder
	Now() (Signer, error)
}

// Signer represents the signer identity assignable
type Signer interface {
	Hash() hash.Hash
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

// SignatureVerifyBuilder represents a signature verify builder
type SignatureVerifyBuilder interface {
	Create() SignatureVerifyBuilder
	WithSignature(signature string) SignatureVerifyBuilder
	WithMessage(message BytesReference) SignatureVerifyBuilder
	Now() (SignatureVerify, error)
}

// SignatureVerify represents a signature verify
type SignatureVerify interface {
	Hash() hash.Hash
	Signature() string
	Message() BytesReference
}

// VoteVerifyBuilder represents a vote verify builder
type VoteVerifyBuilder interface {
	Create() VoteVerifyBuilder
	WithVote(vote string) VoteVerifyBuilder
	WithMessage(msg BytesReference) VoteVerifyBuilder
	WithHashedRing(hashedRing string) VoteVerifyBuilder
	Now() (VoteVerify, error)
}

// VoteVerify represents a vote verify
type VoteVerify interface {
	Hash() hash.Hash
	Vote() string
	Message() BytesReference
	HashedRing() string
}

// VoteBuilder represents a vote builder
type VoteBuilder interface {
	Create() VoteBuilder
	WithRing(ring string) VoteBuilder
	WithMessage(message BytesReference) VoteBuilder
	Now() (Vote, error)
}

// Vote represents a vote
type Vote interface {
	Hash() hash.Hash
	Ring() string
	Message() BytesReference
}

// BytesReferencesBuilder represents the bytes references builder
type BytesReferencesBuilder interface {
	Create() BytesReferencesBuilder
	WithList(list []BytesReference) BytesReferencesBuilder
	Now() (BytesReferences, error)
}

// BytesReferences represents bytes values
type BytesReferences interface {
	Hash() hash.Hash
	List() []BytesReference
}

// BytesReferenceBuilder represents the bytes reference builder
type BytesReferenceBuilder interface {
	Create() BytesReferenceBuilder
	WithVariable(variable string) BytesReferenceBuilder
	WithBytes(bytes []byte) BytesReferenceBuilder
	Now() (BytesReference, error)
}

// BytesReference a bytes value
type BytesReference interface {
	Hash() hash.Hash
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
