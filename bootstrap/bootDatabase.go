package bootstrap

import (
	// 引入gorose orm
	"github.com/gohouse/gorose"
	// 引入数据库配置
	"goWeb/config"
	// 引入mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// 数据库链接对象
var DB gorose.Connection

func init()  {
	var err error
	// 装配数据库信息, 并加载database
	// config.DbConfig 为数据库配置信息
	DB, err = gorose.Open(config.DbConfig)

	if err != nil {
		panic("数据库链接失败" + err.Error())
	}
	errs := DB.Ping()
	if errs != nil {
		panic("数据库链接失败" + err.Error())
	}
}