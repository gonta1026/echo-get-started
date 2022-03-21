package dbconfig

import (
	"echo-get-started/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var Master *sqlx.DB
var Slave *sqlx.DB

func InitMySQL() error {

	master, err := initMaster()
	if err != nil {
		return err
	}

	slave, err := initSlave()
	if err != nil {
		return err
	}

	Master = master
	Slave = slave

	return nil
}

func initMaster() (*sqlx.DB, error) {
	formatConfig := fmt.Sprintf(`%s:%s@/%s`, config.Config.DbUser, config.Config.DbPassword, config.Config.DbName)

	db, err := sqlx.Connect(config.Config.SQLDriver, formatConfig)
	if err != nil {
		return nil, errors.Wrap(err, "mysql master instance connection refused")
	}

	return db, err
}

func initSlave() (*sqlx.DB, error) {
	formatConfig := fmt.Sprintf(`%s:%s@/%s`, config.Config.DbUser, config.Config.DbPassword, config.Config.DbName)

	db, err := sqlx.Connect("mysql", formatConfig)

	if err != nil {
		return nil, errors.Wrap(err, "mysql slave instance connection refused")
	}

	return db, err
}
