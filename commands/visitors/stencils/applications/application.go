package applications

import (
	"bytes"
	"errors"
	"fmt"

	"steve.care/network/libraries/hash"

	admin_accounts "steve.care/network/commands/visitors/admins/domain/accounts"
	identity_accounts "steve.care/network/commands/visitors/admins/identities/domain/accounts"
	"steve.care/network/commands/visitors/admins/identities/domain/accounts/encryptors"
	"steve.care/network/commands/visitors/admins/identities/domain/accounts/signers"
	identity_accounts_signers "steve.care/network/commands/visitors/admins/identities/domain/accounts/signers"
	"steve.care/network/commands/visitors/stencils/domain/layers"
	"steve.care/network/commands/visitors/stencils/domain/links"
	"steve.care/network/commands/visitors/stencils/domain/results"
	"steve.care/network/commands/visitors/stencils/domain/stacks"
)

type application struct {
	hashAdapter               hash.Adapter
	layerRepository           layers.Repository
	linkRepository            links.Repository
	stackBuilder              stacks.Builder
	stackFramesBuilder        stacks.FramesBuilder
	stackFrameBuilder         stacks.FrameBuilder
	stackInstructionsBuilder  stacks.InstructionsBuilder
	stackInstructionBuilder   stacks.InstructionBuilder
	stackAssignmentsBuilder   stacks.AssignmentsBuilder
	stackAssignmentBuilder    stacks.AssignmentBuilder
	stackAssignableBuilder    stacks.AssignableBuilder
	encryptorPublicKeyAdapter encryptors.PublicKeyAdapter
	accountSignerFactory      identity_accounts_signers.Factory
	accountSignerVoteAdapter  identity_accounts_signers.VoteAdapter
	resultBuilder             results.Builder
	resultSuccessBuilder      results.SuccessBuilder
	resultActionBuilder       results.ActionBuilder
	resultFailureBuilder      results.FailureBuilder
}

func createApplication(
	layerRepository layers.Repository,
	linkRepository links.Repository,
	stackBuilder stacks.Builder,
	stackFramesBuilder stacks.FramesBuilder,
	stackFrameBuilder stacks.FrameBuilder,
	stackInstructionsBuilder stacks.InstructionsBuilder,
	stackInstructionBuilder stacks.InstructionBuilder,
	stackAssignmentsBuilder stacks.AssignmentsBuilder,
	stackAssignmentBuilder stacks.AssignmentBuilder,
	stackAssignableBuilder stacks.AssignableBuilder,
	encryptorPublicKeyAdapter encryptors.PublicKeyAdapter,
	accountSignerFactory identity_accounts_signers.Factory,
	accountSignerVoteAdapter identity_accounts_signers.VoteAdapter,
	resultBuilder results.Builder,
	resultSuccessBuilder results.SuccessBuilder,
	resultActionBuilder results.ActionBuilder,
	resultFailureBuilder results.FailureBuilder,
) Application {
	out := application{
		layerRepository:           layerRepository,
		linkRepository:            linkRepository,
		stackBuilder:              stackBuilder,
		stackFramesBuilder:        stackFramesBuilder,
		stackFrameBuilder:         stackFrameBuilder,
		stackInstructionsBuilder:  stackInstructionsBuilder,
		stackInstructionBuilder:   stackInstructionBuilder,
		stackAssignmentsBuilder:   stackAssignmentsBuilder,
		stackAssignmentBuilder:    stackAssignmentBuilder,
		stackAssignableBuilder:    stackAssignableBuilder,
		encryptorPublicKeyAdapter: encryptorPublicKeyAdapter,
		accountSignerFactory:      accountSignerFactory,
		accountSignerVoteAdapter:  accountSignerVoteAdapter,
		resultBuilder:             resultBuilder,
		resultSuccessBuilder:      resultSuccessBuilder,
		resultActionBuilder:       resultActionBuilder,
		resultFailureBuilder:      resultFailureBuilder,
	}

	return &out
}

// Execute executes the program
func (app *application) Execute(
	authorized admin_accounts.Account,
	authenticated identity_accounts.Account,
	stack stacks.Stack,
) (results.Result, error) {
	path := authenticated.Root()
	root, err := app.layerRepository.Retrieve(path)
	if err != nil {
		return nil, err
	}

	return app.executeLayer(
		authenticated,
		root,
		stack,
	)
}

func (app *application) executeLayer(
	authenticated identity_accounts.Account,
	layer layers.Layer,
	stack stacks.Stack,
) (results.Result, error) {
	builder := app.resultBuilder.Create()
	inputVariable := layer.Input()
	inputAssignable, err := stack.Last().Fetch(inputVariable)
	if err != nil {
		failure, err := app.resultFailureBuilder.Create().
			WithCode(results.InputNotFoundError).
			Now()

		if err != nil {
			return nil, err
		}

		return builder.WithFailure(failure).
			Now()
	}

	if !inputAssignable.IsBytes() {
		failure, err := app.resultFailureBuilder.Create().
			WithCode(results.InputNotBytesError).
			Now()

		if err != nil {
			return nil, err
		}

		return builder.WithFailure(failure).
			Now()
	}

	instructions := layer.Instructions()
	updatedStack, failure, err := app.executeInstructions(
		authenticated,
		instructions,
		stack,
	)

	if err != nil {
		return nil, err
	}

	if failure != nil {
		return builder.WithFailure(failure).
			Now()
	}

	outputIns := layer.Output()
	outputVariable := outputIns.Variable()
	outputAssignable, err := updatedStack.Last().Fetch(outputVariable)
	if err != nil {
		failure, err := app.resultFailureBuilder.Create().
			WithCode(results.OutputNotFoundError).
			Now()

		if err != nil {
			return nil, err
		}

		return builder.WithFailure(failure).
			Now()
	}

	if !outputAssignable.IsBytes() {
		failure, err := app.resultFailureBuilder.Create().
			WithCode(results.OutputNotBytesError).
			Now()

		if err != nil {
			return nil, err
		}

		return builder.WithFailure(failure).
			Now()
	}

	outputBytes := outputAssignable.Bytes()
	kind := outputIns.Kind()
	if kind.HasExecute() {
		command := kind.Execute()
		retOutputBytes, err := app.executeNativeCode(outputBytes, command)
		if err != nil {
			return nil, err
		}

		outputBytes = retOutputBytes
	}

	actionBuilder := app.resultActionBuilder.Create()
	if kind.IsContinue() {
		actionBuilder.IsContinue()
	}

	if kind.IsPrompt() {
		actionBuilder.IsPrompt()
	}

	action, err := actionBuilder.Now()
	if err != nil {
		return nil, err
	}

	success, err := app.resultSuccessBuilder.Create().
		WithBytes(outputBytes).
		WithAction(action).
		Now()

	if err != nil {
		return nil, err
	}

	return builder.WithSuccess(success).
		Now()
}

func (app *application) executeNativeCode(
	codeBytes []byte,
	command string,
) ([]byte, error) {
	return nil, nil
}

func (app *application) executeInstructions(
	authenticated identity_accounts.Account,
	instructions layers.Instructions,
	stack stacks.Stack,
) (stacks.Stack, results.Failure, error) {
	updatedStack := stack
	list := instructions.List()
	for idx, oneInstruction := range list {
		retStack, failure, err := app.executeInstruction(
			authenticated,
			oneInstruction,
			uint(idx),
			updatedStack,
		)

		if err != nil {
			return nil, nil, err
		}

		if failure != nil {
			return nil, failure, nil
		}

		// stop
		if retStack == nil {
			return updatedStack, nil, nil
		}

		updatedStack = retStack
	}

	return updatedStack, nil, nil
}

func (app *application) executeInstruction(
	authenticated identity_accounts.Account,
	instruction layers.Instruction,
	index uint,
	stack stacks.Stack,
) (stacks.Stack, results.Failure, error) {
	currentFrameInstructionList := []stacks.Instruction{}
	last := stack.Last()
	if last.HasInstructions() {
		currentFrameInstructionList = last.Instructions().List()
	}

	currentFrameAssignments := []stacks.Assignment{}
	if last.HasAssignments() {
		currentFrameAssignments = last.Assignments().List()
	}

	if instruction.IsStop() {
		return nil, nil, nil
	}

	if instruction.IsRaiseError() {
		code := instruction.RaiseError()
		failure, err := app.resultFailureBuilder.Create().
			WithCode(code).
			IsRaisedInLayer().
			Now()

		if err != nil {
			return nil, nil, err
		}

		return nil, failure, nil
	}

	if instruction.IsCondition() {
		condition := instruction.Condition()
		variable := condition.Variable()
		boolValue, err := stack.Last().FetchBool(variable)
		if err != nil {
			return nil, nil, err
		}

		if boolValue {
			conditionalInstructons := condition.Instructions()
			return app.executeInstructions(
				authenticated,
				conditionalInstructons,
				stack,
			)
		}

		return stack, nil, nil
	}

	if instruction.IsSave() {
		newLayer := instruction.Save()
		instruction, err := app.stackInstructionBuilder.Create().
			WithSave(newLayer).
			Now()

		if err != nil {
			return nil, nil, err
		}

		currentFrameInstructionList = append(
			currentFrameInstructionList,
			instruction,
		)
	}

	if instruction.IsAssignment() {
		assignment := instruction.Assignment()
		stackAssignment, failure, err := app.executeAssignment(authenticated, assignment, stack)
		if err != nil {
			return nil, nil, err
		}

		if failure != nil {
			return nil, failure, nil
		}

		currentFrameAssignments = append(currentFrameAssignments, stackAssignment)
	}

	updatedFrameBuilder := app.stackFrameBuilder.Create()
	if len(currentFrameInstructionList) > 0 {
		updatedStackInstructions, err := app.stackInstructionsBuilder.Create().
			WithList(currentFrameInstructionList).
			Now()

		if err != nil {
			return nil, nil, err
		}

		updatedFrameBuilder.WithInstructions(updatedStackInstructions)
	}

	if len(currentFrameAssignments) > 0 {
		updatedStackAssignments, err := app.stackAssignmentsBuilder.Create().
			WithList(currentFrameAssignments).
			Now()

		if err != nil {
			return nil, nil, err
		}

		updatedFrameBuilder.WithAssignments(updatedStackAssignments)
	}

	updatedFrame, err := updatedFrameBuilder.Now()
	if err != nil {
		return nil, nil, err
	}

	body := stack.Body().List()
	framesList := append(body, updatedFrame)
	updatedFrames, err := app.stackFramesBuilder.Create().
		WithList(framesList).
		Now()

	if err != nil {
		return nil, nil, err
	}

	updatedStack, err := app.stackBuilder.Create().
		WithFrames(updatedFrames).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return updatedStack, nil, nil
}

func (app *application) executeAssignment(
	authenticated identity_accounts.Account,
	assignment layers.Assignment,
	stack stacks.Stack,
) (stacks.Assignment, results.Failure, error) {
	name := assignment.Name()
	assignable := assignment.Assignable()
	stackAssignable, stackFailure, err := app.executeAssignable(authenticated, assignable, stack)
	if err != nil {
		return nil, nil, err
	}

	if stackFailure != nil {
		return nil, stackFailure, nil
	}

	stackAssignment, err := app.stackAssignmentBuilder.Create().
		WithName(name).
		WithAssignable(stackAssignable).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return stackAssignment, nil, nil
}

func (app *application) executeAssignable(
	authenticated identity_accounts.Account,
	assignable layers.Assignable,
	stack stacks.Stack,
) (stacks.Assignable, results.Failure, error) {
	if assignable.IsIdentity() {
		identity := assignable.Identity()
		return app.executeAssignableIdentity(
			authenticated,
			identity,
			stack,
		)
	}

	bytesIns := assignable.Bytes()
	return app.executeAssignableBytes(
		authenticated,
		bytesIns,
		stack,
	)
}

func (app *application) executeAssignableBytes(
	authenticated identity_accounts.Account,
	bytesIns layers.Bytes,
	stack stacks.Stack,
) (stacks.Assignable, results.Failure, error) {
	builder := app.stackAssignableBuilder.Create()
	if bytesIns.IsJoin() {
		joined := []byte{}
		refList := bytesIns.Join().List()
		for _, oneRef := range refList {
			value, err := app.executeBytesReference(oneRef, stack)
			if err != nil {
				return nil, nil, err
			}

			joined = append(joined, value...)
		}

		builder.WithBytes(joined)
	}

	if bytesIns.IsCompare() {
		isEqual := true
		last := []byte{}
		refList := bytesIns.Compare().List()
		for _, oneRef := range refList {
			value, err := app.executeBytesReference(oneRef, stack)
			if err != nil {
				return nil, nil, err
			}

			if len(last) <= 0 {
				last = value
				continue
			}

			if !bytes.Equal(last, last) {
				isEqual = false
			}
		}

		builder.WithBool(isEqual)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}

func (app *application) executeAssignableIdentity(
	authenticated identity_accounts.Account,
	identity layers.Identity,
	stack stacks.Stack,
) (stacks.Assignable, results.Failure, error) {
	if identity.IsEncryptor() {
		encryptor := identity.Encryptor()
		return app.executeAssignableIdentityEncryptor(
			authenticated,
			encryptor,
			stack,
		)
	}

	signer := identity.Signer()
	return app.executeAssignableIdentitySigner(
		authenticated,
		signer,
		stack,
	)
}

func (app *application) executeAssignableIdentitySigner(
	authenticated identity_accounts.Account,
	signer layers.Signer,
	stack stacks.Stack,
) (stacks.Assignable, results.Failure, error) {
	builder := app.stackAssignableBuilder.Create()
	if signer.IsBytes() {
		variable := signer.Bytes()
		assignable, err := stack.Last().Fetch(variable)
		if err != nil {
			return nil, nil, err
		}

		if assignable.IsSignature() {
			bytes, err := assignable.Signature().Bytes()
			if err != nil {
				return nil, nil, err
			}

			ins, err := builder.WithBytes(bytes).
				Now()

			if err != nil {
				return nil, nil, err
			}

			return ins, nil, nil
		}

		if assignable.IsVote() {
			bytes, err := assignable.Vote().Bytes()
			if err != nil {
				return nil, nil, err
			}

			ins, err := builder.WithBytes(bytes).
				Now()

			if err != nil {
				return nil, nil, err
			}

			return ins, nil, nil
		}

		if assignable.IsSignerPublicKey() {
			bytes, err := assignable.SignerPublicKey().Bytes()
			if err != nil {
				return nil, nil, err
			}

			ins, err := builder.WithBytes(bytes).
				Now()

			if err != nil {
				return nil, nil, err
			}

			return ins, nil, nil
		}

		str := fmt.Sprintf("the variable (name: %s) does NOT hold a value that can be converted to bytes (signature, vote, publicKeys)", variable)
		return nil, nil, errors.New(str)
	}

	if signer.IsPublicKey() {
		pubKey := authenticated.Signer().PublicKey()
		builder.WithSignerPublicKey(pubKey)
	}

	if signer.IsVoteVerify() {
		voteVerify := signer.VoteVerify()
		voteVariable := voteVerify.Vote()
		vote, err := stack.Last().FetchVote(voteVariable)
		if err != nil {
			return nil, nil, err
		}

		hashedRingVariable := voteVerify.HashedRing()
		hashedRing, err := stack.Last().FetchHashList(hashedRingVariable)
		if err != nil {
			return nil, nil, err
		}

		msgRef := voteVerify.Message()
		msg, err := app.executeBytesReference(msgRef, stack)
		if err != nil {
			return nil, nil, err
		}

		isValid, err := app.accountSignerVoteAdapter.ToVerification(vote, msg, hashedRing)
		if err != nil {
			return nil, nil, err
		}

		builder.WithBool(isValid)
	}

	if signer.IsSignatureVerify() {
		signatureVerify := signer.SignatureVerify()
		sigVariable := signatureVerify.Signature()
		sig, err := stack.Last().FetchSignature(sigVariable)
		if err != nil {
			return nil, nil, err
		}

		msgRef := signatureVerify.Message()
		msg, err := app.executeBytesReference(msgRef, stack)
		if err != nil {
			return nil, nil, err
		}

		sigPubKey, err := sig.PublicKey(msg)
		if err != nil {
			return nil, nil, err
		}

		isValid := authenticated.Signer().PublicKey().Equals(sigPubKey) && sig.Verify()
		builder.WithBool(isValid)
	}

	if signer.IsGenerateSignerPublicKeys() {
		list := []signers.PublicKey{}
		amount := signer.GenerateSignerPublicKeys()
		castedAmount := int(amount)
		for i := 0; i < castedAmount; i++ {
			signer := app.accountSignerFactory.Create()
			pubKey := signer.PublicKey()
			list = append(list, pubKey)
		}

		builder.WithSignerPublicKeys(list)
	}

	if signer.IsHashPublicKeys() {
		variable := signer.HashPublicKeys()
		pubKeys, err := stack.Last().FetchSignerPublicKeys(variable)
		if err != nil {
			return nil, nil, err
		}

		hashList := []hash.Hash{}
		for _, onePubKey := range pubKeys {
			bytes, err := onePubKey.Bytes()
			if err != nil {
				return nil, nil, err
			}

			pHash, err := app.hashAdapter.FromBytes(bytes)
			if err != nil {
				return nil, nil, err
			}

			hashList = append(hashList, *pHash)
		}

		builder.WithHashList(hashList)
	}

	if signer.IsVote() {
		signerVote := signer.Vote()
		ringVariable := signerVote.Ring()
		ring, err := stack.Last().FetchSignerPublicKeys(ringVariable)
		if err != nil {
			return nil, nil, err
		}

		msgRef := signerVote.Message()
		msg, err := app.executeBytesReference(msgRef, stack)
		if err != nil {
			return nil, nil, err
		}

		vote, err := authenticated.Signer().Vote(msg, ring)
		if err != nil {
			return nil, nil, err
		}

		builder.WithVote(vote)
	}

	if signer.IsSign() {
		msgRef := signer.Sign()
		msg, err := app.executeBytesReference(msgRef, stack)
		if err != nil {
			return nil, nil, err
		}

		sig, err := authenticated.Signer().Sign(msg)
		if err != nil {
			return nil, nil, err
		}

		builder.WithSignature(sig)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}

func (app *application) executeAssignableIdentityEncryptor(
	authenticated identity_accounts.Account,
	encryptor layers.Encryptor,
	stack stacks.Stack,
) (stacks.Assignable, results.Failure, error) {
	builder := app.stackAssignableBuilder.Create()
	if encryptor.IsDecrypt() {
		cipherRef := encryptor.Decrypt()
		cipher, err := app.executeBytesReference(cipherRef, stack)
		if err != nil {
			return nil, nil, err
		}

		msg, err := authenticated.Encryptor().Decrypt(cipher)
		if err != nil {
			return nil, nil, err
		}

		builder.WithBytes(msg)
	}

	if encryptor.IsEncrypt() {
		msgRef := encryptor.Encrypt()
		msg, err := app.executeBytesReference(msgRef, stack)
		if err != nil {
			return nil, nil, err
		}

		cipher, err := authenticated.Encryptor().Public().Encrypt(msg)
		if err != nil {
			return nil, nil, err
		}

		builder.WithBytes(cipher)
	}

	if encryptor.IsPublicKey() {
		publicKey := authenticated.Encryptor().Public()
		bytes := app.encryptorPublicKeyAdapter.ToBytes(publicKey)
		builder.WithBytes(bytes)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}

func (app *application) executeBytesReference(
	bytesRef layers.BytesReference,
	stack stacks.Stack,
) ([]byte, error) {
	return nil, nil
}
