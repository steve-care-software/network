package layers

// Layer represents a layer
type Layer interface {
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
	IsDecrypt() bool
	Decrypt() []byte
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
	Decrypt() []byte
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
	Vote() []byte
	IsVerify() bool
	Verify() []byte
	IsPublicKey() bool
	PublicKey() []byte
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
