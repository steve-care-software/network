package layers

// NewLayersForTests creates a new layers for tests
func NewLayersForTests(list []Layer) Layers {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithInputForTests creates a new layer with input for tests
func NewLayerWithInputForTests(instructions Instructions, output Output, input string) Layer {
	ins, err := NewLayerBuilder().Create().WithInstructions(instructions).WithOutput(output).WithInput(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerForTests creates a new layer for tests
func NewLayerForTests(instructions Instructions, output Output) Layer {
	ins, err := NewLayerBuilder().Create().WithInstructions(instructions).WithOutput(output).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOutputWithExecuteForTests creates a new output with execute for tests
func NewOutputWithExecuteForTests(variable string, kind Kind, execute string) Output {
	ins, err := NewOutputBuilder().Create().WithVariable(variable).WithKind(kind).WithExecute(execute).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOutputForTests creates a new output for tests
func NewOutputForTests(variable string, kind Kind) Output {
	ins, err := NewOutputBuilder().Create().WithVariable(variable).WithKind(kind).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewKindWithContinueForTests creates a new kind with continue for tests
func NewKindWithContinueForTests() Kind {
	ins, err := NewKindBuilder().Create().IsContinue().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewKindWithPromptForTests creates a new kind with prompt for tests
func NewKindWithPromptForTests() Kind {
	ins, err := NewKindBuilder().Create().IsPrompt().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionsForTests creates new instructions for tests
func NewInstructionsForTests(list []Instruction) Instructions {
	ins, err := NewInstructionsBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionWithAssignmentForTests creates a new instruction with assignment for tests
func NewInstructionWithAssignmentForTests(assignment Assignment) Instruction {
	ins, err := NewInstructionBuilder().Create().WithAssignment(assignment).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionWithConditionForTests creates a new instruction with condition for tests
func NewInstructionWithConditionForTests(condition Condition) Instruction {
	ins, err := NewInstructionBuilder().Create().WithCondition(condition).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionWithRaiseErrorForTests creates a new instruction with raiseError for tests
func NewInstructionWithRaiseErrorForTests(raiseError uint) Instruction {
	ins, err := NewInstructionBuilder().Create().WithRaiseError(raiseError).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionWithStopForTests creates a new instruction with stop for tests
func NewInstructionWithStopForTests() Instruction {
	ins, err := NewInstructionBuilder().Create().IsStop().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConditionForTest creates a new condition for tests
func NewConditionForTest(variable string, instructions Instructions) Condition {
	ins, err := NewConditionBuilder().Create().WithVariable(variable).WithInstructions(instructions).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignmentForTests creates a new assignment for tests
func NewAssignmentForTests(name string, assignable Assignable) Assignment {
	ins, err := NewAssignmentBuilder().Create().WithName(name).WithAssignable(assignable).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithIdentityForTests creates a new assignable with identity for tests
func NewAssignableWithIdentityForTests(input Identity) Assignable {
	ins, err := NewAssignableBuilder().Create().WithIdentity(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableWithBytesForTests creates a new assignable with bytes for tests
func NewAssignableWithBytesForTests(input Bytes) Assignable {
	ins, err := NewAssignableBuilder().Create().WithBytes(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewEngineWithExecutionForTests creates a new engine with execution for tests
func NewEngineWithExecutionForTests(input BytesReference) Engine {
	ins, err := NewEngineBuilder().Create().WithExecution(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewEngineWithResourceForTests creates a new engine with resource for tests
func NewEngineWithResourceForTests(input AssignableResource) Engine {
	ins, err := NewEngineBuilder().Create().WithResource(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableResourceWithCompileForTests creates a new assignableResource with compile for tests
func NewAssignableResourceWithCompileForTests(input BytesReference) AssignableResource {
	ins, err := NewAssignableResourceBuilder().Create().WithCompile(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableResourceWithDecompileForTests creates a new assignableResource with decompile for tests
func NewAssignableResourceWithDecompileForTests(input string) AssignableResource {
	ins, err := NewAssignableResourceBuilder().Create().WithDecompile(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableResourceWithAmountByQueryForTests creates a new assignableResource with amountByQuery for tests
func NewAssignableResourceWithAmountByQueryForTests(input BytesReference) AssignableResource {
	ins, err := NewAssignableResourceBuilder().Create().WihAmountByQuery(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableResourceWithRetrieveByQueryForTests creates a new assignableResource with retrieveByQuery for tests
func NewAssignableResourceWithRetrieveByQueryForTests(input BytesReference) AssignableResource {
	ins, err := NewAssignableResourceBuilder().Create().WithRetrieveByQuery(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableResourceWithRetrieveByHashForTests creates a new assignableResource with retrieveByHash for tests
func NewAssignableResourceWithRetrieveByHashForTests(input BytesReference) AssignableResource {
	ins, err := NewAssignableResourceBuilder().Create().WithRetrieveByHash(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewAssignableResourceWithAmountForTests creates a new assignableResource with amount for tests
func NewAssignableResourceWithAmountForTests() AssignableResource {
	ins, err := NewAssignableResourceBuilder().Create().IsAmount().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewBytesWithHashBytesForTests creates a new bytes with hashBytes for tests
func NewBytesWithHashBytesForTests(input BytesReference) Bytes {
	ins, err := NewBytesBuilder().Create().WithHashBytes(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewBytesWithCompareForTests creates a new bytes with compare for tests
func NewBytesWithCompareForTests(input BytesReferences) Bytes {
	ins, err := NewBytesBuilder().Create().WithCompare(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewBytesWithJoinForTests creates a new bytes with join for tests
func NewBytesWithJoinForTests(join BytesReferences) Bytes {
	ins, err := NewBytesBuilder().Create().WithJoin(join).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIdentityWithEncryptorForTests creates a new identity with encryptor for tests
func NewIdentityWithEncryptorForTests(input Encryptor) Identity {
	ins, err := NewIdentityBuilder().Create().WithEncryptor(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIdentityWithSignerForTests creates a new identity with signer for tests
func NewIdentityWithSignerForTests(signer Signer) Identity {
	ins, err := NewIdentityBuilder().Create().WithSigner(signer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewEncryptorWithPublicKeyForTests creates a new encryptor with publicKey for tests
func NewEncryptorWithPublicKeyForTests() Encryptor {
	ins, err := NewEncryptorBuilder().Create().IsPublicKey().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewEncryptorWithEncryptForTests creates a new encryptor with encrypt with tests
func NewEncryptorWithEncryptForTests(input BytesReference) Encryptor {
	ins, err := NewEncryptorBuilder().Create().WithEncrypt(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewEncryptorWithDecryptForTests creates a new encryptor with decrypt with tests
func NewEncryptorWithDecryptForTests(input BytesReference) Encryptor {
	ins, err := NewEncryptorBuilder().Create().WithDecrypt(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSignerWithPublicKeyForTests creates a new signer with publicKey for tests
func NewSignerWithPublicKeyForTests() Signer {
	ins, err := NewSignerBuilder().Create().IsPublicKey().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSignerWithBytesForTests creates a new signer with bytes for tests
func NewSignerWithBytesForTests(input string) Signer {
	ins, err := NewSignerBuilder().Create().WithBytes(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSignerWithSignatureVerifyForTests creates a new signer with signatureVerify for tests
func NewSignerWithSignatureVerifyForTests(input SignatureVerify) Signer {
	ins, err := NewSignerBuilder().Create().WithSignatureVerify(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSignerWithVoteVerifyForTests creates a new signer with voteVerify for tests
func NewSignerWithVoteVerifyForTests(input VoteVerify) Signer {
	ins, err := NewSignerBuilder().Create().WithVoteVerify(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSignerWithHashPublicKeysForTests creates a new signer with hashPublicKeys for tests
func NewSignerWithHashPublicKeysForTests(input string) Signer {
	ins, err := NewSignerBuilder().Create().WithHashPublicKeys(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSignerWithGenerateSignerPublicKeyForTests creates a new signer with generateSignerPublicKey for tests
func NewSignerWithGenerateSignerPublicKeyForTests(input uint) Signer {
	ins, err := NewSignerBuilder().Create().WithGenerateSignerPublicKey(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSignerWithVoteForTests creates a new signer with vote for tests
func NewSignerWithVoteForTests(input Vote) Signer {
	ins, err := NewSignerBuilder().Create().WithVote(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSignerWithSignForTests creates a new signer with sign for tests
func NewSignerWithSignForTests(input BytesReference) Signer {
	ins, err := NewSignerBuilder().Create().WithSign(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSignatureVerifyForTests creates a new signature verify for tests
func NewSignatureVerifyForTests(signature string, message BytesReference) SignatureVerify {
	ins, err := NewSignatureVerifyBuilder().Create().
		WithSignature(signature).
		WithMessage(message).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewVoteVerifyForTests creates a new vote verify for tests
func NewVoteVerifyForTests(vote string, message BytesReference, hashedRing string) VoteVerify {
	ins, err := NewVoteVerifyBuilder().Create().
		WithVote(vote).
		WithMessage(message).
		WithHashedRing(hashedRing).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewVoteForTests creates a new vote for tests
func NewVoteForTests(ring string, message BytesReference) Vote {
	ins, err := NewVoteBuilder().Create().
		WithRing(ring).
		WithMessage(message).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewBytesReferencesForTests creates a new bytes references for tests
func NewBytesReferencesForTests(list []BytesReference) BytesReferences {
	ins, err := NewBytesReferencesBuilder().
		Create().
		WithList(list).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewBytesReferenceWithBytesForTests creates a new bytes reference with bytes for tests
func NewBytesReferenceWithBytesForTests(bytes []byte) BytesReference {
	ins, err := NewBytesReferenceBuilder().
		Create().
		WithBytes(bytes).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewBytesReferenceWithVariableForTests creates a new bytes reference with variable for tests
func NewBytesReferenceWithVariableForTests(variable string) BytesReference {
	ins, err := NewBytesReferenceBuilder().
		Create().
		WithVariable(variable).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
