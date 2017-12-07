package models

import (
	_ "github.com/go-sql-driver/mysql" //remember: import xxx driver to init
	"github.com/go-xorm/xorm"
	"github.com/hhxsv5/go-db-proxy-api/core/model"
	"github.com/go-xorm/core"
)

const connection = "default" //default connection in file db.toml

var (
	Db  *model.DB
	Orm *xorm.Engine
)

func init() {
	Db = model.NewDB(connection)
	Orm = Db.Engine

	//...some settings for xorm
	Orm.SetMapper(core.GonicMapper{})
}
