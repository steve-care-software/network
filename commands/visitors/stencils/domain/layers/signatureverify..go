package layers

import "steve.care/network/libraries/hash"

type signatureVerify struct {
	hash      hash.Hash
	signature string
	message   BytesReference
}

func createSignatureVerify(
	hash hash.Hash,
	signature string,
	message BytesReference,
) SignatureVerify {
	out := signatureVerify{
		hash:      hash,
		signature: signature,
		message:   message,
	}

	return &out
}

// Hash returns the hash
func (obj *signatureVerify) Hash() hash.Hash {
	return obj.hash
}

// Signature returns the signature
func (obj *signatureVerify) Signature() string {
	return obj.signature
}

// Message returns the message
func (obj *signatureVerify) Message() BytesReference {
	return obj.message
}
