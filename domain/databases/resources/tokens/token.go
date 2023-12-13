package tokens

import (
	"time"

	"steve.care/network/domain/hash"
)

type token struct {
	hash      hash.Hash
	content   Content
	createdOn time.Time
}

func createToken(
	hash hash.Hash,
	content Content,
	createdOn time.Time,
) Token {
	out := token{
		hash:      hash,
		content:   content,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *token) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *token) Content() Content {
	return obj.content
}

// CreatedOn returns the creation time
func (obj *token) CreatedOn() time.Time {
	return obj.createdOn
}
