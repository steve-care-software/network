package blocks

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewDifficultyBuilder creates a new difficulty builder
func NewDifficultyBuilder() DifficultyBuilder {
	return createDifficultyBuilder()
}

// NewContentBuilder creates a content builder
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createContentBuilder(hashAdapter)
}

// Compute computes an hash
func Compute(msg []byte, result []byte) (*hash.Hash, error) {
	return hash.NewAdapter().FromMultiBytes([][]byte{
		result,
		msg,
	})
}

// Builder represents the block builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithResult(result []byte) Builder
	Now() (Block, error)
}

// Block representa a block
type Block interface {
	Hash() hash.Hash
	Content() Content
	Difficulty() uint
	MinedHash() hash.Hash
	Result() []byte
}

// DifficultyBuilder represents the difficulty builder
type DifficultyBuilder interface {
	Create() DifficultyBuilder
	WithTargetTrxAmount(targetTrxAmount uint) DifficultyBuilder
	WithMultiplier(multiplier uint) DifficultyBuilder
	WithBase(base uint) DifficultyBuilder
	WithMineValue(mineValue byte) DifficultyBuilder
	Now() (Difficulty, error)
}

// Difficulty represents the difficulty
type Difficulty interface {
	Fetch(prevTrxAmount uint) (*uint, []byte, error)
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithTransactions(transactions transactions.Transactions) ContentBuilder
	WithParent(parent hash.Hash) ContentBuilder
	Now() (Content, error)
}

// Content represents the block content
type Content interface {
	Hash() hash.Hash
	Transactions() transactions.Transactions
	HasParent() bool
	Parent() hash.Hash
}

// Repository represents a block repository
type Repository interface {
	Retrieve(hash hash.Hash) (Block, error)
	RetrieveByParent(parent hash.Hash) (Block, error)
	RetrieveHeadByProgram(program hash.Hash) (Block, error)
}

// Service represents a block service
type Service interface {
	Insert(block Block) error
	Delete(hash hash.Hash) error
}
