package router

import (
	// 引入dotweb
	"github.com/devfeel/dotweb"
	// 引入dotweb中间件
	"github.com/devfeel/middleware/cors"
	// 引入控制器
	"goWeb/controller"
)

func Run(Router *dotweb.HttpServer)  {
	// 设置cors选项中间件，使用默认的跨域配置
	option := cors.NewConfig().UseDefault()

	// 默认路由
	Router.GET("/", func(context dotweb.Context) error {
		return context.WriteString("demo 首页")
	})

	// 返回值json事例
	Router.GET("/json", func(context dotweb.Context) error {
		return context.WriteJson("json 测试")
	}).Use(cors.Middleware(option))

	// 查询商品信息，并返回json格式
	Router.GET("/goodsList", controller.GetGoodsList).Use(cors.Middleware(option))
}