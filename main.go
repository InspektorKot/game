package main

import (
	"fmt"
	"github.com/InspektorKot/game.git/classes"
	"github.com/InspektorKot/game.git/menu"
	"github.com/go-pg/pg/v9"
	"os"
	//"github.com/go-pg/pg/v9/orm"
)

func main() {

	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "example",
		Database: "game",
	})
	defer db.Close()

	var classDB []classes.Enemy

	err := db.Model(&classDB).Select()
	if err != nil {
		panic(err)
	}

	fmt.Println(classDB)

	os.Exit(1)

	hero, name := menu.MainMenu()

	day := 1

	for day <= 100 {

		fmt.Printf("День %d ", day)
		fmt.Println()

		var enemy = classes.CreateEnemy()

		fmt.Printf("Вы встретили %s (%d HP)", enemy.Name, enemy.Health)
		fmt.Println()

		for {
			key := menu.SelectAction(&hero, &enemy)
			if key == 1 {
				classes.SaveHero(hero, name)
				break
			}
			enemy.CheckHealth(&hero)

			for key, skill := range hero.SKills {
				if skill.Type == "active" && skill.CoolDown > 0 {
					var buf = hero.SKills[key]
					buf.CoolDown = buf.CoolDown - 1
					hero.SKills[key] = buf
				}
			}

			if enemy.Health <= 0 {
				for key, skill := range hero.SKills {
					if skill.Type == "active" && skill.CoolDown > 0 {
						var buf = hero.SKills[key]
						buf.CoolDown = 0
						hero.SKills[key] = buf
					}
				}
				classes.SaveHero(hero, name)
				break
			}

			enemy.Attack(&hero)
			hero.CheckHealth()

		}
		day++
	}
}
