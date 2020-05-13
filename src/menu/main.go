package menu

import (
	"encoding/json"
	"fmt"
	"github.com/InspektorKot/game.git/src/managers"
	"github.com/InspektorKot/game.git/src/models"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"os"
)

func MainMenu(menuManager managers.MenuDataManager) (models.Character, string) {
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

			hero := SelectHero(managers.GetClassNameList(menuManager))
			file, _ := json.MarshalIndent(hero, "", " ")
			_ = ioutil.WriteFile(fmt.Sprintf("%s.json", buf), file, 0644)
			return hero, buf
		}
	case 1:
		{
			fmt.Println("Введите имя персонажа")
			fmt.Fscan(os.Stdin, &buf)
			if _, err := os.Stat(fmt.Sprintf("%s.json", buf)); err == nil {
				data, _ := ioutil.ReadFile(fmt.Sprintf("%s.json", buf))
				var hero models.Character
				json.Unmarshal([]byte(data), &hero)
				return hero, buf
			} else {
				fmt.Printf("Такого персонажа не существует")
				fmt.Println()
				os.Exit(1)
			}
		}
	case 2:
		os.Exit(1)
	}
	return models.Character{}, buf
}
