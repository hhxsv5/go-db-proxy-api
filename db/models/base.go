package models

import (
	"github.com/go-xorm/core"
	"github.com/hhxsv5/go-db-proxy-api/db"
)

func NewDB(conn string) *db.DB {
	d := db.NewDB(conn)

	//...some settings for xorm
	d.Engine.SetMapper(core.SnakeMapper{})
	return d
}
