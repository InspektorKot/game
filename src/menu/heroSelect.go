package menu

import (
	"github.com/manifoldco/promptui"
)

func SelectHero(data []string) string {

	prompt := promptui.Select{
		Label: "Select Character",
		Items: data,
	}

	_, value := selectFromMenu(prompt)

	return value
}
