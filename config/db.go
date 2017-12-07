package config

import (
	"github.com/BurntSushi/toml"
	"errors"
)

type Connection struct {
	Driver   string
	Dsn      string
	Log      bool
	LogLevel int
	LogFile  string
}

type Connections struct {
	Connections map[string]Connection
}

func GetDbConfigs() map[string]Connection {

	var conns Connections
	if _, err := toml.DecodeFile("config/db.toml", &conns); err != nil {
		panic(errors.New("parse db.toml fail"))
	}
	return conns.Connections
}

func GetDbConfig(id string) Connection {
	conns := GetDbConfigs()
	conn, err := conns[id]
	if !err {
		panic(errors.New("connection " + id + " is not available"))
	}
	return conn
}
