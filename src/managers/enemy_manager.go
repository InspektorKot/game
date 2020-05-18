package managers

import (
	"github.com/InspektorKot/game.git/src/models"
	"github.com/InspektorKot/game.git/src/storage"
	sq "github.com/Masterminds/squirrel"
)

type EnemyManager struct {
	store *storage.Storage
}

func NewEnemyManager(store *storage.Storage) EnemyManager {
	return EnemyManager{
		store: store,
	}
}

func GetEnemy(manager EnemyManager) models.Enemy {
	enemy := new(models.Enemy)
	sql := sq.Select("id", "name", "level", "exp", "min_damage", "max_damage", "health").From("enemy").OrderBy("random()").PlaceholderFormat(sq.Dollar)
	query, values := sql.MustSql()

	row := manager.store.Db.QueryRow(query, values...)

	err := row.Scan(&enemy.Id, &enemy.Name, &enemy.Level, &enemy.Exp, &enemy.MinDamage, &enemy.MaxDamage, &enemy.Health)
	if err != nil {
		panic(err)
	}

	return *enemy
}
