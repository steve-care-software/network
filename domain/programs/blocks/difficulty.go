package blocks

import "errors"

type difficulty struct {
	targetTrxAmount uint
	multiplier      uint
	base            uint
	mineValue       byte
}

func createDifficulty(
	targetTrxAmount uint,
	multiplier uint,
	base uint,
	mineValue byte,
) Difficulty {
	out := difficulty{
		targetTrxAmount: targetTrxAmount,
		multiplier:      multiplier,
		base:            base,
		mineValue:       mineValue,
	}

	return &out
}

// Fetch fetches the difficulty
func (obj *difficulty) Fetch(prevTrxAmount uint) (*uint, []byte, error) {
	if prevTrxAmount <= 0 {
		return nil, nil, errors.New("the previous transaction amount must be greater than zero (0) in order to fetch its difficulty and prefix bytes")
	}

	// calculate the difficulty:
	amount := uint(prevTrxAmount / obj.targetTrxAmount)
	difficulty := int((obj.multiplier * amount) + obj.base)

	// build the prefix:
	prefix := []byte{}
	for i := 0; i < difficulty; i++ {
		prefix = append(prefix, obj.mineValue)
	}

	// return the values:
	casted := uint(difficulty)
	return &casted, prefix, nil
}
