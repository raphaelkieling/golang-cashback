package models

import "fmt"

type Person struct {
	Name                 string
	CurrentCashbackValue float64
	CurrentLevel         Level
	ValueToUpgrade       float64
	CurrentValue         float64
}

func (p *Person) getNextLevel(levels []Level) Level {
	var foundedIndex = 0
	for i, value := range levels {
		if value.Id == p.CurrentLevel.Id {
			foundedIndex = i
		}
	}

	return levels[foundedIndex+1]
}

func (p *Person) UpgradeLevelByValue(levels []Level, value float64) {
	fmt.Println("Calculando por: ", value)
	resultRemaining := p.CurrentLevel.RemainingToUpgrade(p.CurrentValue, value)
	fmt.Println(resultRemaining)

	if resultRemaining.needUpgrade {
		resultCashback := p.calcCashback(resultRemaining.remaining)
		fmt.Println("Cashback gerado de ", resultCashback, " baseado em ", resultRemaining.remaining)
		p.CurrentCashbackValue += resultCashback
		p.ValueToUpgrade = resultRemaining.remaining
		p.CurrentValue += value
		p.CurrentLevel = p.getNextLevel(levels)
		p.UpgradeLevelByValue(levels, resultRemaining.remaining)
		return
	}

	resultCashback := p.calcCashback(value)
	p.CurrentCashbackValue += resultCashback
	p.ValueToUpgrade = resultRemaining.remaining
}

func (p *Person) calcCashback(value float64) float64 {
	return value * p.CurrentLevel.Percent
}
