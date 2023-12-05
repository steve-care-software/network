package layers

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/layers"
)

type layer struct {
	layer            layers.Layer
	output           layers.Output
	kind             layers.Kind
	instruction      layers.Instruction
	linkInstruction  layers.LinkInstruction
	layerInstruction layers.LayerInstruction
	condition        layers.Condition
	assignment       layers.Assignment
	assignable       layers.Assignable
	bytes            layers.Bytes
	identity         layers.Identity
	encryptor        layers.Encryptor
	signer           layers.Signer
	signatureVerify  layers.SignatureVerify
	voteVerify       layers.VoteVerify
	vote             layers.Vote
	bytesReference   layers.BytesReference
}

func createLayerWithLayer(
	layerIns layers.Layer,
) Layer {
	return createLayerInternally(
		layerIns,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLayerWithOutput(
	output layers.Output,
) Layer {
	return createLayerInternally(
		nil,
		output,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLayerWithKind(
	kind layers.Kind,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		kind,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLayerWithInstruction(
	instruction layers.Instruction,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		nil,
		instruction,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLayerWithLinkInstruction(
	linkInstruction layers.LinkInstruction,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		nil,
		nil,
		linkInstruction,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLayerWithLayerInstruction(
	layerInstruction layers.LayerInstruction,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		layerInstruction,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLayerWithCondition(
	condition layers.Condition,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		condition,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLayerWithAssignment(
	assignment layers.Assignment,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		assignment,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLayerWithAssignable(
	assignable layers.Assignable,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		assignable,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLayerWithBytes(
	bytes layers.Bytes,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		bytes,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLayerWithIdentity(
	identity layers.Identity,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		identity,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLayerWithEncryptor(
	encryptor layers.Encryptor,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		encryptor,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLayerWithSigner(
	signer layers.Signer,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		signer,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLayerWithSignatureVerify(
	signatureVerify layers.SignatureVerify,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		signatureVerify,
		nil,
		nil,
		nil,
	)
}

func createLayerWithVoteVerify(
	voteVerify layers.VoteVerify,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		voteVerify,
		nil,
		nil,
	)
}

func createLayerWithVote(
	vote layers.Vote,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		vote,
		nil,
	)
}

func createLayerWithBytesReference(
	bytesReference layers.BytesReference,
) Layer {
	return createLayerInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		bytesReference,
	)
}

func createLayerInternally(
	layerIns layers.Layer,
	output layers.Output,
	kind layers.Kind,
	instruction layers.Instruction,
	linkInstruction layers.LinkInstruction,
	layerInstruction layers.LayerInstruction,
	condition layers.Condition,
	assignment layers.Assignment,
	assignable layers.Assignable,
	bytes layers.Bytes,
	identity layers.Identity,
	encryptor layers.Encryptor,
	signer layers.Signer,
	signatureVerify layers.SignatureVerify,
	voteVerify layers.VoteVerify,
	vote layers.Vote,
	bytesReference layers.BytesReference,
) Layer {
	out := layer{
		layer:            layerIns,
		output:           output,
		kind:             kind,
		instruction:      instruction,
		linkInstruction:  linkInstruction,
		layerInstruction: layerInstruction,
		condition:        condition,
		assignment:       assignment,
		assignable:       assignable,
		bytes:            bytes,
		identity:         identity,
		encryptor:        encryptor,
		signer:           signer,
		signatureVerify:  signatureVerify,
		voteVerify:       voteVerify,
		vote:             vote,
		bytesReference:   bytesReference,
	}

	return &out
}

// Hash returns the hash
func (obj *layer) Hash() hash.Hash {
	if obj.IsLayer() {
		return obj.layer.Hash()
	}

	if obj.IsOutput() {
		return obj.output.Hash()
	}

	if obj.IsInstruction() {
		return obj.instruction.Hash()
	}

	if obj.IsLinkInstruction() {
		return obj.linkInstruction.Hash()
	}

	if obj.IsLayerInstruction() {
		return obj.layerInstruction.Hash()
	}

	if obj.IsCondition() {
		return obj.condition.Hash()
	}

	if obj.IsAssignment() {
		return obj.assignment.Hash()
	}

	if obj.IsAssignable() {
		return obj.assignable.Hash()
	}

	if obj.IsBytes() {
		return obj.bytes.Hash()
	}

	if obj.IsIdentity() {
		return obj.identity.Hash()
	}

	if obj.IsEncryptor() {
		return obj.encryptor.Hash()
	}

	if obj.IsSigner() {
		return obj.signer.Hash()
	}

	if obj.IsSignatureVerify() {
		return obj.signatureVerify.Hash()
	}

	if obj.IsVoteVerify() {
		return obj.voteVerify.Hash()
	}

	if obj.IsVote() {
		return obj.vote.Hash()
	}

	return obj.bytesReference.Hash()
}

// IsLayer returns true if there is a layer, false otherwise
func (obj *layer) IsLayer() bool {
	return obj.layer != nil
}

// Layer returns the layer, if any
func (obj *layer) Layer() layers.Layer {
	return obj.layer
}

// IsOutput returns true if there is an output, false otherwise
func (obj *layer) IsOutput() bool {
	return obj.output != nil
}

// Output returns the output, if any
func (obj *layer) Output() layers.Output {
	return obj.output
}

// IsKind returns true if there is a kind, false otherwise
func (obj *layer) IsKind() bool {
	return obj.kind != nil
}

// Kind returns the kind, if any
func (obj *layer) Kind() layers.Kind {
	return obj.kind
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *layer) IsInstruction() bool {
	return obj.instruction != nil
}

// Instruction returns the instruction, if any
func (obj *layer) Instruction() layers.Instruction {
	return obj.instruction
}

// IsLinkInstruction returns true if there is a link instruction, false otherwise
func (obj *layer) IsLinkInstruction() bool {
	return obj.linkInstruction != nil
}

// LinkInstruction returns the link instruction, if any
func (obj *layer) LinkInstruction() layers.LinkInstruction {
	return obj.linkInstruction
}

// IsLayerInstruction returns true if there is a layer instruction, false otherwise
func (obj *layer) IsLayerInstruction() bool {
	return obj.layerInstruction != nil
}

// LayerInstruction returns the layer instruction, if any
func (obj *layer) LayerInstruction() layers.LayerInstruction {
	return obj.layerInstruction
}

// IsCondition returns true if there is a condition, false otherwise
func (obj *layer) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *layer) Condition() layers.Condition {
	return obj.condition
}

// IsAssignment returns true if there is an assignment, false otherwise
func (obj *layer) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *layer) Assignment() layers.Assignment {
	return obj.assignment
}

// IsAssignable returns true if there is an assignable, false otherwise
func (obj *layer) IsAssignable() bool {
	return obj.assignable != nil
}

// Assignable returns the assignable, if any
func (obj *layer) Assignable() layers.Assignable {
	return obj.assignable
}

// IsBytes returns true if there is a bytes, false otherwise
func (obj *layer) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes, if any
func (obj *layer) Bytes() layers.Bytes {
	return obj.bytes
}

// IsIdentity returns true if there is an identity, false otherwise
func (obj *layer) IsIdentity() bool {
	return obj.identity != nil
}

// Identity returns the identity, if any
func (obj *layer) Identity() layers.Identity {
	return obj.identity
}

// IsEncryptor returns true if there is an encryptor, false otherwise
func (obj *layer) IsEncryptor() bool {
	return obj.encryptor != nil
}

// Encryptor returns the encryptor, if any
func (obj *layer) Encryptor() layers.Encryptor {
	return obj.encryptor
}

// IsSigner returns true if there is a signer, false otherwise
func (obj *layer) IsSigner() bool {
	return obj.signer != nil
}

// Signer returns the signer, if any
func (obj *layer) Signer() layers.Signer {
	return obj.signer
}

// IsSignatureVerify returns true if there is a signatureVerify, false otherwise
func (obj *layer) IsSignatureVerify() bool {
	return obj.signatureVerify != nil
}

// SignatureVerify returns the signatureVerify, if any
func (obj *layer) SignatureVerify() layers.SignatureVerify {
	return obj.signatureVerify
}

// IsVoteVerify returns true if there is a voteVerify, false otherwise
func (obj *layer) IsVoteVerify() bool {
	return obj.voteVerify != nil
}

// VoteVerfy returns the voteVerify, if any
func (obj *layer) VoteVerfy() layers.VoteVerify {
	return obj.voteVerify
}

// IsVote returns true if there is a vote, false otherwise
func (obj *layer) IsVote() bool {
	return obj.vote != nil
}

// Vote returns the vote, if any
func (obj *layer) Vote() layers.Vote {
	return obj.vote
}

// IsBytesReference returns true if there is a bytesReference, false otherwise
func (obj *layer) IsBytesReference() bool {
	return obj.bytesReference != nil
}

// BytesRefernece returns the bytesReference, if any
func (obj *layer) BytesRefernece() layers.BytesReference {
	return obj.bytesReference
}
