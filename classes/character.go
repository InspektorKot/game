package classes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
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

func (hero Character) Attack(enemy *Enemy) {
	rand.Seed(time.Now().UnixNano())
	var criticalMultiplier float64 = 1
	var chance = rand.Intn(100)
	if chance < hero.CriticalChance {
		criticalMultiplier = 2
	}
	var heroDamage = int(float64(rand.Intn(hero.MaxDamage-hero.MinDamage)+hero.MinDamage) * hero.DamageMultiplier * criticalMultiplier)
	enemy.Health = enemy.Health - heroDamage

	fmt.Printf("Вы  наносите %d , у противника осталось %d здоровья", heroDamage, enemy.Health)
	fmt.Println()
}

func TrueDamage(enemy *Enemy, damage int) {
	enemy.Health = enemy.Health - damage

	fmt.Printf("Вы  наносите %d , у противника осталось %d здоровья", damage, enemy.Health)
	fmt.Println()
}

func (hero Character) CheckHealth() {
	if hero.Health <= 0 {
		fmt.Println("Вы погибли. Игра окончена")
		fmt.Println()

		fmt.Println("Спасибо за игру")
		os.Exit(0)
	}
}

func (hero *Character) gainExp(exp int) {

	fmt.Printf("Вы получули %d опыта", exp)
	hero.CurrentExp += exp
	hero.DamageMultiplier = 1

	for hero.CurrentExp >= hero.LevelExp {
		hero.Level += 1
		hero.CurrentExp -= hero.LevelExp
		hero.LevelExp += 100
		hero.MaxHealth += 5
		hero.Health = hero.MaxHealth
		hero.MinDamage += 1
		hero.MaxDamage += 1

		for key, skill := range hero.SKills {
			if skill.Type == "passive" && skill.LevelUnlock == hero.Level {
				var buf = hero.SKills[key]
				hero.SKills[key] = buf

				switch key {
				case "Fat":
					hero.MaxHealth += 100
				case "CriticalDamage":
					hero.CriticalChance += 20
				}
			}
		}

		fmt.Printf("Вы получили уровень %d", hero.Level)
	}

	fmt.Printf("Ваш уровень %d. Опыт %d/%d", hero.Level, hero.CurrentExp, hero.LevelExp)
	fmt.Println()

}

func (hero *Character) GainItem(item Item) {
	if _, ok := hero.Inventory[item.Name]; ok {
		item.Count = hero.Inventory[item.Name].Count + 1
	}
	hero.Inventory[item.Name] = item
}

func Create(data int) Character {
	var result Character

	switch data {
	case 0:
		{
			result = Character{
				MinDamage: 3,
				MaxDamage: 6,
				Health:    30,
			}
			result.SKills = make(map[string]Skill)
			result.SKills["DoubleAttack"] = Skill{
				Name:         "DoubleAttack",
				LevelUnlock:  5,
				Description:  "Deals double damage",
				Type:         "active",
				CoolDown:     0,
				BaseCoolDown: 3,
			}
			result.SKills["Fat"] = Skill{
				Name:        "Fat",
				LevelUnlock: 20,
				Description: "+100 passive HP",
				Type:        "passive",
			}
		}
	case 1:
		{
			result = Character{
				MinDamage: 10,
				MaxDamage: 15,
				Health:    30,
			}
			result.SKills = make(map[string]Skill)
			result.SKills["DrainStrike"] = Skill{
				Name:         "DrainStrike",
				LevelUnlock:  5,
				Description:  "Deals 1,5x damage and heals 0,5x damage",
				Type:         "active",
				CoolDown:     0,
				BaseCoolDown: 4,
			}
			result.SKills["CriticalDamage"] = Skill{
				Name:        "CriticalDamage",
				LevelUnlock: 20,
				Description: "20% CriticalChance",
				Type:        "passive",
			}
		}
	case 2:
		{
			result = Character{
				MinDamage: 4,
				MaxDamage: 5,
				Health:    25,
			}
			result.SKills = make(map[string]Skill)
			result.SKills["Fireball"] = Skill{
				Name:         "Fireball",
				LevelUnlock:  5,
				Description:  "Deals 3x damage",
				Type:         "active",
				CoolDown:     0,
				BaseCoolDown: 3,
			}
			result.SKills["Heal"] = Skill{
				Name:         "Heal",
				LevelUnlock:  20,
				Description:  "Heals 30% HP",
				Type:         "active",
				CoolDown:     0,
				BaseCoolDown: 5,
			}
		}
	}

	result.Level = 1
	result.DamageMultiplier = 1
	result.MaxHealth = result.Health
	result.CurrentExp = 0
	result.LevelExp = 200
	result.Inventory = make(map[string]Item)

	return result
}

func SaveHero(hero Character, name string) {
	file, _ := json.MarshalIndent(hero, "", " ")
	_ = ioutil.WriteFile(fmt.Sprintf("%s.json", name), file, 0644)
}
