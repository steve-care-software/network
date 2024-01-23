package blocks

import "steve.care/network/domain/hash"

type block struct {
	hash       hash.Hash
	content    Content
	difficulty uint
	minedHash  hash.Hash
	result     []byte
}

func createBlock(
	hash hash.Hash,
	content Content,
	difficulty uint,
	minedHash hash.Hash,
	result []byte,
) Block {
	out := block{
		hash:       hash,
		content:    content,
		difficulty: difficulty,
		minedHash:  minedHash,
		result:     result,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *block) Content() Content {
	return obj.content
}

// Difficulty returns the difficulty
func (obj *block) Difficulty() uint {
	return obj.difficulty
}

// MinedHash returns the minedHash
func (obj *block) MinedHash() hash.Hash {
	return obj.minedHash
}

// Result returns the result
func (obj *block) Result() []byte {
	return obj.result
}
