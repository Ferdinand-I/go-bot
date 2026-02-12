package db_models

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        int64  `db:"id"`
	TGID      int64  `db:"tg_id"`
	FirstName string `db:"first_name"`
	Username  string `db:"username"`
}

func GetUserByTgID(db *sqlx.DB, ID int64) *User {
	user := &User{}

	err := db.Get(user, "SELECT * FROM users WHERE tg_id = $1", ID)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			log.Print(err.Error())
		}

		return nil
	}

	return user
}

func CreateUser(db *sqlx.DB, user *User) error {
	_, err := db.NamedExec("INSERT INTO users (tg_id, first_name, username) VALUES (:tg_id, :first_name, :username)", user)
	if err != nil {
		log.Print(err.Error())

		return err
	}

	return nil
}
