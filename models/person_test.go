package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCashback(t *testing.T) {
	levels := []Level{
		{
			Id:             1,
			Name:           "a",
			Percent:        0.01,
			ValueToUpgrade: 30.0,
		},
		{
			Id:             1,
			Name:           "b",
			Percent:        0.02,
			ValueToUpgrade: 40.0,
		},
	}

	person := Person{
		Name:                 "c",
		CurrentCashbackValue: 0,
		CurrentLevel:         levels[0],
		ValueToUpgrade:       0,
	}

	person.UpgradeLevelByValue(levels, 25.0)

	assert.Equal(t, person.CurrentLevel.Id, 1)
	assert.Equal(t, person.CurrentCashbackValue, 0.25)
}

func TestGenerateCashbackAndUpgradeOneLevel(t *testing.T) {
	levels := []Level{
		{
			Id:             1,
			Name:           "a",
			Percent:        0.01,
			ValueToUpgrade: 30.0,
		},
		{
			Id:             1,
			Name:           "b",
			Percent:        0.02,
			ValueToUpgrade: 40.0,
		},
	}

	person := Person{
		Name:                 "c",
		CurrentCashbackValue: 0,
		CurrentLevel:         levels[0],
		ValueToUpgrade:       0,
	}

	person.UpgradeLevelByValue(levels, 43)

	assert.Equal(t, person.CurrentLevel.Id, 2)
	assert.Equal(t, person.CurrentCashbackValue, 37.4)
}

/**
*	Se eu pagar 43 reais 30 são computador em bronze e
*	13 em prata
*              bronze           prata
*	|            |               |          |
*	|---------------------------------------|
*                30 (10%)        40 (10%)
 */
