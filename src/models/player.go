package models

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Player struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Class            int    `json:"class_id"`
	Level            int    `json:"level"`
	CurrentExp       int    `json:"current_exp"`
	LevelExp         int    `json:"level_exp"`
	MinDamage        int    `json:"min_damage"`
	MaxDamage        int    `json:"max_damage"`
	Health           int    `json:"health"`
	MaxHealth        int    `json:"max_health"`
	CriticalChance   int    `json:"critical_chance"`
	DamageMultiplier float64
}

func (player Player) Attack(enemy *Enemy) {
	rand.Seed(time.Now().UnixNano())
	var criticalMultiplier float64 = 1
	var chance = rand.Intn(100)
	if chance < player.CriticalChance {
		criticalMultiplier = 2
	}
	var playerDamage = int(float64(rand.Intn(player.MaxDamage-player.MinDamage)+player.MinDamage) * player.DamageMultiplier * criticalMultiplier)
	enemy.Health = enemy.Health - playerDamage

	fmt.Printf("Вы  наносите %d , у противника осталось %d здоровья", playerDamage, enemy.Health)
	fmt.Println()
}

func (player *Player) GainExp(exp int) {

	fmt.Printf("Вы получули %d опыта", exp)
	player.CurrentExp += exp
	player.DamageMultiplier = 1

	for player.CurrentExp >= player.LevelExp {
		player.Level += 1
		player.CurrentExp -= player.LevelExp
		player.LevelExp += 100
		player.MaxHealth += 5
		player.Health = player.MaxHealth
		player.MinDamage += 1
		player.MaxDamage += 1

		//for key, skill := range player.SKills {
		//	if skill.Type == "passive" && skill.LevelUnlock == player.Level {
		//		var buf = player.SKills[key]
		//		player.SKills[key] = buf
		//
		//		switch key {
		//		case "Fat":
		//			player.MaxHealth += 100
		//		case "CriticalDamage":
		//			player.CriticalChance += 20
		//		}
		//	}
		//}

		fmt.Printf("Вы получили уровень %d", player.Level)
	}

	fmt.Printf("Ваш уровень %d. Опыт %d/%d", player.Level, player.CurrentExp, player.LevelExp)
	fmt.Println()

}

func (player Player) CheckHealth() {
	if player.Health <= 0 {
		fmt.Println("Вы погибли. Игра окончена")
		fmt.Println()

		fmt.Println("Спасибо за игру")
		os.Exit(0)
	}
}
