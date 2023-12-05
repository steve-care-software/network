package layers

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/layers"
)

// Builder represents a layer builder
type Builder interface {
	Create() Builder
	WithLayer(layer layers.Layer) Builder
	WithOutput(output layers.Output) Builder
	WithKind(kind layers.Kind) Builder
	WithInstruction(ins layers.Instruction) Builder
	WithLayerInstructions(layerIns layers.LayerInstruction) Builder
	WithCondition(condition layers.Condition) Builder
	WithAssignment(assignment layers.Assignment) Builder
	WithAssignable(assignable layers.Assignable) Builder
	WithBytes(bytes layers.Bytes) Builder
	WithIdentity(identity layers.Identity) Builder
	WithEncryptor(encryptor layers.Encryptor) Builder
	WithSigner(signer layers.Signer) Builder
	WithSignatureVerify(sigVerify layers.SignatureVerify) Builder
	WithVoteVerify(voteVerify layers.VoteVerify) Builder
	WithVote(vote layers.Vote) Builder
	WithBytesReference(bytesReference layers.BytesReference) Builder
	Now() (Layer, error)
}

// Layer represents a layer resource
type Layer interface {
	Hash() hash.Hash
	IsLayer() bool
	Layer() layers.Layer
	IsOutput() bool
	Output() layers.Output
	IsKind() bool
	Kind() layers.Kind
	IsInstruction() bool
	Instruction() layers.Instruction
	IsLinkInstruction() bool
	LinkInstruction() layers.LinkInstruction
	IsLayerInstruction() bool
	LayerInstruction() layers.LayerInstruction
	IsCondition() bool
	Condition() layers.Condition
	IsAssignment() bool
	Assignment() layers.Assignment
	IsAssignable() bool
	Assignable() layers.Assignable
	IsBytes() bool
	Bytes() layers.Bytes
	IsIdentity() bool
	Identity() layers.Identity
	IsEncryptor() bool
	Encryptor() layers.Encryptor
	IsSigner() bool
	Signer() layers.Signer
	IsSignatureVerify() bool
	SignatureVerify() layers.SignatureVerify
	IsVoteVerify() bool
	VoteVerfy() layers.VoteVerify
	IsVote() bool
	Vote() layers.Vote
	IsBytesReference() bool
	BytesRefernece() layers.BytesReference
}
