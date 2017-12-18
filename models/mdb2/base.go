package mdb2

import (
	_ "github.com/go-sql-driver/mysql" //remember: import xxx driver to init
	"github.com/hhxsv5/go-db-proxy-api/models"
	"github.com/hhxsv5/go-db-proxy-api/core"
	"github.com/go-xorm/xorm"
)

//db2 connection in file db.toml
const conn = "db2"

var (
	db  *core.DB
	orm *xorm.Engine
)

func init() {
	db = models.NewDB(conn)
	orm = db.Engine
}
