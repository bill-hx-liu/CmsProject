package model

/**
用户订单详情结构体
*/
type OrderDetail struct {
	UserOrder 		`xorm:"extends"`// "extends"代表着不仅能访问这些结构体还能访问这些结构体中的字段,
									// 从而可以将数据分别解析到每个表中的每个字段中
	User 			`xorm:"extends"`
	OrderStatus 	`xorm:"extends"`
	Shop			`xorm:"extends"`
	Address			`xorm:"extends"`
}

func (detail *OrderDetail)OrderDetail2Resp() interface{} {
	respDesc := map[string]interface{}{
		"id" :					detail.UserOrder.Id,
		"total_amount":			detail.UserOrder.SumMoney,
		"user_id":				detail.User.UserName, //用户名
		"status":				detail.OrderStatus.StatusDesc,//订单状态
		"restaurant_id":		detail.Shop.ShopId,//商铺Id
		"restaurant_image_url":	detail.Shop.ImagePath,//商铺图片
		"restaurant_name":		detail.Shop.Name,//商铺名称
		"formatted_created_at":	detail.Time,//创建时间
		"status_code":			0,
		"address_id":			detail.Address.AddressId,//订单地址
	}
	statusDesc := map[string]interface{}{
		"color":	"f60",
		"sub_title":"15分钟内支付",
		"title":	detail.OrderStatus.StatusDesc,
	}
	respDesc["status_bar"] = statusDesc
	return respDesc
}