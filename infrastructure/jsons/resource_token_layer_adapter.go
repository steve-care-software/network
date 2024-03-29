package jsons

import (
	"steve.care/network/domain/hash"
	resources_layers "steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/logics/libraries/layers"
	structs_tokens "steve.care/network/infrastructure/jsons/resources/tokens"
	structs_layers "steve.care/network/infrastructure/jsons/resources/tokens/layers"
)

type resourceTokenLayerAdapter struct {
	hashAdapter            hash.Adapter
	builder                resources_layers.Builder
	linkAdapter            *resourceTokenLinkAdapter
	layerBuilder           layers.LayerBuilder
	outputBuilder          layers.OutputBuilder
	kindBuilder            layers.KindBuilder
	instructionsBuilder    layers.InstructionsBuilder
	instructionBuilder     layers.InstructionBuilder
	conditionBuilder       layers.ConditionBuilder
	assignmentBuilder      layers.AssignmentBuilder
	assignableBuilder      layers.AssignableBuilder
	bytesBuilder           layers.BytesBuilder
	identityBuilder        layers.IdentityBuilder
	encryptorBuilder       layers.EncryptorBuilder
	signerBuilder          layers.SignerBuilder
	signatureVerifyBuilder layers.SignatureVerifyBuilder
	voteVerifyBuilder      layers.VoteVerifyBuilder
	voteBuilder            layers.VoteBuilder
}

func (app *resourceTokenLayerAdapter) toStruct(ins resources_layers.Layer) structs_tokens.Layer {
	output := structs_tokens.Layer{}
	if ins.IsLayer() {
		layer := app.layerToStruct(ins.Layer())
		output.Layer = &layer
	}

	if ins.IsOutput() {
		outputIns := app.outputToStruct(ins.Output())
		output.Output = &outputIns
	}

	if ins.IsKind() {
		kind := app.kindToStruct(ins.Kind())
		output.Kind = &kind
	}

	if ins.IsInstruction() {
		instruction := app.instructionToStruct(ins.Instruction())
		output.Instruction = &instruction
	}

	if ins.IsCondition() {
		condition := app.conditionToStruct(ins.Condition())
		output.Condition = &condition
	}

	if ins.IsAssignment() {
		assignment := app.assignmentToStruct(ins.Assignment())
		output.Assignment = &assignment
	}

	if ins.IsAssignable() {
		assignable := app.assignableToStruct(ins.Assignable())
		output.Assignable = &assignable
	}

	if ins.IsBytes() {
		bytes := app.bytesToStruct(ins.Bytes())
		output.Bytes = &bytes
	}

	if ins.IsIdentity() {
		identity := app.identityToStruct(ins.Identity())
		output.Identity = &identity
	}

	if ins.IsEncryptor() {
		encryptor := app.encryptorToStruct(ins.Encryptor())
		output.Encryptor = &encryptor
	}

	if ins.IsSigner() {
		signer := app.signerToStruct(ins.Signer())
		output.Signer = &signer
	}

	if ins.IsSignatureVerify() {
		signatureVerify := app.signatureVerifyToStruct(ins.SignatureVerify())
		output.SignatureVerify = &signatureVerify
	}

	if ins.IsVoteVerify() {
		voteVerify := app.voteVerifyToStruct(ins.VoteVerify())
		output.VoteVerify = &voteVerify
	}

	if ins.IsVote() {
		vote := app.voteToStruct(ins.Vote())
		output.Vote = &vote
	}

	return output
}

func (app *resourceTokenLayerAdapter) toInstance(ins structs_tokens.Layer) (resources_layers.Layer, error) {
	builder := app.builder.Create()
	if ins.Layer != nil {
		layer, err := app.structToLayer(*ins.Layer)
		if err != nil {
			return nil, err
		}

		builder.WithLayer(layer)
	}

	if ins.Output != nil {
		output, err := app.structToOutput(*ins.Output)
		if err != nil {
			return nil, err
		}

		builder.WithOutput(output)
	}

	if ins.Kind != nil {
		kind, err := app.structToKind(*ins.Kind)
		if err != nil {
			return nil, err
		}

		builder.WithKind(kind)
	}

	if ins.Instruction != nil {
		instruction, err := app.structToInstruction(*ins.Instruction)
		if err != nil {
			return nil, err
		}

		builder.WithInstruction(instruction)
	}

	if ins.Condition != nil {
		condition, err := app.structToCondition(*ins.Condition)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	if ins.Assignment != nil {
		assignment, err := app.structToAssignment(*ins.Assignment)
		if err != nil {
			return nil, err
		}

		builder.WithAssignment(assignment)
	}

	if ins.Assignable != nil {
		assignable, err := app.structToAssignable(*ins.Assignable)
		if err != nil {
			return nil, err
		}

		builder.WithAssignable(assignable)
	}

	if ins.Bytes != nil {
		bytes, err := app.structToBytes(*ins.Bytes)
		if err != nil {
			return nil, err
		}

		builder.WithBytes(bytes)
	}

	if ins.Identity != nil {
		identity, err := app.structToIdentity(*ins.Identity)
		if err != nil {
			return nil, err
		}

		builder.WithIdentity(identity)
	}

	if ins.Encryptor != nil {
		encryptor, err := app.structToEncryptor(*ins.Encryptor)
		if err != nil {
			return nil, err
		}

		builder.WithEncryptor(encryptor)
	}

	if ins.Signer != nil {
		signer, err := app.structToSigner(*ins.Signer)
		if err != nil {
			return nil, err
		}

		builder.WithSigner(signer)
	}

	if ins.SignatureVerify != nil {
		signatureVerify, err := app.structToSignatureVerify(*ins.SignatureVerify)
		if err != nil {
			return nil, err
		}

		builder.WithSignatureVerify(signatureVerify)
	}

	if ins.VoteVerify != nil {
		voteVerify, err := app.structToVoteVerify(*ins.VoteVerify)
		if err != nil {
			return nil, err
		}

		builder.WithVoteVerify(voteVerify)
	}

	if ins.Vote != nil {
		vote, err := app.structToVote(*ins.Vote)
		if err != nil {
			return nil, err
		}

		builder.WithVote(vote)
	}

	return builder.Now()
}

func (app *resourceTokenLayerAdapter) layerToStruct(
	ins layers.Layer,
) structs_layers.Layer {
	instructions := app.instructionsToStructs(ins.Instructions())
	output := app.outputToStruct(ins.Output())
	outputLayer := structs_layers.Layer{
		Instructions: instructions,
		Output:       output,
		Input:        ins.Input(),
	}

	return outputLayer
}

func (app *resourceTokenLayerAdapter) structToLayer(
	ins structs_layers.Layer,
) (layers.Layer, error) {
	instructions, err := app.structsToInstructions(ins.Instructions)
	if err != nil {
		return nil, err
	}

	output, err := app.structToOutput(ins.Output)
	if err != nil {
		return nil, err
	}

	builder := app.layerBuilder.Create().WithInstructions(instructions).WithOutput(output)
	if ins.Input != "" {
		builder.WithInput(ins.Input)
	}

	return builder.Now()
}

func (app *resourceTokenLayerAdapter) outputToStruct(
	ins layers.Output,
) structs_layers.Output {
	kind := app.kindToStruct(ins.Kind())
	output := structs_layers.Output{
		Variable: ins.Variable(),
		Kind:     kind,
	}

	if ins.HasExecute() {
		output.Execute = ins.Execute()
	}

	return output
}

func (app *resourceTokenLayerAdapter) structToOutput(
	ins structs_layers.Output,
) (layers.Output, error) {
	kind, err := app.structToKind(ins.Kind)
	if err != nil {
		return nil, err
	}

	return app.outputBuilder.Create().
		WithVariable(ins.Variable).
		WithKind(kind).
		WithExecute(ins.Execute).
		Now()
}

func (app *resourceTokenLayerAdapter) kindToStruct(
	ins layers.Kind,
) structs_layers.Kind {
	output := structs_layers.Kind{}
	if ins.IsContinue() {
		output.IsContinue = ins.IsContinue()
	}

	if ins.IsPrompt() {
		output.IsPrompt = ins.IsPrompt()
	}

	return output
}

func (app *resourceTokenLayerAdapter) structToKind(
	ins structs_layers.Kind,
) (layers.Kind, error) {
	builder := app.kindBuilder.Create()
	if ins.IsContinue {
		builder.IsContinue()
	}

	if ins.IsPrompt {
		builder.IsPrompt()
	}

	return builder.Now()
}

func (app *resourceTokenLayerAdapter) structsToInstructions(
	list []structs_layers.Instruction,
) (layers.Instructions, error) {
	output := []layers.Instruction{}
	for _, oneStruct := range list {
		ins, err := app.structToInstruction(oneStruct)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.instructionsBuilder.Create().
		WithList(output).
		Now()
}

func (app *resourceTokenLayerAdapter) instructionsToStructs(
	ins layers.Instructions,
) []structs_layers.Instruction {
	list := ins.List()
	output := []structs_layers.Instruction{}
	for _, oneIns := range list {
		instruction := app.instructionToStruct(oneIns)
		output = append(output, instruction)
	}

	return output
}

func (app *resourceTokenLayerAdapter) structToInstruction(
	ins structs_layers.Instruction,
) (layers.Instruction, error) {
	builder := app.instructionBuilder.Create()
	if ins.Stop {
		builder.IsStop()
	}

	if ins.RaiseError != nil {
		builder.WithRaiseError(*ins.RaiseError)
	}

	if ins.Condition != nil {
		condition, err := app.structToCondition(*ins.Condition)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	if ins.Assignment != nil {
		assignment, err := app.structToAssignment(*ins.Assignment)
		if err != nil {
			return nil, err
		}

		builder.WithAssignment(assignment)
	}

	return builder.Now()
}

func (app *resourceTokenLayerAdapter) instructionToStruct(
	ins layers.Instruction,
) structs_layers.Instruction {
	output := structs_layers.Instruction{}
	if ins.IsStop() {
		output.Stop = ins.IsStop()
	}

	if ins.IsRaiseError() {
		raiseError := ins.RaiseError()
		output.RaiseError = &raiseError
	}

	if ins.IsCondition() {
		condition := app.conditionToStruct(ins.Condition())
		output.Condition = &condition
	}

	if ins.IsAssignment() {
		assignment := app.assignmentToStruct(ins.Assignment())
		output.Assignment = &assignment
	}

	return output
}

func (app *resourceTokenLayerAdapter) structToCondition(
	ins structs_layers.Condition,
) (layers.Condition, error) {
	instructions, err := app.structsToInstructions(ins.Instructions)
	if err != nil {
		return nil, err
	}

	return app.conditionBuilder.Create().
		WithInstructions(instructions).
		WithVariable(ins.Variable).
		Now()
}

func (app *resourceTokenLayerAdapter) conditionToStruct(
	ins layers.Condition,
) structs_layers.Condition {
	instructions := app.instructionsToStructs(ins.Instructions())
	return structs_layers.Condition{
		Variable:     ins.Variable(),
		Instructions: instructions,
	}
}

func (app *resourceTokenLayerAdapter) structToAssignment(
	ins structs_layers.Assignment,
) (layers.Assignment, error) {
	assignable, err := app.structToAssignable(ins.Assignable)
	if err != nil {
		return nil, err
	}

	return app.assignmentBuilder.Create().
		WithName(ins.Name).
		WithAssignable(assignable).
		Now()
}

func (app *resourceTokenLayerAdapter) assignmentToStruct(
	ins layers.Assignment,
) structs_layers.Assignment {
	assignable := app.assignableToStruct(ins.Assignable())
	return structs_layers.Assignment{
		Name:       ins.Name(),
		Assignable: assignable,
	}
}

func (app *resourceTokenLayerAdapter) structToAssignable(
	ins structs_layers.Assignable,
) (layers.Assignable, error) {
	builder := app.assignableBuilder.Create()
	if ins.Bytes != nil {
		bytes, err := app.structToBytes(*ins.Bytes)
		if err != nil {
			return nil, err
		}

		builder.WithBytes(bytes)
	}

	if ins.Identity != nil {
		identity, err := app.structToIdentity(*ins.Identity)
		if err != nil {
			return nil, err
		}

		builder.WithIdentity(identity)
	}

	return builder.Now()
}

func (app *resourceTokenLayerAdapter) assignableToStruct(
	ins layers.Assignable,
) structs_layers.Assignable {
	output := structs_layers.Assignable{}
	if ins.IsBytes() {
		bytes := app.bytesToStruct(ins.Bytes())
		output.Bytes = &bytes
	}

	if ins.IsIdentity() {
		identity := app.identityToStruct(ins.Identity())
		output.Identity = &identity
	}

	return output
}

func (app *resourceTokenLayerAdapter) bytesToStruct(
	ins layers.Bytes,
) structs_layers.Bytes {
	output := structs_layers.Bytes{}
	if ins.IsJoin() {
		output.Join = ins.Join()
	}

	if ins.IsCompare() {
		compare := ins.Compare()
		output.Compare = compare
	}

	if ins.IsHashBytes() {
		output.Hash = ins.HashBytes()
	}

	return output
}

func (app *resourceTokenLayerAdapter) structToBytes(
	ins structs_layers.Bytes,
) (layers.Bytes, error) {
	builder := app.bytesBuilder.Create()
	if ins.Join != nil && len(ins.Join) > 0 {
		builder.WithJoin(ins.Join)
	}

	if ins.Compare != nil && len(ins.Compare) > 0 {
		builder.WithCompare(ins.Compare)
	}

	if ins.Hash != "" {
		builder.WithHashBytes(ins.Hash)
	}

	return builder.Now()
}

func (app *resourceTokenLayerAdapter) identityToStruct(
	ins layers.Identity,
) structs_layers.Identity {
	output := structs_layers.Identity{}
	if ins.IsSigner() {
		signer := app.signerToStruct(ins.Signer())
		output.Signer = &signer
	}

	if ins.IsEncryptor() {
		encryptor := app.encryptorToStruct(ins.Encryptor())
		output.Encryptor = &encryptor
	}

	return output
}

func (app *resourceTokenLayerAdapter) structToIdentity(
	ins structs_layers.Identity,
) (layers.Identity, error) {
	builder := app.identityBuilder.Create()
	if ins.Signer != nil {
		signer, err := app.structToSigner(*ins.Signer)
		if err != nil {
			return nil, err
		}

		builder.WithSigner(signer)
	}

	if ins.Encryptor != nil {
		encryptor, err := app.structToEncryptor(*ins.Encryptor)
		if err != nil {
			return nil, err
		}

		builder.WithEncryptor(encryptor)
	}

	return builder.Now()
}

func (app *resourceTokenLayerAdapter) encryptorToStruct(
	ins layers.Encryptor,
) structs_layers.Encryptor {
	output := structs_layers.Encryptor{}
	if ins.IsEncrypt() {
		output.Encrypt = ins.Encrypt()
	}

	if ins.IsDecrypt() {
		output.Decrypt = ins.Decrypt()
	}

	return output
}

func (app *resourceTokenLayerAdapter) structToEncryptor(
	ins structs_layers.Encryptor,
) (layers.Encryptor, error) {
	builder := app.encryptorBuilder.Create()
	if ins.Encrypt != "" {
		builder.WithEncrypt(ins.Encrypt)
	}

	if ins.Decrypt != "" {
		builder.WithDecrypt(ins.Decrypt)
	}

	return builder.Now()
}

func (app *resourceTokenLayerAdapter) signerToStruct(
	ins layers.Signer,
) structs_layers.Signer {
	output := structs_layers.Signer{}
	if ins.IsSign() {
		output.Sign = ins.Sign()
	}

	if ins.IsVote() {
		vote := app.voteToStruct(ins.Vote())
		output.Vote = &vote
	}

	if ins.IsGenerateSignerPublicKeys() {
		output.GenSignerPubKeys = ins.GenerateSignerPublicKeys()
	}

	if ins.IsHashPublicKeys() {
		output.HashPublicKeys = ins.HashPublicKeys()
	}

	if ins.IsVoteVerify() {
		voteVerify := app.voteVerifyToStruct(ins.VoteVerify())
		output.VoteVerify = &voteVerify
	}

	if ins.IsSignatureVerify() {
		sigVerify := app.signatureVerifyToStruct(ins.SignatureVerify())
		output.SignatureVerify = &sigVerify
	}

	if ins.IsBytes() {
		output.Bytes = ins.Bytes()
	}

	if ins.IsPublicKey() {
		output.IsPublicKey = ins.IsPublicKey()
	}

	return output
}

func (app *resourceTokenLayerAdapter) structToSigner(
	ins structs_layers.Signer,
) (layers.Signer, error) {
	builder := app.signerBuilder.Create()
	if ins.Sign != "" {
		builder.WithSign(ins.Sign)
	}

	if ins.Vote != nil {
		vote, err := app.structToVote(*ins.Vote)
		if err != nil {
			return nil, err
		}

		builder.WithVote(vote)
	}

	if ins.GenSignerPubKeys > 0 {
		builder.WithGenerateSignerPublicKey(ins.GenSignerPubKeys)
	}

	if ins.HashPublicKeys != "" {
		builder.WithHashPublicKeys(ins.HashPublicKeys)
	}

	if ins.VoteVerify != nil {
		voteVerify, err := app.structToVoteVerify(*ins.VoteVerify)
		if err != nil {
			return nil, err
		}

		builder.WithVoteVerify(voteVerify)
	}

	if ins.SignatureVerify != nil {
		signatureVerify, err := app.structToSignatureVerify(*ins.SignatureVerify)
		if err != nil {
			return nil, err
		}

		builder.WithSignatureVerify(signatureVerify)
	}

	if ins.Bytes != "" {
		builder.WithBytes(ins.Bytes)
	}

	if ins.IsPublicKey {
		builder.IsPublicKey()
	}

	return builder.Now()
}

func (app *resourceTokenLayerAdapter) signatureVerifyToStruct(
	ins layers.SignatureVerify,
) structs_layers.SignatureVerify {
	return structs_layers.SignatureVerify{
		Signature: ins.Signature(),
		Message:   ins.Message(),
	}
}

func (app *resourceTokenLayerAdapter) structToSignatureVerify(
	ins structs_layers.SignatureVerify,
) (layers.SignatureVerify, error) {
	return app.signatureVerifyBuilder.Create().
		WithMessage(ins.Message).
		WithSignature(ins.Signature).
		Now()
}

func (app *resourceTokenLayerAdapter) voteVerifyToStruct(
	ins layers.VoteVerify,
) structs_layers.VoteVerify {
	return structs_layers.VoteVerify{
		Vote:       ins.Vote(),
		Message:    ins.Message(),
		HashedRing: ins.HashedRing(),
	}
}

func (app *resourceTokenLayerAdapter) structToVoteVerify(
	ins structs_layers.VoteVerify,
) (layers.VoteVerify, error) {
	return app.voteVerifyBuilder.Create().
		WithVote(ins.Vote).
		WithHashedRing(ins.HashedRing).
		WithMessage(ins.Message).
		Now()
}

func (app *resourceTokenLayerAdapter) voteToStruct(
	ins layers.Vote,
) structs_layers.Vote {
	return structs_layers.Vote{
		Ring:    ins.Ring(),
		Message: ins.Message(),
	}
}

func (app *resourceTokenLayerAdapter) structToVote(
	ins structs_layers.Vote,
) (layers.Vote, error) {
	return app.voteBuilder.Create().
		WithMessage(ins.Message).
		WithRing(ins.Ring).
		Now()
}
