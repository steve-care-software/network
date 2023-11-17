package layers

// Layer represents a layer
type Layer interface {
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
	IsExecute() bool
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
	Save() Resource
	IsAssignment() bool
	Assignment() Assignment
}

// Resource represents a resource
type Resource interface {
	Path() []string
	Layer() Layer
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
	IsAuthorized() bool
	IsAuthenticated() bool
	IsIdentity() bool
	Identity() Identity
}

// Identity represents the assignable identity
type Identity interface {
	IsSigner() bool
	Signer() Signer
	IsVoter() bool
	Voter() Voter
	IsEncryptor() bool
	Encryptor() Encryptor
}

// Bytes represents the bytes assignable
type Bytes interface {
	IsJoin() bool
	Join() Values
	IsCompare() bool
	Compare() Values
}

// Encryptor represents encryptor
type Encryptor interface {
	IsDecrypt() bool
	Decrypt() Value
	IsEncrypt() bool
	Encrypt() Value
	IsPublicKey() bool
}

// Signer represents the signer identity assignable
type Signer interface {
	IsSign() bool
	Sign() []byte
	IsVerify() bool
	IsPublicKey() bool
}

// Voter represents the vote identity assignable
type Voter interface {
	IsVote() bool
	Vote() Value
	IsVerify() bool
	Verify() Value
	IsPublicKey() bool
	PublicKey() Value
}

// Values represents values
type Values interface {
	List() []Value
}

// Value a value
type Value interface {
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
