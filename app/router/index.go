package router

import (
	"echo-get-started/config"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func getUsers(c echo.Context) error {
	var userlist []User

	db := config.Db()

	rows, err := db.Queryx("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	var user User
	for rows.Next() {
		//rows.Scanの代わりにrows.StructScanを使う
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
		}
		userlist = append(userlist, user)
	}
	return c.JSON(http.StatusOK, userlist)
}

func Router() {
	e := echo.New()
	e.GET("/users", getUsers)
	e.Logger.Fatal(e.Start(fmt.Sprintf(`:%s`, config.Config.Port)))
}
