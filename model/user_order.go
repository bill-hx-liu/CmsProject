package model

import "time"

/**
 * 用户订单结构实体定义
 */
type UserOrder struct {
	Id            int64        `xorm:"pk autoincr" json:"id"` //主键
	SumMoney      int64        `xorm:"default 0" json:"sum_money"`
	Time          time.Time    `xorm:"DateTime" json:"time"`         //时间
	//订单信息
	OrderTime     uint64       `json:"order_time"`                   //订单创建时间
	OrderStatusId int64        `xorm:"index" json:"order_status_id"` //订台状态id index:索引 具体状态在order_status.go中定义的主键:StatusId
	OrderStatus   *OrderStatus `xorm:"-"`                            //订单对象 "-" 代表不映射
	//用户信息
	UserId        int64        `xorm:"index" json:"user_id"`         //用户编号Id index:索引 具体定义在user.go文件中
	User          *User        `xorm:"-"`                            //订单对应的账户，并不进行结构体字段映射
	//商铺信息
	ShopId        int64        `xorm:"index" json:"shop_id"`         //用户购买的商品编号 index同上, 在shop.go中
	Shop          *Shop        `xorm:"-"`                            //商品结构体，不进行映射

	AddressId     int64        `xorm:"index" json:"address_id"`      //地址结构体的Id index 同上,见address.go
	Address       *Address     `xorm:"-"`                            //地址结构体，不进行映射

	DelFlag       int64        `xorm:"default 0" json:"del_flag"`    //删除标志 0为正常 1为已删除
}
