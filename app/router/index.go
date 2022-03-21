package router

import (
	"echo-get-started/app/presentation/di"
	"echo-get-started/config"
	"fmt"

	"github.com/labstack/echo"
)

type User struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func Router() {
	hander := di.InitializeHandler()
	e := echo.New()
	apiGroup := e.Group("/api")

	apiGroup.GET("/users", hander.User.Users)
	e.Logger.Fatal(e.Start(fmt.Sprintf(`:%s`, config.Config.Port)))
}
