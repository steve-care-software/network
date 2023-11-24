package commands

import (
	"bytes"
	"errors"
	"fmt"

	"steve.care/network/applications/authenticates/accounts"
	"steve.care/network/domain/databases"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/receipts/commands/layers"
	"steve.care/network/domain/receipts/commands/links"

	identity_accounts "steve.care/network/domain/accounts"
	"steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	identity_accounts_signers "steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/receipts/commands/results"
	"steve.care/network/domain/stacks"
)

type application struct {
	accountApp                accounts.Application
	hashAdapter               hash.Adapter
	database                  databases.Database
	stackBuilder              stacks.Builder
	stackFramesBuilder        stacks.FramesBuilder
	stackFrameBuilder         stacks.FrameBuilder
	stackAssignmentsBuilder   stacks.AssignmentsBuilder
	stackAssignmentBuilder    stacks.AssignmentBuilder
	stackAssignableBuilder    stacks.AssignableBuilder
	encryptorPublicKeyAdapter encryptors.PublicKeyAdapter
	accountSignerFactory      identity_accounts_signers.Factory
	accountSignerVoteAdapter  identity_accounts_signers.VoteAdapter
	resultBuilder             results.Builder
	resultSuccessBuilder      results.SuccessBuilder
	resultFailureBuilder      results.FailureBuilder
}

func createApplication(
	accountApp accounts.Application,
	database databases.Database,
	stackBuilder stacks.Builder,
	stackFramesBuilder stacks.FramesBuilder,
	stackFrameBuilder stacks.FrameBuilder,
	stackAssignmentsBuilder stacks.AssignmentsBuilder,
	stackAssignmentBuilder stacks.AssignmentBuilder,
	stackAssignableBuilder stacks.AssignableBuilder,
	encryptorPublicKeyAdapter encryptors.PublicKeyAdapter,
	accountSignerFactory identity_accounts_signers.Factory,
	accountSignerVoteAdapter identity_accounts_signers.VoteAdapter,
	resultBuilder results.Builder,
	resultSuccessBuilder results.SuccessBuilder,
	resultFailureBuilder results.FailureBuilder,
) Application {
	out := application{
		accountApp:                accountApp,
		database:                  database,
		stackBuilder:              stackBuilder,
		stackFramesBuilder:        stackFramesBuilder,
		stackFrameBuilder:         stackFrameBuilder,
		stackAssignmentsBuilder:   stackAssignmentsBuilder,
		stackAssignmentBuilder:    stackAssignmentBuilder,
		stackAssignableBuilder:    stackAssignableBuilder,
		encryptorPublicKeyAdapter: encryptorPublicKeyAdapter,
		accountSignerFactory:      accountSignerFactory,
		accountSignerVoteAdapter:  accountSignerVoteAdapter,
		resultBuilder:             resultBuilder,
		resultSuccessBuilder:      resultSuccessBuilder,
		resultFailureBuilder:      resultFailureBuilder,
	}

	return &out
}

// Delete deletes the authenticated account
func (app *application) Delete(password []byte) error {
	return nil
}

// Update updates the authenticated account
func (app *application) Update(currentPassword []byte, newPassword []byte) error {
	return nil
}

// Exists returns true if the layer exists, false otherwise
func (app *application) Exists(hash hash.Hash) (bool, error) {
	return false, nil
}

// Execute executes a layer
func (app *application) Execute(hash hash.Hash, input []byte) (receipts.Receipt, error) {
	authenticated, err := app.accountApp.Retrieve()
	if err != nil {
		// failure
	}

	root, err := app.database.Repository().Layer().Retrieve(hash)
	if err != nil {
		return nil, err
	}

	service, err := app.database.Begin()
	if err != nil {
		return nil, err
	}

	assignable, err := app.stackAssignableBuilder.Create().
		WithBytes(input).
		Now()

	if err != nil {
		return nil, err
	}

	variable := root.Input()
	assignment, err := app.stackAssignmentBuilder.Create().
		WithName(variable).
		WithAssignable(assignable).
		Now()

	if err != nil {
		return nil, err
	}

	assignments, err := app.stackAssignmentsBuilder.Create().
		WithList([]stacks.Assignment{
			assignment,
		}).Now()

	if err != nil {
		return nil, err
	}

	frame, err := app.stackFrameBuilder.Create().
		WithAssignments(assignments).
		Now()

	if err != nil {
		return nil, err
	}

	frames, err := app.stackFramesBuilder.Create().
		WithList([]stacks.Frame{
			frame,
		}).Now()

	if err != nil {
		return nil, err
	}

	stack, err := app.stackBuilder.Create().
		WithFrames(frames).
		Now()

	if err != nil {
		return nil, err
	}

	_, err = app.executeLayer(
		service,
		authenticated,
		root,
		stack,
	)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Links returns the link based on the executed layers
func (app *application) Links(executed []hash.Hash) (links.Link, error) {
	return nil, nil
}

// Clear clears the session
func (app *application) Clear() error {
	return nil
}

func (app *application) executeLayer(
	service databases.Service,
	authenticated identity_accounts.Account,
	layer layers.Layer,
	stack stacks.Stack,
) (results.Result, error) {
	builder := app.resultBuilder.Create()
	inputVariable := layer.Input()
	inputAssignable, err := stack.Head().Fetch(inputVariable)
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
		service,
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
	outputAssignable, err := updatedStack.Head().Fetch(outputVariable)
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
	if outputIns.HasExecute() {
		command := outputIns.Execute()
		retOutputBytes, err := app.executeNativeCode(outputBytes, command)
		if err != nil {
			return nil, err
		}

		outputBytes = retOutputBytes
	}

	kind := outputIns.Kind()
	success, err := app.resultSuccessBuilder.Create().
		WithBytes(outputBytes).
		WithKind(kind).
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
	service databases.Service,
	authenticated identity_accounts.Account,
	instructions layers.Instructions,
	stack stacks.Stack,
) (stacks.Stack, results.Failure, error) {
	updatedStack := stack
	list := instructions.List()
	for idx, oneInstruction := range list {
		retStack, failure, err := app.executeInstruction(
			service,
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
	service databases.Service,
	authenticated identity_accounts.Account,
	instruction layers.Instruction,
	index uint,
	stack stacks.Stack,
) (stacks.Stack, results.Failure, error) {
	last := stack.Head()
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
		boolValue, err := stack.Head().FetchBool(variable)
		if err != nil {
			return nil, nil, err
		}

		if boolValue {
			conditionalInstructons := condition.Instructions()
			return app.executeInstructions(
				service,
				authenticated,
				conditionalInstructons,
				stack,
			)
		}

		return stack, nil, nil
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

	if instruction.IsLayer() {
		layerService := service.Layer()
		layerInstruction := instruction.Layer()
		if layerInstruction.IsSave() {

		}

		if layerInstruction.IsDelete() {
			hash := layerInstruction.Delete()
			err := layerService.Delete(hash)
			if err != nil {
				// failure
			}
		}
	}

	if instruction.IsLink() {

	}

	updatedFrameBuilder := app.stackFrameBuilder.Create()
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

	if bytesIns.IsHashBytes() {

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
		assignable, err := stack.Head().Fetch(variable)
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
		vote, err := stack.Head().FetchVote(voteVariable)
		if err != nil {
			return nil, nil, err
		}

		hashedRingVariable := voteVerify.HashedRing()
		hashedRing, err := stack.Head().FetchHashList(hashedRingVariable)
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
		sig, err := stack.Head().FetchSignature(sigVariable)
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
		pubKeys, err := stack.Head().FetchSignerPublicKeys(variable)
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
		ring, err := stack.Head().FetchSignerPublicKeys(ringVariable)
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
