package main

import (
	"CmsProject/config"
	"CmsProject/controller"
	"CmsProject/datasource"
	"CmsProject/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"time"
)

/**
*程序主入口
*/
func main() {
	app := newApp()

	//应用app设置
	Configation(app)
	//路由设置
	mvcHandle(app)
	//配置参数
	config := config.InitConfig()
	addr := ":"  + config.Port//端口号形式
	app.Run(
		iris.Addr(addr),//在端口8080监听
		iris.WithoutServerError(iris.ErrServerClosed),//无服务错误提示
		iris.WithOptimizations,//对json数据序列化更快的配置
		)


	////监听
	//app.Run(iris.Addr(":8000"),iris.WithoutServerError(iris.ErrServerClosed))
}

//构建APP
func newApp() *iris.Application  {
	app := iris.New()

	//设置日志级别,开发阶段为debug
	app.Logger().SetLevel("debug")

	//注册静态资源路径//文件映射，url路径映射到本地静态文件
	//app.StaticWeb"/static","./static") 已经被删除了
	/*
	这几天在跟着视频做iris框架的小应用，发现iris框架更新后，删除了StaticWeb方法，在网上找了半天没找到替换方法，
	最后在官方实例中找到了。实例代码目录为:iris/_examples/file-server/embedding-files-into-app
	b.HandleDir("public","./public")
	第一个变量声明url路径，第二个参数变量是静态文件夹目录。
	*/

	app.HandleDir("/static","./static")
	app.HandleDir("/manage/static","./static")
	app.HandleDir("/img","./static/img")
	//注册视图文件
	app.RegisterView(iris.HTML("./static",".html"))
	app.Get("/", func(context *context.Context) {
		context.View("index.html")
	})
	//fmt.Println("------------------------")

	return app

}
/*
项目设置
*/
func Configation(app *iris.Application) {
	//配置字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset:"UTF-8",
	}))

	//错位配置
	//未发现错误
	app.OnErrorCode(iris.StatusNotFound, func(context *context.Context) {
		context.JSON(iris.Map{
			"errmsg":iris.StatusNotFound,
			"msg":"not found",
			"data":iris.Map{},
		})
	})
	//服务内部错误
	app.OnErrorCode(iris.StatusInternalServerError, func(context *context.Context) {
		context.JSON(iris.Map{
			"errmsg":iris.StatusInternalServerError,
			"msg":"interal error",
			"data":iris.Map{},
		})
	})

}
/*
mvc 架构模式处理
*/
func mvcHandle(app *iris.Application) {
	//启用session
	sessManage := sessions.New(sessions.Config{
		Cookie:                      "sessioncookie",
		Expires:                     24 * time.Hour,
	})

	//构造数据库引擎
	engine := datasource.NewMysqlEngine()
	//管理员模块功能
	adminService := service.NewAdminService(engine)//j将上面的引擎放到模块中
	admin := mvc.New(app.Party("/admin"))//路由组
	admin.Register(adminService,sessManage.Start,)//注册处理业的务逻辑
	admin.Handle(new(controller.AdminController))//将注册的业务逻辑赋值给一个新得AdminController,并传入控制handle
}

