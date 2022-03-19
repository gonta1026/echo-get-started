package main

import (
	"echo-get-started/app/router"
	"echo-get-started/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.LoadConfig() // 初期設定の読み込み
	router.Router()
}
