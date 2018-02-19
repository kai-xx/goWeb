package router

import (
	"github.com/devfeel/dotweb"
	"github.com/devfeel/middleware/cors"
	"goWeb/controller"
)

func Run(Router *dotweb.HttpServer)  {
	option := cors.NewConfig().UseDefault()

	Router.GET("/", func(context dotweb.Context) error {
		return context.WriteString("demo 首页")
	})

	Router.GET("/json", func(context dotweb.Context) error {
		return context.WriteJson("json 测试")
	}).Use(cors.Middleware(option))

	Router.GET("/goodsList", controller.GetGoodsList).Use(cors.Middleware(option))
}