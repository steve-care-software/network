package applications

import (
	admin_accounts "steve.care/network/commands/visitors/admins/domain/accounts"
	identity_accounts "steve.care/network/commands/visitors/admins/identities/domain/accounts"
	"steve.care/network/commands/visitors/admins/identities/domain/accounts/encryptors"
	"steve.care/network/commands/visitors/stencils/domain/layers"
	"steve.care/network/commands/visitors/stencils/domain/links"
	"steve.care/network/commands/visitors/stencils/domain/stacks"
)

type application struct {
	linkRepository            links.Repository
	stackAssignableBuilder    stacks.AssignableBuilder
	encryptorPublicKeyAdapter encryptors.PublicKeyAdapter
}

func createApplication(
	linkRepository links.Repository,
	stackAssignableBuilder stacks.AssignableBuilder,
	encryptorPublicKeyAdapter encryptors.PublicKeyAdapter,
) Application {
	out := application{
		linkRepository:            linkRepository,
		stackAssignableBuilder:    stackAssignableBuilder,
		encryptorPublicKeyAdapter: encryptorPublicKeyAdapter,
	}

	return &out
}

// Execute executes the program
func (app *application) Execute(
	authorized admin_accounts.Account,
	authenticated identity_accounts.Account,
	stack stacks.Stack,
) ([]byte, error) {
	root := authenticated.Root()
	return app.executeLayer(root, stack)
}

func (app *application) executeLayer(layer layers.Layer, stack stacks.Stack) ([]byte, error) {
	input := []byte{}
	last := stack.Last()
	if last.HasInput() {
		input = last.Input()
	}

	instructions := layer.Instructions()
	_, err := app.executeInstructions(instructions, input, stack)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (app *application) executeInstructions(instructions layers.Instructions, input []byte, stack stacks.Stack) ([]byte, error) {
	return nil, nil
}

func (app *application) executeInstruction(instruction layers.Instruction, input []byte, stack stacks.Stack) ([]byte, error) {
	return nil, nil
}

func (app *application) executeAssignment(assignment layers.Assignment, input []byte, stack stacks.Stack) (stacks.Assignment, error) {
	return nil, nil
}

func (app *application) executeAssignable(assignable layers.Assignable, input []byte, stack stacks.Stack) (stacks.Assignable, error) {
	if assignable.IsIdentity() {
		identity := assignable.Identity()
		return app.executeAssignableIdentity(
			identity,
			input,
			stack,
		)
	}
	return nil, nil
}

func (app *application) executeAssignableIdentity(identity layers.Identity, input []byte, stack stacks.Stack) (stacks.Assignable, error) {
	return nil, nil
}

func (app *application) executeAssignableIdentitySigner(signer layers.Signer, input []byte, stack stacks.Stack) (stacks.Assignable, error) {
	return nil, nil
}

func (app *application) executeAssignableIdentityVoter(authenticated identity_accounts.Account, voter layers.Voter, stack stacks.Stack) (stacks.Assignable, error) {
	if voter.IsVote() {
		msgValue := voter.Vote()
		_, err := app.executeValue(msgValue, stack)
		if err != nil {
			return nil, err
		}

	}

	if voter.IsVerify() {

	}

	if voter.IsPublicKey() {

	}

	return nil, nil
}

func (app *application) executeAssignableIdentityEncryptor(
	authenticated identity_accounts.Account,
	encryptor layers.Encryptor,
	stack stacks.Stack,
) (stacks.Assignable, error) {
	builder := app.stackAssignableBuilder.Create()
	if encryptor.IsDecrypt() {
		cipherValue := encryptor.Decrypt()
		cipher, err := app.executeValue(cipherValue, stack)
		if err != nil {
			return nil, err
		}

		msg, err := authenticated.Encryptor().Decrypt(cipher)
		if err != nil {
			return nil, err
		}

		builder.WithBytes(msg)
	}

	if encryptor.IsEncrypt() {
		msgValue := encryptor.Encrypt()
		msg, err := app.executeValue(msgValue, stack)
		if err != nil {
			return nil, err
		}

		cipher, err := authenticated.Encryptor().Public().Encrypt(msg)
		if err != nil {
			return nil, err
		}

		builder.WithBytes(cipher)
	}

	if encryptor.IsPublicKey() {
		publicKey := authenticated.Encryptor().Public()
		bytes := app.encryptorPublicKeyAdapter.ToBytes(publicKey)
		builder.WithBytes(bytes)
	}

	return builder.Now()
}

func (app *application) executeValue(value layers.Value, stack stacks.Stack) ([]byte, error) {
	return nil, nil
}
