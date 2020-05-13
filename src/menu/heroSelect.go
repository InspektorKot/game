package menu

import (
	"github.com/InspektorKot/game.git/src/models"
	"github.com/manifoldco/promptui"
)

func SelectHero(data []string) models.Character {


	prompt := promptui.Select{
		Label: "Select Character",
		Items: data,
	}

	key, _ := selectFromMenu(prompt)

	hero := models.Create(key)

	return hero
}
