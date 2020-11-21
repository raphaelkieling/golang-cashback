package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotNeedUpgrade(t *testing.T) {
	bronze := Level{
		Id:             1,
		Name:           "",
		Percent:        0.1,
		ValueToUpgrade: 30,
	}

	result := bronze.RemainingToUpgrade(10, 15)

	assert.Equal(t, result.needUpgrade, false)
	assert.Equal(t, result.valueToUseInNext, 0)
	assert.Equal(t, result.remaining, 5)
}

// func TestNeedReturn29Remaining(t *testing.T) {
// 	bronze := Level{
// 		Id:             1,
// 		Name:           "",
// 		Percent:        0.1,
// 		ValueToUpgrade: 30,
// 	}

// 	result := bronze.RemainingToUpgrade(1)

// 	assert.NotEqual(t, result.remaining, 29)
// }

// func TestNeedReturnNeedUpgrade(t *testing.T) {
// 	bronze := Level{
// 		Id:             1,
// 		Name:           "",
// 		Percent:        0.1,
// 		ValueToUpgrade: 30,
// 	}

// 	result := bronze.RemainingToUpgrade(31)

// 	assert.Equal(t, result.needUpgrade, true)
// 	assert.Equal(t, result.remaining, float64(1))
// }
