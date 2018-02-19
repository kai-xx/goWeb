package main

import (

	"fmt"
	"goWeb/bootstrap"
	"goWeb/router"
)

func main () {
	fmt.Println("start……，port:8099,visit:http://localhost:8099")

	// 加载数据库
	defer bootstrap.DB.Close()

	router.Run(bootstrap.App.HttpServer)

	err := bootstrap.App.StartServer(8099)

	if err != nil {
		fmt.Println(err)
	}


}
