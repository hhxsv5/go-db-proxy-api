package models

import (
	xcore "github.com/go-xorm/core"
	"github.com/hhxsv5/go-db-proxy-api/db"
)

func NewDB(conn string) *db.DB {
	db := db.NewDB(conn)

	//...some settings for xorm
	db.Engine.SetMapper(xcore.GonicMapper{})
	return db
}
