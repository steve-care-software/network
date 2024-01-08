package layers

import "steve.care/network/domain/programs/logics/libraries/layers"

// NewLayerWithLayerForTests creates a new layer with layer for tests
func NewLayerWithLayerForTests(input layers.Layer) Layer {
	ins, err := NewBuilder().Create().WithLayer(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithOutputForTests creates a new layer with output for tests
func NewLayerWithOutputForTests(input layers.Output) Layer {
	ins, err := NewBuilder().Create().WithOutput(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithKindForTests creates a new layer with kind for tests
func NewLayerWithKindForTests(input layers.Kind) Layer {
	ins, err := NewBuilder().Create().WithKind(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithInstructionForTests creates a new layer with instruction for tests
func NewLayerWithInstructionForTests(input layers.Instruction) Layer {
	ins, err := NewBuilder().Create().WithInstruction(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithConditionForTests creates a new layer with condition for tests
func NewLayerWithConditionForTests(input layers.Condition) Layer {
	ins, err := NewBuilder().Create().WithCondition(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithAssignmentForTests creates a new layer with assignment for tests
func NewLayerWithAssignmentForTests(input layers.Assignment) Layer {
	ins, err := NewBuilder().Create().WithAssignment(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithAssignableForTests creates a new layer with assignable for tests
func NewLayerWithAssignableForTests(input layers.Assignable) Layer {
	ins, err := NewBuilder().Create().WithAssignable(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithEngineForTests creates a new layer with engine for tests
func NewLayerWithEngineForTests(input layers.Engine) Layer {
	ins, err := NewBuilder().Create().WithEngine(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithAssignableResourceForTests creates a new layer with assignableResource for tests
func NewLayerWithAssignableResourceForTests(input layers.AssignableResource) Layer {
	ins, err := NewBuilder().Create().WithAssignableResource(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithBytesForTests creates a new layer with bytes for tests
func NewLayerWithBytesForTests(input layers.Bytes) Layer {
	ins, err := NewBuilder().Create().WithBytes(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithIdentityForTests creates a new layer with identity for tests
func NewLayerWithIdentityForTests(input layers.Identity) Layer {
	ins, err := NewBuilder().Create().WithIdentity(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithEncryptorForTests creates a new layer with encryptor for tests
func NewLayerWithEncryptorForTests(input layers.Encryptor) Layer {
	ins, err := NewBuilder().Create().WithEncryptor(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithSignerForTests creates a new layer with signer for tests
func NewLayerWithSignerForTests(input layers.Signer) Layer {
	ins, err := NewBuilder().Create().WithSigner(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithSignatureVerifyForTests creates a new layer with signatureVerify for tests
func NewLayerWithSignatureVerifyForTests(input layers.SignatureVerify) Layer {
	ins, err := NewBuilder().Create().WithSignatureVerify(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithVoteVerifyForTests creates a new layer with voteVerify for tests
func NewLayerWithVoteVerifyForTests(input layers.VoteVerify) Layer {
	ins, err := NewBuilder().Create().WithVoteVerify(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithVoteForTests creates a new layer with vote for tests
func NewLayerWithVoteForTests(input layers.Vote) Layer {
	ins, err := NewBuilder().Create().WithVote(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithBytesReferenceForTests creates a new layer with bytes reference for tests
func NewLayerWithBytesReferenceForTests(input layers.BytesReference) Layer {
	ins, err := NewBuilder().Create().WithBytesReference(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
