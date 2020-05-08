package menu

import (
	"github.com/InspektorKot/game.git/classes"
	"github.com/manifoldco/promptui"
)

func SelectHero() classes.Character {
	prompt := promptui.Select{
		Label: "Select Character",
		Items: []string{"Knight", "Archer", "Mage"},
	}

	key, _ := selectFromMenu(prompt)

	hero := classes.Create(key)

	return hero
}
