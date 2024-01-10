package blocks

import "errors"

type difficultyBuilder struct {
	targetTrxAmount uint
	pMultiplier     *uint
	pBase           *uint
	pMineValue      *byte
}

func createDifficultyBuilder() DifficultyBuilder {
	out := difficultyBuilder{
		targetTrxAmount: 0,
		pMultiplier:     nil,
		pBase:           nil,
		pMineValue:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *difficultyBuilder) Create() DifficultyBuilder {
	return createDifficultyBuilder()
}

// WithTargetTrxAmount adds a target trx amount to the builder
func (app *difficultyBuilder) WithTargetTrxAmount(targetTrxAmount uint) DifficultyBuilder {
	app.targetTrxAmount = targetTrxAmount
	return app
}

// WithMultiplier adds a multiplier
func (app *difficultyBuilder) WithMultiplier(multiplier uint) DifficultyBuilder {
	app.pMultiplier = &multiplier
	return app
}

// WithBase adds a base
func (app *difficultyBuilder) WithBase(base uint) DifficultyBuilder {
	app.pBase = &base
	return app
}

// WithMineValue adds a mineValue
func (app *difficultyBuilder) WithMineValue(mineValue byte) DifficultyBuilder {
	app.pMineValue = &mineValue
	return app
}

// Now builds a new Difficulty instance
func (app *difficultyBuilder) Now() (Difficulty, error) {
	if app.targetTrxAmount <= 0 {
		return nil, errors.New("the target transaction amount must be greater than zero (0) in order to build a Difficulty instance")
	}

	if app.pMultiplier == nil {
		return nil, errors.New("the multiplier is mandatory in order to build a Difficulty instance")
	}

	if app.pBase == nil {
		return nil, errors.New("the base is mandatory in order to build a Difficulty instance")
	}

	if app.pMineValue == nil {
		return nil, errors.New("the mineValue is mandatory in order to build a Difficulty instance")
	}

	return createDifficulty(
		app.targetTrxAmount,
		*app.pMultiplier,
		*app.pBase,
		*app.pMineValue,
	), nil
}
