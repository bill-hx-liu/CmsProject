package main

import (
	"CmsProject/config"
	"CmsProject/controller"
	"CmsProject/datasource"
	"CmsProject/model"
	"CmsProject/service"
	"CmsProject/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"io"
	"os"
	"strconv"
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

	//设定应用图标
	app.Favicon("./static/favicons/favicon.ico")

	//设置日志级别,开发阶段为debug
	app.Logger().SetLevel("debug")

	//注册静态资源路径//文件映射，url路径映射到本地静态文件
	//app.StaticWeb"/static","./static") 已经被删除了
	/*
	这几天在跟着视频做iris框架的小应用，发现iris框架更新后，删除了StaticWeb方法，在网上找了半天没找到替换方法，
	最后在官方实例中找到了。实例代码目录为:iris/_examples/file-server/embedding-files-into-app
	b.HandleDir("public","./public")
	第一个变量声明url路径，第二个参数变量是项目工程中的静态文件夹目录。
	*/
	//有更新了,详情见笔记
	//app.HandleDir("/static", "./static")
	//app.HandleDir("/manage/static","./static")
	//app.HandleDir("/img","./static/img")
	app.HandleDir("/static", iris.Dir("./static"))
	app.HandleDir("/manage/static",iris.Dir("./static"))
	app.HandleDir("/img",iris.Dir("./static/img"))



	//注册视图文件
	app.RegisterView(iris.HTML("./static",".html"))//将./static目录下的.html文件注册到RegisterView中
	app.Get("/", func(context *context.Context) {
		context.View("index.html")//渲染index.html文件
	})
	//fmt.Println("------------------------")
	//app.RegisterView(iris.Handlebars())//iris.Handlebars() js语言的框架
	//iris.Django:python语言框架
	//iris.HTML:HTML语言
	//iris.Jet等等六种视图引擎

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

	//获取redis实例
	redis := datasource.NewRedis()
	//设置session的同步位置为redis
	sessManage.UseDatabase(redis)//将session中的数据进行持久化的动作(同步到redis)

	//实例化mysql数据库引擎
	engine := datasource.NewMysqlEngine()
	//fmt.Println(engine)

	//管理员模块功能
	adminService := service.NewAdminService(engine)//j将上面的引擎放到模块中
	admin := mvc.New(app.Party("/admin"))//路由组
	admin.Register(adminService,sessManage.Start,)//注册处理业的务逻辑
	admin.Handle(new(controller.AdminController))//将注册的业务逻辑赋值给一个新得AdminController,并传入控制handle

	//统计功能模块
	statisService := service.NewStatisService(engine)
	statis := mvc.New(app.Party("/statis/{model}/{date}/"))//正则表达式
	statis.Register(statisService,sessManage.Start,)
	statis.Handle(new(controller.StatisController))

	//订单模块
	//orderService := service.

	//用户功能模块
	userService := service.NewUserService(engine)
	user := mvc.New(app.Party("/v1/users"))
	user.Register(
		userService,
		sessManage.Start,
	)
	user.Handle(new(controller.UserController))

	//获取用户详细信息
	app.Get("v1/user/{user_name}", func(context *context.Context) {
		userName := context.Params().Get("user_name")
		var user model.User
		_,err := engine.Where("user_name = ?",userName).Get(&user)
		if err != nil{
			context.JSON(iris.Map{
				"status":utils.RECODE_FAIL,
				"type":utils.RESPMSG_ERROR_USERINFO,
				"message":utils.Recode2Text(utils.RESPMSG_ERROR_USERINFO),
			})
		}else{
			context.JSON(user)
		}
		
	})

	//获取地址信息
	app.Get("v1/address/{address_id}", func(context *context.Context) {
		address_id := context.Params().Get("address_id")

		addressID,err := strconv.Atoi(address_id)
		if err != nil{
			context.JSON(iris.Map{
				"status":utils.RECODE_FAIL,
				"type":utils.RESPMSG_ERROR_ORDERINFO,
				"message":utils.Recode2Text(utils.RESPMSG_ERROR_ORDERINFO),
			})
		}
		var address model.Address
		_,err = engine.Id(addressID).Get(&address)
		if err != nil{
			context.JSON(iris.Map{
				"status":utils.RECODE_FAIL,
				"type":utils.RESPMSG_ERROR_ORDERINFO,
				"message":utils.Recode2Text(utils.RESPMSG_ERROR_ORDERINFO),
			})
		}
		//查询数据成功
		context.JSON(address)
	})

	//文件上传模块
	app.Post("/admin/update/avatar/{adminId}", func(context *context.Context) {
		adminId := context.Params().Get("adminId")
		iris.New().Logger().Info(adminId)

		file,info,err := context.FormFile("file")//文件内容,文件信息,错误
		if err != nil{
			iris.New().Logger().Info(err.Error())
			context.JSON(iris.Map{
				"status":utils.RECODE_FAIL,
				"type":utils.RESPMSG_ERROR_PICTUREADD,
				"failure":utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}
		defer file.Close()

		fname := info.Filename//xx.png等上传文件的名字
		out,err := os.OpenFile("./uploads/" + fname,os.O_WRONLY | os.O_CREATE,0666)
		if err != nil{
			iris.New().Logger().Info(err.Error())
			context.JSON(iris.Map{
				"status":utils.RECODE_FAIL,
				"type":utils.RESPMSG_ERROR_PICTUREADD,
				"failure":utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}
		iris.New().Logger().Info("文件路径:" + out.Name())
		defer out.Close()
		_,err = io.Copy(out,file)//out:要copy到的文件.file:被copy的文件
		if err != nil{
			context.JSON(iris.Map{
				"status":utils.RECODE_FAIL,
				"type":utils.RESPMSG_ERROR_PICTUREADD,
				"failure":utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}

		intAdminId,_ := strconv.Atoi(adminId)
		adminService.SaveAvatarImg(int64(intAdminId),fname)//待开发fname:存到uploads中的文件名称
		context.JSON(iris.Map{
			"status":utils.RECODE_OK,
			"image_path":fname,
		})
		
	})
}

