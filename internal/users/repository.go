package users

import (
	"tgbot/internal/store"
	"tgbot/models"
)

type Repository interface {
	Create(u *models.User) error
	GetById(id int64) (*models.User, error)
	GetAll() ([]models.User, error)
}

type repository struct {
	Store *store.Store
}

func NewRepository(s *store.Store) Repository {
	return &repository{s}
}

func (r *repository) Create(u *models.User) error {
	stmt, err := r.Store.Db.Prepare("INSERT INTO users (id, name, hp) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Id, u.Name, u.Hp)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetById(id int64) (*models.User, error) {
	var u models.User

	stmt, err := r.Store.Db.Prepare("select * from users where id = $1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&u.Id, &u.Name, &u.Hp)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *repository) GetAll() ([]models.User, error) {
	var users []models.User

	stmt, err := r.Store.Db.Prepare("select * from users")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	defer rows.Close()

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Hp); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
