package service

import (
	"CmsProject/model"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)

/**
 * 订单服务接口
 */
type OrderService interface {
	GetCount() (int64,error)
	GetOrderList(offset,limit int) []model.OrderDetail
}

type orderService struct {
	engine *xorm.Engine
}

/*
获取订单列表
*/
func (orderService *orderService)GetOrderList(offset,limit int)[]model.OrderDetail  {
	orderList := make([]model.OrderDetail,0)

	//查询用户订单信息
	//使用了多表查询Join(多表关联查询)(INNER:内关联)
	//将user_order表的字段与其他表的字段关联到一起
	err := orderService.engine.Table("user_order").
		Join("INNER","order_status","order_status.status_id = user_order.order_status_id").
		Join("INNER","user","user.id = user_order.user_id").
		Join("INNER","shop","shop.shop_id = user_order.shop_id").
		Join("INNER","address","address.address_id = user_order.address_id").
		Find(&orderList)//很多数据,所以model中的order_detail.go中的使用"extends"
	iris.New().Logger().Info(orderList[0])
	if err != nil{
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	//iris.New().Logger().Info(orderList)
	return orderList
}