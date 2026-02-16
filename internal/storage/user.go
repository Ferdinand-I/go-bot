package storage

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        int64  `db:"id"`
	TGID      int64  `db:"tg_id"`
	FirstName string `db:"first_name"`
	Username  string `db:"username"`
}

type UserRepo struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) TgUserExists(id int64) (bool, error) {
	var exists bool

	err := r.DB.Get(&exists, "SELECT EXISTS(SELECT 1 FROM users WHERE tg_id = $1)", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return exists, nil
}

func (r *UserRepo) Create(user *User) error {
	_, err := r.DB.NamedExec(
		"INSERT INTO users (tg_id, first_name, username) VALUES (:tg_id, :first_name, :username)",
		user,
	)
	return err
}
