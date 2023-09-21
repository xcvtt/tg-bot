package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	Db *sql.DB
}

func New() *Store {
	return &Store{}
}

func (s *Store) Open() error {
	db, err := sql.Open(
		"postgres",
		"postgres://localhost:5433/shitbot?sslmode=disable&user=postgres&password=123")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.Db = db

	return nil
}

func (s *Store) Close() error {
	err := s.Db.Close()
	if err != nil {
		return err
	}

	return nil
}
