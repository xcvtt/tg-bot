package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type Store struct {
	Db *sql.DB
}

func New() *Store {
	return &Store{}
}

func (s *Store) Open() error {
	var connectionString = fmt.Sprintf("postgres://postgres:%s/%s?sslmode=disable&user=%s&password=%s", os.Getenv("POSTGRES_DB_PORT"), os.Getenv("POSTGRES_SHITBOT"), os.Getenv("POSTGRES_SHITBOT"), os.Getenv("POSTGRES_SHITBOT"))
	db, err := sql.Open(
		"postgres",
		connectionString)
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
