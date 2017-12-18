package models

import (
	_ "github.com/go-sql-driver/mysql" //remember: import xxx driver to init
	"github.com/go-xorm/xorm"
	xcore "github.com/go-xorm/core"
	"github.com/hhxsv5/go-db-proxy-api/core"
)

const connection = "default" //default connection in file db.toml

var (
	Db  *core.DB
	Orm *xorm.Engine
)

func init() {
	Db = core.NewDB(connection)
	Orm = Db.Engine

	//...some settings for xorm
	Orm.SetMapper(xcore.GonicMapper{})
}
