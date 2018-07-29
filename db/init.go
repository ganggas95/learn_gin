package db

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
)

var DbEngine *xorm.Engine

func InitDB() *xorm.Engine{

	DbEngine, _ = xorm.NewEngine(
		"mysql",
		"root:b1sm1llah@/todo_db?charset=utf8")
	DbEngine.ShowSQL(true)
	DbEngine.Logger().SetLevel(core.LOG_DEBUG)
	return DbEngine
}
