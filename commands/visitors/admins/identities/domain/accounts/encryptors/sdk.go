package encryptors

// Factory represents an encryptor factory
type Factory interface {
	Create() Encryptor
}

// Encryptor represents an encryptor
type Encryptor interface {
	Public() PublicKey
	Decrypt(cipher []byte) ([]byte, error)
	Bytes() []byte
}

// PublicKey represents a public key
type PublicKey interface {
	Encrypt(msg []byte) ([]byte, error)
	Bytes() []byte
}
