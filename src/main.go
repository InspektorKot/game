package main

import (
	"database/sql"
	"fmt"
	"github.com/InspektorKot/game.git/src/managers"
	"github.com/InspektorKot/game.git/src/menu"
	"github.com/InspektorKot/game.git/src/models"
	"github.com/InspektorKot/game.git/src/storage"
	//sq "github.com/Masterminds/squirrel"
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

	//query := sq.Insert("models").
	//	Columns("name","min_damage","max_damage","health","critical_chance").
	//	Values("Warrior",4,7,22,15).
	//	RunWith(conn).
	//	PlaceholderFormat(sq.Dollar)
	//_, err = query.Exec()

	//fmt.Printf( "%s\n", data[0].Health);
	//
	//fmt.Println(err)
	//os.Exit(1)

	hero, name := menu.MainMenu(menuDataManager)

	day := 1

	for day <= 100 {

		fmt.Printf("День %d ", day)
		fmt.Println()

		var enemy = models.CreateEnemy()

		fmt.Printf("Вы встретили %s (%d HP)", enemy.Name, enemy.Health)
		fmt.Println()

		for {
			key := menu.SelectAction(&hero, &enemy)
			if key == 1 {
				models.SaveHero(hero, name)
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
				models.SaveHero(hero, name)
				break
			}

			enemy.Attack(&hero)
			hero.CheckHealth()

		}
		day++
	}
}
