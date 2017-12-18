package core

import (
	"github.com/go-xorm/xorm"
	"os"
	"runtime"
	"path"
	"github.com/BurntSushi/toml"
	"errors"
	xcore "github.com/go-xorm/core"
)

type DB struct {
	Cfg    Connection
	Engine *xorm.Engine
}

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
	_, filename, _, _ := runtime.Caller(1)
	cfg := path.Join(path.Dir(filename), "/../config/db.toml")
	if _, err := toml.DecodeFile(cfg, &conns); err != nil {
		panic(errors.New("parse db.toml fail: " + err.Error()))
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

func NewDB(conn string) *DB {
	cfg := GetDbConfig(conn)
	db := &DB{cfg, nil}

	var err error
	db.Engine, err = xorm.NewEngine(cfg.Driver, cfg.Dsn)
	if err != nil {
		panic(err)
	}

	if cfg.Log {
		//log into file
		if len(cfg.LogFile) > 0 {
			f, err := os.OpenFile(cfg.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				panic(err)
			}
			db.Engine.SetLogger(xorm.NewSimpleLogger(f))
		} //else log into console

		db.Engine.ShowSQL(true)
		db.Engine.ShowExecTime(true)
		db.Engine.Logger().SetLevel(xcore.LogLevel(cfg.LogLevel))
	}

	return db
}
