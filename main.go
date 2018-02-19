package main

import (

	"fmt"
	// 引入驱动
	"goWeb/bootstrap"
	// 引入项目路由
	"goWeb/router"
)

func main () {
	fmt.Println("start……，port:8099,visit:http://localhost:8099")

	// 加载数据库
	defer bootstrap.DB.Close()

	// 加载路由
	router.Run(bootstrap.App.HttpServer)

	// 启动web服务
	err := bootstrap.App.StartServer(8099)

	if err != nil {
		fmt.Println(err)
	}


}
