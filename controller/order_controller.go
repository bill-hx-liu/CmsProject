package controller

import (
	"CmsProject/service"
	"CmsProject/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"strconv"
)

type OrderController struct {
	Ctx iris.Context
	Service service.OrderService
	Session *sessions.Session
}

/**
获取订单列表
*/
func (orderController *OrderController) Get() mvc.Result{
	iris.New().Logger().Info("查询订单列表")

	offsetStr := orderController.Ctx.FormValue("offset")
	limitStr :=orderController.Ctx.FormValue("limit")
	var offset int
	var limit int

	//判断offset,limit两个变量任意一个都不能为空
	if offsetStr == "" || limitStr == ""{
		return mvc.Response{
			Object:map[string]interface{}{
				"status":utils.RECODE_FAIL,
				"type":utils.RESPMSG_ERROR_ORDERLIST,
				"message":utils.Recode2Text(utils.RESPMSG_ERROR_ORDERLIST),
			},
		}
	}

	offset,err := strconv.Atoi(offsetStr)
	limit,err = strconv.Atoi(limitStr)
	if err != nil{
		return mvc.Response{
			Object:map[string]interface{}{
				"status":utils.RECODE_FAIL,
				"type":utils.RESPMSG_ERROR_ORDERLIST,
				"message":utils.Recode2Text(utils.RESPMSG_ERROR_ORDERLIST),
			},
		}
	}

	//做页数的限制
	if offset <= 0{
		offset = 0
	}

	//做页数最大限制
	if limit > MaxLimit{
		limit = MaxLimit
	}

	//
	orderList := orderController.Service.GetOrderList(offset,limit)//

	//
	if len(orderList) == 0{
		return  mvc.Response{
			Object:map[string]interface{}{
				"status":utils.RECODE_FAIL,
				"type":utils.RESPMSG_ERROR_ORDERLIST,
				"message":utils.Recode2Text(utils.RESPMSG_ERROR_ORDERLIST),
			},
		}
	}

	//否则将查询到的用户数据进行转换，变成前端需要的内容
	var respList []interface{}
	for _,detail := range orderList{
		respList = append(respList,detail.OrderDetail2Resp())//装填成前端需要的内容
	}

	//返回用户列表
	return mvc.Response{
		Object:&respList,
	}
}

/**
查询订单记录总数
待开发
*/
