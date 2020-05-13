package storage

import "database/sql"

type Storage struct {
	Db sql.DB
}

func New(db *sql.DB) *Storage {
	s:= new (Storage)
	s.Db = *db
	return s
}