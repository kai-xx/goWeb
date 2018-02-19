# goWeb
gorose orm+dotweb框架实现接口 -- 类似PHP的laravel
# 项目结构说明
因为 gorose orm 和 dotweb 框架 都是基于独立的组件, 所以我们可以对项目的目录做出自由的规划.
在这里, 我们便规划一个遵循一般的mvc结构的项目目录.
为了项目的易维护性和高可用性, 我们尽量采取成熟的架构, 分离业务模块, 这里就借鉴 php 界最火爆的框架拉 laravel 的设计思想, 做出简单的涉设计架构
目录大致如下:

1. 路由: /router/
放置所有路由, 这里会用到dotweb的路由组件, 所有请求的入口都会经过这里

2. 控制器: /controller/
放置所有控制器, 这里就是基本的业务逻辑处理曾

3. 数据操作: /model/
放置所有model, 这里就是gorose orm的主战场, 所有的数据库操作都在这里完成

4. 中间件: /middleware/
放置所有中间件, 这里可以放置各种狂拽酷炫吊炸天的组件

5. 入口目录: /bootstrap/
系统中用到的组件, 统一在这里驱动, 便于维护管理, 本项目主要包含gorose和dotweb在这里驱动

6. 配置目录: /config/
必须是放置所有的配置文件在这里

7. 日志目录: /log/
这个没什么说的

成型目录结构
goWeb
    bootstrap
        bootDatabase.go
        bootRouter.go
    config
        database.go
    middleware
        Auth.go
    controller
        Goods.go
    model
        Goods.go
    router
        route.go
    main.go
以上各个组件, 除了model模块采用gorose的orm组件之外, 路由和部分中间件模块均由dotweb中对应的组件来完成, 这里就更能提现出组件化框架dotweb的强大之处了

# 路由说明
  这里用到了 dotweb 的相关组件, 有 HttpServer , Context 和 cors 中间件
  他们分别作用于 路由, 请求参数等相关和中间件(比如跨域) 等

  Route.GET 代表该请求位 get 请求
  Route.POST 代表该请求位 post 请求
  Route.Group 将请求归类
  Route.Group.Use 使用中间件
  option 中间件初始化
  cors.Middleware(option) 加载中间件option
  controller.xxxxx 加载控制器的对应方法

# 控制器和模型说明
````
    // 控制器
    func GetGoodsById(context dotweb.Context) error {
         res := model.GetGoodsById(context)

         ctx.WriteJson(utils.SuccessReturn(res))

         return nil
     }
````
说明：
- context 请求的内容
- context.WriteJson 返回 json 格式数据
- utils.SuccessReturn 工具包中封装的成功返回函数
- model 引入的模型目录, 内容如下:

````
    // 模型
    func GetGoodsById(context dotweb.Context) interface{} {
          res, err := bootstrap.DB.Table("goods").
              Where("id", context.FormValue("id")).
              First()
          if err != nil {
              return ""
          }
          return res
      }
````

# 测试
- 启动服务
````
    go run main.go
````
出现如下信息
````
    start……，port:8099,visit:http://localhost:8099
````

# 部署上线
- 打包项目
````
    go build main.go
````
- 运行项目
````
    nohup ./main &
````