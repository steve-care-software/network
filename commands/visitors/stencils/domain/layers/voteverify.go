package layers

import "steve.care/network/libraries/hash"

type voteVerify struct {
	hash       hash.Hash
	vote       string
	message    BytesReference
	hashedRing string
}

func createVoteVerify(
	hash hash.Hash,
	vote string,
	message BytesReference,
	hashedRing string,
) VoteVerify {
	out := voteVerify{
		hash:       hash,
		vote:       vote,
		message:    message,
		hashedRing: hashedRing,
	}

	return &out
}

// Hash returns the hash
func (obj *voteVerify) Hash() hash.Hash {
	return obj.hash
}

// Vote returns the vote
func (obj *voteVerify) Vote() string {
	return obj.vote
}

// Message returns the message
func (obj *voteVerify) Message() BytesReference {
	return obj.message
}

// HashedRing returns the hashed ring
func (obj *voteVerify) HashedRing() string {
	return obj.hashedRing
}
