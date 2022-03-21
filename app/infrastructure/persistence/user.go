package persistence

import (
	"echo-get-started/app/domain/model/user"
	"log"

	"github.com/jmoiron/sqlx"
)

type UserPersistence struct {
	Master *sqlx.DB
	Slave  *sqlx.DB
}

func NewUserPersistence(master *sqlx.DB, slave *sqlx.DB) *UserPersistence {
	return &UserPersistence{
		Master: master,
		Slave:  slave,
	}
}

func (userPersistence UserPersistence) List() ([]user.Entity, error) {

	var userlist []user.Entity
	rows, err := userPersistence.Master.Queryx("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	var user user.Entity
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
		}
		userlist = append(userlist, user)
	}

	return userlist, err
}
