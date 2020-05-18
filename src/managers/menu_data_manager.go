package managers

import (
	"github.com/InspektorKot/game.git/src/storage"
	sq "github.com/Masterminds/squirrel"
)

type MenuDataManager struct {
	store *storage.Storage
}

func NewMenuDataManager(store *storage.Storage) MenuDataManager {
	return MenuDataManager{
		store: store,
	}
}

func GetClassNameList(mdm MenuDataManager) []string {
	classes := sq.Select("name").From("class").PlaceholderFormat(sq.Dollar)
	query, _ := classes.MustSql()

	rows, err := mdm.store.Db.Query(query)

	if err != nil {
		panic("Query error")
	}

	var data []string
	for rows.Next() {
		var test string
		rows.Scan(&test)
		data = append(data, test)
	}
	return data
}

//func GetPlayerItemList(mdm MenuDataManager) []string {
//
//}
