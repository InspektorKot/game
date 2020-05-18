package models

import (
	"fmt"
)

type Character struct {
	MinDamage        int
	MaxDamage        int
	Health           int
	MaxHealth        int
	DamageMultiplier float64
	Level            int
	CurrentExp       int
	LevelExp         int
	CriticalChance   int
	Inventory        map[string]Item
	SKills           map[string]Skill
}

func TrueDamage(enemy *Enemy, damage int) {
	enemy.Health = enemy.Health - damage

	fmt.Printf("Вы  наносите %d , у противника осталось %d здоровья", damage, enemy.Health)
	fmt.Println()
}

func (hero *Character) GainItem(item Item) {
	if _, ok := hero.Inventory[item.Name]; ok {
		item.Count = hero.Inventory[item.Name].Count + 1
	}
	hero.Inventory[item.Name] = item
}
