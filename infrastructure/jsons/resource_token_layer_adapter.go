package jsons

import (
	"steve.care/network/domain/receipts/commands/layers"
	resources_layers "steve.care/network/domain/resources/tokens/layers"
	structs_layers "steve.care/network/infrastructure/jsons/resources/tokens/layers"
)

type resourceTokenLayerAdapter struct {
	bytesBuilder           layers.BytesBuilder
	identityBuilder        layers.IdentityBuilder
	encryptorBuilder       layers.EncryptorBuilder
	signerBuilder          layers.SignerBuilder
	signatureVerifyBuilder layers.SignatureVerifyBuilder
	voteVerifyBuilder      layers.VoteVerifyBuilder
	voteBuilder            layers.VoteBuilder
	bytesReferencesBuilder layers.BytesReferencesBuilder
	bytesReferenceBuilder  layers.BytesReferenceBuilder
}

// ToStruct converts a resource layer to struct
func (app *resourceTokenLayerAdapter) ToStruct(ins resources_layers.Layer) structs_layers.Layer {
	return structs_layers.Layer{}
}

// ToInstance converts bytes to resource layer instance
func (app *resourceTokenLayerAdapter) ToInstance(ins structs_layers.Layer) (resources_layers.Layer, error) {
	return nil, nil
}

func (app *resourceTokenLayerAdapter) bytesToStruct(
	ins layers.Bytes,
) structs_layers.Bytes {
	output := structs_layers.Bytes{}
	if ins.IsJoin() {
		join := app.bytesReferencesToStructs(ins.Join())
		output.Join = join
	}

	if ins.IsCompare() {
		compare := app.bytesReferencesToStructs(ins.Compare())
		output.Compare = compare
	}

	if ins.IsHashBytes() {
		hash := app.bytesReferenceToStruct(ins.HashBytes())
		output.Hash = &hash
	}

	return output
}

func (app *resourceTokenLayerAdapter) structToBytes(
	ins structs_layers.Bytes,
) (layers.Bytes, error) {
	builder := app.bytesBuilder.Create()
	if ins.Join != nil && len(ins.Join) > 0 {
		join, err := app.structsToBytesReferences(ins.Join)
		if err != nil {
			return nil, err
		}

		builder.WithJoin(join)
	}

	if ins.Compare != nil && len(ins.Compare) > 0 {
		compare, err := app.structsToBytesReferences(ins.Compare)
		if err != nil {
			return nil, err
		}

		builder.WithCompare(compare)
	}

	if ins.Hash != nil {
		hash, err := app.structToBytesReference(*ins.Hash)
		if err != nil {
			return nil, err
		}

		builder.WithHashBytes(hash)
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
		encrypt := app.bytesReferenceToStruct(ins.Encrypt())
		output.Encrypt = &encrypt
	}

	if ins.IsDecrypt() {
		decrypt := app.bytesReferenceToStruct(ins.Decrypt())
		output.Decrypt = &decrypt
	}

	return output
}

func (app *resourceTokenLayerAdapter) structToEncryptor(
	ins structs_layers.Encryptor,
) (layers.Encryptor, error) {
	builder := app.encryptorBuilder.Create()
	if ins.Encrypt != nil {
		encrypt, err := app.structToBytesReference(*ins.Encrypt)
		if err != nil {
			return nil, err
		}

		builder.WithEncrypt(encrypt)
	}

	if ins.Decrypt != nil {
		decrypt, err := app.structToBytesReference(*ins.Decrypt)
		if err != nil {
			return nil, err
		}

		builder.WithDecrypt(decrypt)
	}

	return builder.Now()
}

func (app *resourceTokenLayerAdapter) signerToStruct(
	ins layers.Signer,
) structs_layers.Signer {
	output := structs_layers.Signer{}
	if ins.IsSign() {
		sign := app.bytesReferenceToStruct(ins.Sign())
		output.Sign = &sign
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
	if ins.Sign != nil {
		sign, err := app.structToBytesReference(*ins.Sign)
		if err != nil {
			return nil, err
		}

		builder.WithSign(sign)
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
	message := app.bytesReferenceToStruct(ins.Message())
	return structs_layers.SignatureVerify{
		Signature: ins.Signature(),
		Message:   message,
	}
}

func (app *resourceTokenLayerAdapter) structToSignatureVerify(
	ins structs_layers.SignatureVerify,
) (layers.SignatureVerify, error) {
	message, err := app.structToBytesReference(ins.Message)
	if err != nil {
		return nil, err
	}

	return app.signatureVerifyBuilder.Create().
		WithMessage(message).
		WithSignature(ins.Signature).
		Now()
}

func (app *resourceTokenLayerAdapter) voteVerifyToStruct(
	ins layers.VoteVerify,
) structs_layers.VoteVerify {
	message := app.bytesReferenceToStruct(ins.Message())
	return structs_layers.VoteVerify{
		Vote:       ins.Vote(),
		Message:    message,
		HashedRing: ins.HashedRing(),
	}
}

func (app *resourceTokenLayerAdapter) structToVoteVerify(
	ins structs_layers.VoteVerify,
) (layers.VoteVerify, error) {
	message, err := app.structToBytesReference(ins.Message)
	if err != nil {
		return nil, err
	}

	return app.voteVerifyBuilder.Create().
		WithVote(ins.Vote).
		WithHashedRing(ins.HashedRing).
		WithMessage(message).
		Now()
}

func (app *resourceTokenLayerAdapter) voteToStruct(
	ins layers.Vote,
) structs_layers.Vote {
	message := app.bytesReferenceToStruct(ins.Message())
	return structs_layers.Vote{
		Ring:    ins.Ring(),
		Message: message,
	}
}

func (app *resourceTokenLayerAdapter) structToVote(
	ins structs_layers.Vote,
) (layers.Vote, error) {
	message, err := app.structToBytesReference(ins.Message)
	if err != nil {
		return nil, err
	}

	return app.voteBuilder.Create().
		WithMessage(message).
		WithRing(ins.Ring).
		Now()
}

func (app *resourceTokenLayerAdapter) bytesReferencesToStructs(
	ins layers.BytesReferences,
) []structs_layers.BytesReference {
	list := ins.List()
	output := []structs_layers.BytesReference{}
	for _, oneIns := range list {
		str := app.bytesReferenceToStruct(oneIns)
		output = append(output, str)
	}

	return output
}

func (app *resourceTokenLayerAdapter) structsToBytesReferences(
	list []structs_layers.BytesReference,
) (layers.BytesReferences, error) {
	output := []layers.BytesReference{}
	for _, oneIns := range list {
		ins, err := app.structToBytesReference(oneIns)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}
	return app.bytesReferencesBuilder.Create().
		WithList(output).
		Now()
}

func (app *resourceTokenLayerAdapter) bytesReferenceToStruct(
	ins layers.BytesReference,
) structs_layers.BytesReference {
	output := structs_layers.BytesReference{}
	if ins.IsVariable() {
		output.Variable = ins.Variable()
	}

	if ins.IsBytes() {
		output.Bytes = ins.Bytes()
	}

	return output
}

func (app *resourceTokenLayerAdapter) structToBytesReference(
	ins structs_layers.BytesReference,
) (layers.BytesReference, error) {
	builder := app.bytesReferenceBuilder.Create()
	if ins.Variable != "" {
		builder.WithVariable(ins.Variable)
	}

	if ins.Bytes != nil && len(ins.Bytes) > 0 {
		builder.WithBytes(ins.Bytes)
	}

	return builder.Now()
}
