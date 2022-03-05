package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const port string = ":1323"

func helloWord(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// e.GET("/users/:id", getUser)
func getUser(context echo.Context) error {
	// User ID from path `users/:id`
	id := context.Param("id")

	return context.String(http.StatusNotExtended, "idは"+id+"です")
}

// e.GET("/show", show)
func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	arr := strings.Split(team, ",")

	buf := []byte("あいうえお")
	buf = append(buf, "かきくけこ"...)
	buf = append(buf, "さしすせそ"...)

	println(string(buf))
	tmpStr := ""
	for index, str := range arr {
		if len(arr)-1 == index {
			tmpStr += str + "です。"
			break
		}
		tmpStr += str + "と"

	}
	return c.String(http.StatusOK, "team:"+tmpStr+" member:"+member)
}

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
func main() {
	e := echo.New()
	e.GET("/", helloWord)
	e.POST("/save", save)
	e.GET("/show", show)
	e.GET("/users/:id", getUser)
	fmt.Println("起動！")
	e.Logger.Fatal(e.Start(port))
}
