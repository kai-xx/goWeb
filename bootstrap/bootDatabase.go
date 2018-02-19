package bootstrap

import (
	"github.com/gohouse/gorose"
	"goWeb/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB gorose.Connection

func init()  {
	var err error
	DB, err = gorose.Open(config.DbConfig)

	if err != nil {
		panic("数据库链接失败" + err.Error())
	}
	errs := DB.Ping()
	if errs != nil {
		panic("数据库链接失败" + err.Error())
	}
}