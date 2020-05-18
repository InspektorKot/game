package models

type Class struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	MinDamage      int    `json:"min_damage"`
	MaxDamage      int    `json:"max_damage"`
	Health         int    `json:"health"`
	CriticalChance int    `json:"critical_chance"`
}
