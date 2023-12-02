package accounts

// Account represents an account
type Account struct {
	Username  string `json:"username"`
	Encryptor []byte `json:"encryptor"`
	Signer    []byte `json:"signer"`
}
