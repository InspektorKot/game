package main

import (
	"database/sql"
	"fmt"
	"github.com/InspektorKot/game.git/src/managers"
	"github.com/InspektorKot/game.git/src/menu"
	"github.com/InspektorKot/game.git/src/storage"
	_ "github.com/lib/pq" // here
)

func main() {
	conn, err := sql.Open("postgres", "host=localhost  dbname=game user=postgres password=example sslmode=disable ")
	if err != nil {
		panic("Can't connect to DB")
	}
	defer conn.Close()

	s := storage.New(conn)

	menuDataManager := managers.NewMenuDataManager(s)
	playerManager := managers.NewPlayerManager(s)
	enemyManager := managers.NewEnemyManager(s)
	itemManager := managers.NewItemManaqer(s)

	hero := menu.MainMenu(menuDataManager, playerManager)

	day := 1

	for day <= 100 {

		fmt.Printf("День %d ", day)
		fmt.Println()

		var enemy = managers.GetEnemy(enemyManager)

		fmt.Printf("Вы встретили %s (%d HP)", enemy.Name, enemy.Health)
		fmt.Println()

		for {
			key := menu.SelectAction(menuDataManager, &hero, &enemy)
			if key == 1 {
				managers.Save(playerManager, hero)
				break
			}
			if enemy.Health <= 0 {
				fmt.Println("Вы победили")
				fmt.Println()

				hero.GainExp(enemy.Exp)
				managers.Drop(itemManager, enemy.Id, hero.Id)
			}
			//
			//for key, skill := range hero.SKills {
			//	if skill.Type == "active" && skill.CoolDown > 0 {
			//		var buf = hero.SKills[key]
			//		buf.CoolDown = buf.CoolDown - 1
			//		hero.SKills[key] = buf
			//	}
			//}
			//
			if enemy.Health <= 0 {
				//	for key, skill := range hero.SKills {
				//		if skill.Type == "active" && skill.CoolDown > 0 {
				//			var buf = hero.SKills[key]
				//			buf.CoolDown = 0
				//			hero.SKills[key] = buf
				//		}
				//	}
				managers.Save(playerManager, hero)
				break
			}
			//
			enemy.Attack(&hero)
			hero.CheckHealth()

		}
		day++
	}
}
