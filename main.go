package main

import (
	"cashback/models"
	"fmt"
)

func main() {
	levels := []models.Level{
		{
			Id:             1,
			Name:           "Bronze",
			Percent:        0.1,
			ValueToUpgrade: 30,
		},
		{
			Id:             2,
			Name:           "Prata",
			Percent:        0.2,
			ValueToUpgrade: 40,
		},
		{
			Id:             3,
			Name:           "Ouro",
			Percent:        0.3,
			ValueToUpgrade: 60,
		},
	}

	me := models.Person{
		Name:                 "Raphael",
		CurrentCashbackValue: 0,
		ValueToUpgrade:       0,
		CurrentLevel:         levels[0],
		CurrentValue:         0,
	}

	me.UpgradeLevelByValue(levels, 35)

	fmt.Println(me)
}
