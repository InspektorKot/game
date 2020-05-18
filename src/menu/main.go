package menu

import (
	"fmt"
	"github.com/InspektorKot/game.git/src/managers"
	"github.com/InspektorKot/game.git/src/models"
	"github.com/manifoldco/promptui"
	"os"
)

func MainMenu(menuManager managers.MenuDataManager, playerManager managers.PlayerManager) models.Player {
	prompt := promptui.Select{
		Label: "Главное меню",
		Items: []string{"Новая игра", "Загрузить игру", "Выйти"},
	}

	key, _ := selectFromMenu(prompt)

	var buf string
	switch key {
	case 0:
		{
			fmt.Println("Введите имя персонажа")
			fmt.Fscan(os.Stdin, &buf)

			selectedClass := SelectHero(managers.GetClassNameList(menuManager))

			hero := managers.NewPlayer(playerManager, selectedClass, buf)

			return *hero
		}
	case 1:
		{
			fmt.Println("Введите имя персонажа")
			fmt.Fscan(os.Stdin, &buf)
			hero := managers.Load(playerManager, buf)

			return *hero
		}
	case 2:
		os.Exit(1)
	}
	return models.Player{}
}
