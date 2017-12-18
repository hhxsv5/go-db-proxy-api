package db

import (
	"github.com/go-xorm/xorm"
	"os"
	"github.com/go-xorm/core"
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
		if len(cfg.LogFile) > 0 {
			f, err := os.OpenFile(cfg.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				panic(err)
			}
			db.Engine.SetLogger(xorm.NewSimpleLogger(f))
		} //else log into console

		db.Engine.ShowSQL(true)
		db.Engine.ShowExecTime(true)
		db.Engine.Logger().SetLevel(core.LogLevel(cfg.LogLevel))
	}

	return db
}
