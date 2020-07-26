package controller

import (
	"CmsProject/service"
	"CmsProject/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"strings"
)

/**
 * 统计功能控制者
 */
type StatisController struct {
	//上下文环境对象
	Ctx iris.Context

	//统计功能的服务实现接口
	Service service.StatisService

	//session
	Session *sessions.Session//临时缓存的机制
}

var (
	ADMINMODULE = "ADMIN_"
	USERMODULE  = "USER_"
	ORDERMODULE = "ORDER_"
)


/**
 * 解析统计功能路由请求
 */
func (sc *StatisController) GetCount() mvc.Result {
	// /statis/user/2019-03-10/count
	path := sc.Ctx.Path()

	var pathSlice []string
	if path != "" {
		pathSlice = strings.Split(path, "/")
	}

	//不符合请求格式
	if len(pathSlice) != 5 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
			},
		}
	}

	//将最前面的去掉
	pathSlice = pathSlice[1:]
	model := pathSlice[1]
	date := pathSlice[2]
	var result int64
	switch model {
	case "user":
		//加上缓存机制(session(redis))
		//因为当日的增长可能会变,但过去七天的数据是不会改变的，所以可以使用session + redis缓存机制(适用于改变小的数据)
		userResult := sc.Session.Get(USERMODULE + date)
		if userResult != nil{//先查询redis缓存中的数据是否有,有则返回无则查询mysql数据库并写入redis
			userResult = userResult.(float64)//语法:类型断言
			return mvc.Response{
				Object:map[string]interface{}{
					"status":utils.RECODE_OK,
					"count":userResult,
				},
			}
		}else {
			iris.New().Logger().Error(date) //时间
			//将不在session(redis)中的缓存数据从mysql读出来，并写入缓存中
			result = sc.Service.GetUserDailyCount(date)
			//读入缓存中
			//数据转变:
			//raw: date:2020-04-02
			//to:(USERMODULE)USER_2020-04-02:result(key:value)
			sc.Session.Set(USERMODULE + date,result)
		}
	case "order":
		orderStatis := sc.Session.Get(ORDERMODULE + date)
		if orderStatis != nil{
			orderStatis = orderStatis.(float64)
			return mvc.Response{
				Object:map[string]interface{}{
					"status":utils.RECODE_OK,
					"count":orderStatis,
				},
			}
		}else {
			result = sc.Service.GetOrderDailyCount(date)
			sc.Session.Set(ORDERMODULE + date,result)
		}
	case "admin":
		adminStatis := sc.Session.Get(ADMINMODULE + date)
		if adminStatis != nil{
			adminStatis = adminStatis.(float64)
			return mvc.Response{
				Object:map[string]interface{}{
					"status":utils.RECODE_OK,
					"count":adminStatis,
				},
			}
		}else{
			result = sc.Service.GetAdminDailyCount(date)
			sc.Session.Set(ADMINMODULE + date,result)
		}
	}
	//必须有这一步返回,假设缓存中没有会走到这一步
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  result,
		},
	}
}
