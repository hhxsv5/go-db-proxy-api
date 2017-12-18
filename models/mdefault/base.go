package mdefault

import (
	_ "github.com/go-sql-driver/mysql" //remember: import xxx driver to init
	"github.com/hhxsv5/go-db-proxy-api/models"
	"github.com/hhxsv5/go-db-proxy-api/core"
	"github.com/go-xorm/xorm"
)

//default connection in file db.toml
const conn = "default"

var (
	db  *core.DB
	orm *xorm.Engine
)

func init() {
	db = models.NewDB(conn)
	orm = db.Engine
}
