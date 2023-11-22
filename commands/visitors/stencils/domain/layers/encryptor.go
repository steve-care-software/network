package layers

import "steve.care/network/libraries/hash"

type encryptor struct {
	hash        hash.Hash
	decrypt     BytesReference
	encrypt     BytesReference
	isPublicKey bool
}

func createEncryptorWithDecrypt(
	hash hash.Hash,
	decrypt BytesReference,
) Encryptor {
	return createEncryptorInternally(hash, decrypt, nil, false)
}

func createEncryptorWithEncrypt(
	hash hash.Hash,
	encrypt BytesReference,
) Encryptor {
	return createEncryptorInternally(hash, nil, encrypt, false)
}

func createEncryptorWithIsPublicKey(
	hash hash.Hash,
) Encryptor {
	return createEncryptorInternally(hash, nil, nil, true)
}

func createEncryptorInternally(
	hash hash.Hash,
	decrypt BytesReference,
	encrypt BytesReference,
	isPublicKey bool,
) Encryptor {
	out := encryptor{
		hash:        hash,
		decrypt:     decrypt,
		encrypt:     encrypt,
		isPublicKey: isPublicKey,
	}

	return &out
}

// Hash returns the hash
func (obj *encryptor) Hash() hash.Hash {
	return obj.hash
}

// IsDecrypt returns true if there is a decrypt, false otherwise
func (obj *encryptor) IsDecrypt() bool {
	return obj.decrypt != nil
}

// Decrypt returns the decrypt, if any
func (obj *encryptor) Decrypt() BytesReference {
	return obj.decrypt
}

// IsEncrypt returns true if there is an encrypt, false otherwise
func (obj *encryptor) IsEncrypt() bool {
	return obj.encrypt != nil
}

// Encrypt returns the encrypt, if any
func (obj *encryptor) Encrypt() BytesReference {
	return obj.encrypt
}

// IsPublicKey returns true if isPublicKey, false otherwise
func (obj *encryptor) IsPublicKey() bool {
	return obj.isPublicKey
}
