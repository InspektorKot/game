package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Enemy struct {
	Id        int
	MinDamage int
	MaxDamage int
	Health    int
	Exp       int
	Level     int
	Name      string
}

func (enemy Enemy) Attack(player *Player) {
	rand.Seed(time.Now().UnixNano())
	var enemyDamage = rand.Intn(enemy.MaxDamage-enemy.MinDamage) + enemy.MinDamage
	player.Health = player.Health - enemyDamage

	fmt.Printf("Вам  наносят %d , у вас осталось %d здоровья", enemyDamage, player.Health)
	fmt.Println()
}
