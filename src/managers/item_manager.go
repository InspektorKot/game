package managers

import (
	"github.com/InspektorKot/game.git/src/storage"
	sq "github.com/Masterminds/squirrel"
	"math/rand"
	"time"
)

type ItemManaqer struct {
	store *storage.Storage
}

func NewItemManaqer(store *storage.Storage) ItemManaqer {
	return ItemManaqer{
		store: store,
	}
}

func Drop(manager ItemManaqer, enemy_id int, player_id int) {
	classes := sq.Select("item_id", "amount", "chance").From("enemy_drop").Where(sq.Eq{"enemy_id": enemy_id}).PlaceholderFormat(sq.Dollar)
	query, values := classes.MustSql()

	rows, err := manager.store.Db.Query(query, values...)

	if err != nil {
		panic("Query error")
	}

	var item_id, amount, dropChance int
	for rows.Next() {
		rows.Scan(&item_id, &amount, &dropChance)
		rand.Seed(time.Now().UnixNano())
		var chance = rand.Intn(100)
		if chance < dropChance {
			Add(manager, item_id, amount, player_id)
		}
	}
}

func Add(manager ItemManaqer, item_id int, amount int, player_id int) {
	items := sq.Select("amount", "id").From("player_item").Where(sq.Eq{"player_id": player_id, "item_id": item_id}).PlaceholderFormat(sq.Dollar)

	query, values := items.MustSql()
	row := manager.store.Db.QueryRow(query, values...)

	var itemsAmount, playerItemId int
	err := row.Scan(&itemsAmount, &playerItemId)

	if playerItemId == 0 {
		query := sq.Insert("player_item").
			Columns("player_id", "item_id", "amount").
			Values(player_id, item_id, amount).
			RunWith(&manager.store.Db).
			PlaceholderFormat(sq.Dollar)

		_, err = query.Exec()

		if err != nil {
			panic(err)
		}
	} else {
		query := sq.Update("player_item").
			Set("amount", itemsAmount+amount).
			Set("updated_at", "now()").
			Where(sq.Eq{"id": playerItemId}).
			RunWith(&manager.store.Db).
			PlaceholderFormat(sq.Dollar)

		_, err = query.Exec()

		if err != nil {
			panic(err)
		}
	}
}
