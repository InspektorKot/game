package models

import "github.com/InspektorKot/game.git/src/storage"

type Class struct {
	Db storage.Storage
	Id             int `json:"id"`
	Name           string `json:"name"`
	MinDamage      int `json:"min_damage"`
	MaxDamage      int `json:"max_damage"`
	Health         int `json:"health"`
	CriticalChance int `json:"critical_chance"`
}
