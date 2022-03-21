package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port    string
	LogFile string

	SQLDriver  string
	DbName     string
	DbUser     string
	DbPassword string
}

var Config ConfigList

func loggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	loggingSettings("webapp.log")

	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port:       cfg.Section("web").Key("port").String(),
		LogFile:    cfg.Section("web").Key("logfile").String(),
		SQLDriver:  cfg.Section("db").Key("driver").String(),
		DbUser:     cfg.Section("db").Key("db_user").String(),
		DbName:     cfg.Section("db").Key("db_name").String(),
		DbPassword: cfg.Section("db").Key("db_password").String(),
	}
}
