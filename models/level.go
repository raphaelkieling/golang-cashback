package models

import (
	"math"
)

type Level struct {
	Id             int
	Name           string
	ValueToUpgrade float64
	Percent        float64
}

type LevelRemainingResult struct {
	valueToUseInNext float64
	remaining        float64
	needUpgrade      bool
}

func (l *Level) RemainingToUpgrade(fromValue float64, value float64) LevelRemainingResult {
	remainingCurrent := fromValue - l.ValueToUpgrade
	needUpgrade := false

	if remainingCurrent > 0 {
		needUpgrade = true
	} else {
		needUpgrade = false
	}

	result := LevelRemainingResult{
		remaining:        math.Abs(remainingCurrent),
		needUpgrade:      needUpgrade,
		valueToUseInNext: value - l.ValueToUpgrade,
	}

	return result
}
