package blocks

import (
	"bytes"
	"testing"
)

type difficultyFunc func() Difficulty

type expectation struct {
	expectedDifficulty uint
	expectedPrefix     []byte
	prevTrxAmount      uint
	fn                 difficultyFunc
}

func TestDifficulty_Success(t *testing.T) {
	expectations := []*expectation{
		{
			expectedDifficulty: 3,
			expectedPrefix:     []byte{0, 0, 0},
			prevTrxAmount:      30,
			fn: func() Difficulty {
				targetTrxAmount := uint(20)
				multiplier := uint(2)
				base := uint(1)
				mineValue := byte(0)
				ins, err := NewDifficultyBuilder().Create().WithTargetTrxAmount(targetTrxAmount).WithMultiplier(multiplier).WithBase(base).WithMineValue(mineValue).Now()
				if err != nil {
					t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				}

				return ins
			},
		},
		{
			expectedDifficulty: 1,
			expectedPrefix:     []byte{0},
			prevTrxAmount:      10,
			fn: func() Difficulty {
				targetTrxAmount := uint(20)
				multiplier := uint(2)
				base := uint(1)
				mineValue := byte(0)
				ins, err := NewDifficultyBuilder().Create().WithTargetTrxAmount(targetTrxAmount).WithMultiplier(multiplier).WithBase(base).WithMineValue(mineValue).Now()
				if err != nil {
					t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				}

				return ins
			},
		},
	}

	for idx, oneExpectation := range expectations {
		ins := oneExpectation.fn()
		pNextDifficulty, prefix, err := ins.Fetch(oneExpectation.prevTrxAmount)
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}

		if oneExpectation.expectedDifficulty != *pNextDifficulty {
			t.Errorf("index: %d, the difficulty was expected to be %d, %d returned", idx, oneExpectation.expectedDifficulty, *pNextDifficulty)
			return
		}

		if !bytes.Equal(oneExpectation.expectedPrefix, prefix) {
			t.Errorf("the prefix was expected to be %v, %v returned", oneExpectation.expectedPrefix, prefix)
			return
		}
	}
}

func TestDifficulty_withZeroPrevTrx_returnsError(t *testing.T) {
	targetTrxAmount := uint(20)
	multiplier := uint(2)
	base := uint(1)
	mineValue := byte(0)
	ins, err := NewDifficultyBuilder().Create().WithTargetTrxAmount(targetTrxAmount).WithMultiplier(multiplier).WithBase(base).WithMineValue(mineValue).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
	}

	_, _, err = ins.Fetch(0)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestDifficulty_withZeroTargetTrxAmount_returnsError(t *testing.T) {
	targetTrxAmount := uint(0)
	multiplier := uint(2)
	base := uint(1)
	mineValue := byte(0)
	_, err := NewDifficultyBuilder().Create().WithTargetTrxAmount(targetTrxAmount).WithMultiplier(multiplier).WithBase(base).WithMineValue(mineValue).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestDifficulty_withoutTargetTrxAmount_returnsError(t *testing.T) {
	multiplier := uint(2)
	base := uint(1)
	mineValue := byte(0)
	_, err := NewDifficultyBuilder().Create().WithMultiplier(multiplier).WithBase(base).WithMineValue(mineValue).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestDifficulty_withoutMultiplier_returnsError(t *testing.T) {
	targetTrxAmount := uint(20)
	base := uint(1)
	mineValue := byte(0)
	_, err := NewDifficultyBuilder().Create().WithTargetTrxAmount(targetTrxAmount).WithBase(base).WithMineValue(mineValue).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestDifficulty_withoutBase_returnsError(t *testing.T) {
	targetTrxAmount := uint(20)
	multiplier := uint(2)
	mineValue := byte(0)
	_, err := NewDifficultyBuilder().Create().WithTargetTrxAmount(targetTrxAmount).WithMultiplier(multiplier).WithMineValue(mineValue).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestDifficulty_withoutMineValue_returnsError(t *testing.T) {
	targetTrxAmount := uint(20)
	multiplier := uint(2)
	base := uint(1)
	_, err := NewDifficultyBuilder().Create().WithTargetTrxAmount(targetTrxAmount).WithMultiplier(multiplier).WithBase(base).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
