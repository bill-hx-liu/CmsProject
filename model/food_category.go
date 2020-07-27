package model
/**
食品种类结构体定义
*/
type FoodCategory struct {
	Id				int64 `xorm:"pk autoincr" json:"id"`//食品的id
	CategoryName	string`json:"name"`//食品种类的名称
	CatrgoryDesc	string`json:"description"`//食品种类描述
	Level			int64`json:"level"`//食品种类的层级
	ParentCategoryID int64`json:"parent_category_id"`//父一级的类型id
	RestaurantId	int64`xorm:"index" json:"restaurant_id"`//所对应的商铺ID
}
