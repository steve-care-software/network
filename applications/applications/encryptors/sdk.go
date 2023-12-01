package encryptors

// Application represents the encryptor application
type Application interface {
	Encrypt(bytes []byte, password []byte) ([]byte, error)
	Decrypt(cipher []byte, password []byte) ([]byte, error)
}
