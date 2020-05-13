package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Enemy struct {
	MinDamage int
	MaxDamage int
	Health    int
	Exp       int
	Level     int
	Name      string
	Inventory []Item
}

func (enemy Enemy) Attack(hero *Character) {
	rand.Seed(time.Now().UnixNano())
	var enemyDamage = rand.Intn(enemy.MaxDamage-enemy.MinDamage) + enemy.MinDamage
	hero.Health = hero.Health - enemyDamage

	fmt.Printf("Вам  наносят %d , у вас осталось %d здоровья", enemyDamage, hero.Health)
	fmt.Println()
}

func (enemy Enemy) CheckHealth(hero *Character) {
	if enemy.Health <= 0 {
		fmt.Println("Вы победили")
		fmt.Println()

		hero.gainExp(enemy.Exp)
		enemy.drop(hero)
	}
}

func (enemy Enemy) drop(hero *Character) {
	for _, item := range enemy.Inventory {
		rand.Seed(time.Now().UnixNano())
		var chance = rand.Intn(100)
		if chance < item.Chance {
			hero.GainItem(item)
		}
	}
}

func CreateEnemy() Enemy {
	rand.Seed(time.Now().UnixNano())
	var data = rand.Intn(3)
	switch data {
	case 0:
		{
			return Enemy{
				MinDamage: 1,
				MaxDamage: 2,
				Health:    20,
				Exp:       100,
				Name:      "Human Peasant",
				Inventory: []Item{
					{
						Name:        "Health Potion",
						Count:       1,
						Description: "+25% HP",
						Chance:      30,
					},
				},
			}
		}
	case 1:
		{
			return Enemy{
				MinDamage: 3,
				MaxDamage: 5,
				Health:    40,
				Exp:       300,
				Name:      "Orc Grunt",
				Inventory: []Item{
					{
						Name:        "Health Potion",
						Count:       1,
						Description: "+25% HP",
						Chance:      20,
					},
					{
						Name:        "Attack Potion",
						Count:       1,
						Description: "+25% Damage in this fight",
						Chance:      25,
					},
					{
						Name:        "Fire Potion",
						Count:       1,
						Description: "Deals 100 dmg",
						Chance:      10,
					},
				},
			}
		}
	case 2:
		{
			return Enemy{
				MinDamage: 1,
				MaxDamage: 2,
				Health:    100,
				Exp:       600,
				Name:      "Elf Treant",
				Inventory: []Item{
					{
						Name:        "Attack Potion",
						Count:       1,
						Description: "+25% Damage in this fight",
						Chance:      30,
					},
					{
						Name:        "Fire Potion",
						Count:       1,
						Description: "Deals 100 dmg",
						Chance:      30,
					},
				},
			}
		}
	}
	return Enemy{
		MinDamage: 1000,
		MaxDamage: 1200,
		Health:    1000,
		Exp:       100000,
	}
}
