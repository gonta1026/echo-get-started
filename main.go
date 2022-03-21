package main

import (
	"echo-get-started/app/infrastructure/dbconfig"
	"echo-get-started/app/router"
	"echo-get-started/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	setup()
	router.Router()
}

func setup() {
	// 初期設定の読み込み
	config.LoadConfig()
	if err := dbconfig.InitMySQL(); err != nil {
		fmt.Println(err)
	}
}
