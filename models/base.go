package models

import (
	xcore "github.com/go-xorm/core"
	"github.com/hhxsv5/go-db-proxy-api/core"
)

func NewDB(conn string) *core.DB {
	db := core.NewDB(conn)

	//...some settings for xorm
	db.Engine.SetMapper(xcore.GonicMapper{})
	return db
}
