package db

import (
	"github.com/go-xorm/xorm"
	"os"
	"github.com/go-xorm/core"
	"time"
)

type DB struct {
	Cfg    Connection
	Engine *xorm.Engine
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
		if cfg.LogFile != "" {
			f, err := os.OpenFile(cfg.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				panic(err)
			}
			db.Engine.SetLogger(xorm.NewSimpleLogger(f))
		} //else log into console

		if cfg.Timezone != "" {
			db.Engine.TZLocation, _ = time.LoadLocation(cfg.Timezone)
		}

		db.Engine.ShowSQL(true)
		db.Engine.ShowExecTime(true)
		db.Engine.Logger().SetLevel(core.LogLevel(cfg.LogLevel))
	}

	return db
}
