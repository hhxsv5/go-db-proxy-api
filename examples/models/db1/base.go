package db1

import (
	_ "github.com/go-sql-driver/mysql" //remember: import xxx driver to init
	"github.com/hhxsv5/go-db-proxy-api/db/models"
	mydb "github.com/hhxsv5/go-db-proxy-api/db"
	"github.com/go-xorm/xorm"
)

//db1 connection in file db.toml
const conn = "db1"

var (
	db  *mydb.DB
	orm *xorm.Engine
)

func init() {
	db = models.NewDB(conn)
	orm = db.Engine
}
