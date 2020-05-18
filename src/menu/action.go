package menu

import (
	"github.com/InspektorKot/game.git/src/managers"
	"github.com/InspektorKot/game.git/src/models"
	"github.com/manifoldco/promptui"
)

func SelectAction(menuManager managers.MenuDataManager, player *models.Player, enemy *models.Enemy) int {
	prompt := promptui.Select{
		Label: "Выберите действие",
		Items: []string{
			"Атаковать",
			"Сбежать",
			//"Использовать инвентарь",
			//"Использовать навыки",
		},
	}

	key, _ := selectFromMenu(prompt)

	switch key {
	case 0:
		player.Attack(enemy)
	case 1:
		return 1
	case 2:
		{
			return 0
			//Get
			//SelectItem(character, enemy)
		}
	case 3:
		{
			//SelectSkill(character, enemy)
		}
	}

	return 0
}

//func SelectItem(character *models.Character, enemy *models.Enemy) {
//
//	var Inventory []string
//	for _, item := range character.Inventory {
//		Inventory = append(Inventory, item.Name)
//	}
//	if len(Inventory) == 0 {
//		fmt.Println("У вас пустой инвентарь")
//		return
//	}
//	prompt := promptui.Select{
//		Label: "Выберите Item",
//		Items: Inventory,
//	}
//
//	_, value := selectFromMenu(prompt)
//
//	switch value {
//	case "Health Potion":
//		character.Health = int(float64(character.MaxHealth) * 1.25)
//		if character.Health > character.MaxHealth {
//			character.Health = character.MaxHealth
//		}
//	case "Attack Potion":
//		character.DamageMultiplier = 1.25
//	case "Fire Potion":
//		models.TrueDamage(enemy, 100)
//	}
//
//	var buf = character.Inventory[value]
//	buf.Count -= 1
//	if buf.Count == 0 {
//		delete(character.Inventory, value)
//	} else {
//		character.Inventory[value] = buf
//	}
//}
//
//func SelectSkill(character *models.Character, enemy *models.Enemy) {
//
//	var Skills []string
//	for _, skill := range character.SKills {
//		if skill.Type == "active" && character.Level >= skill.LevelUnlock && skill.CoolDown == 0 {
//			Skills = append(Skills, skill.Name)
//		}
//	}
//	if len(Skills) == 0 {
//		fmt.Println("У вас нет доступных активных навыков")
//		return
//	}
//	prompt := promptui.Select{
//		Label: "Выберите Skill",
//		Items: Skills,
//	}
//
//	_, value := selectFromMenu(prompt)
//
//	switch value {
//	case "DoubleAttack":
//		models.TrueDamage(enemy, character.MaxDamage*2)
//	case "DrainStrike":
//		damage := int(float64(character.MaxDamage) * 1.5)
//		models.TrueDamage(enemy, damage)
//		character.Health += damage
//		if character.Health > character.MaxHealth {
//			character.Health = character.MaxHealth
//		}
//	case "Fireball":
//		models.TrueDamage(enemy, character.MaxDamage*3)
//	case "Heal":
//		character.Health += int(float64(character.MaxHealth) * 1.3)
//		if character.Health > character.MaxHealth {
//			character.Health = character.MaxHealth
//		}
//	}
//
//	var buf = character.SKills[value]
//	buf.CoolDown = buf.BaseCoolDown
//	character.SKills[value] = buf
//}
