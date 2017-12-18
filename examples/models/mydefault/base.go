package mydefault

import (
	_ "github.com/go-sql-driver/mysql" //remember: import xxx driver to init
	"github.com/go-xorm/xorm"
	"github.com/hhxsv5/go-db-proxy-api/db/models"
	mydb "github.com/hhxsv5/go-db-proxy-api/db"
)

//default connection in file db.toml
const conn = "default"

var (
	db  *mydb.DB
	orm *xorm.Engine
)

func init() {
	db = models.NewDB(conn)
	orm = db.Engine
}
