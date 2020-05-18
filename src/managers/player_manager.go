package managers

import (
	"fmt"
	"github.com/InspektorKot/game.git/src/models"
	"github.com/InspektorKot/game.git/src/storage"
	sq "github.com/Masterminds/squirrel"
	"os"
)

type PlayerManager struct {
	store *storage.Storage
}

func NewPlayerManager(store *storage.Storage) PlayerManager {
	return PlayerManager{
		store: store,
	}
}

func NewPlayer(pm PlayerManager, className string, nickName string) *models.Player {
	player := new(models.Player)
	player.Name = nickName

	classes := sq.Select("health", "health", "critical_chance", "min_damage", "max_damage", "id").From("class").Where(sq.Like{"name": className}).PlaceholderFormat(sq.Dollar)
	query, values := classes.MustSql()
	row := pm.store.Db.QueryRow(query, values...)

	err := row.Scan(&player.Health, &player.MaxHealth, &player.CriticalChance, &player.MinDamage, &player.MaxDamage, &player.Class)
	if err != nil {
		panic(err)
	}

	player.Level = 1
	player.DamageMultiplier = 1
	player.CurrentExp = 0
	player.LevelExp = 200

	new_query := sq.Insert("player").
		Columns("name", "class_id", "level", "current_exp", "level_exp", "min_damage", "max_damage", "health", "max_health", "critical_chance").
		Values(player.Name, player.Class, player.Level, player.CurrentExp, player.LevelExp, player.MinDamage, player.MaxDamage, player.Health, player.MaxHealth, player.CriticalChance).
		Suffix("RETURNING \"id\"").
		RunWith(&pm.store.Db).
		PlaceholderFormat(sq.Dollar)

	//_, err = new_query.Exec()
	new_query.QueryRow().Scan(&player.Id)

	if err != nil {
		panic(err)
	}

	return player
}

func Load(pm PlayerManager, nickName string) *models.Player {
	player := new(models.Player)

	hero := sq.Select("id", "name", "class_id", "current_exp", "level", "level_exp", "min_damage", "max_damage", "health", "max_health", "critical_chance").From("player").Where(sq.Like{"name": nickName}).PlaceholderFormat(sq.Dollar)
	query, values := hero.MustSql()
	row := pm.store.Db.QueryRow(query, values...)

	err := row.Scan(&player.Id, &player.Name, &player.Class, &player.CurrentExp, &player.Level, &player.LevelExp, &player.MinDamage, &player.MaxDamage, &player.Health, &player.MaxHealth, &player.CriticalChance)
	if err != nil {
		fmt.Println("Нет такого персонажа")
		os.Exit(1)
	}

	player.DamageMultiplier = 1

	return player
}

func Save(pm PlayerManager, player models.Player) {
	query := sq.Update("player").
		Set("current_exp", player.CurrentExp).
		Set("level_exp", player.LevelExp).
		Set("level", player.Level).
		Set("min_damage", player.MinDamage).
		Set("max_damage", player.MaxDamage).
		Set("health", player.Health).
		Set("max_health", player.MaxHealth).
		Set("critical_chance", player.CriticalChance).
		Set("updated_at", "now()").
		Where(sq.Eq{"id": player.Id}).
		RunWith(&pm.store.Db).
		PlaceholderFormat(sq.Dollar)

	_, err := query.Exec()

	if err != nil {
		panic(err)
	}
}
