package service

import (
	"CmsProject/model"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)
//待开发完成


type CategoryService interface {
	//添加食品
	//AddCategory(model *model.FoodCategory) bool

	//GetCategoryByShopId(shopid int64)([]model.FoodCategory,error)

	//GetAllCategory()([]model.FoodCategory,error)

	//获取商铺信息操作
	//GetRestaurantInfo(shop_id int64) (model.Shop,error)

	//保存记录
	//SaveFood(food model.Food) bool
	SaveShop(shop model.Shop) bool

	//删除记录操作
	DeleteShop(restaurantId int) bool
	DeleteFood(foodId int) bool
}

//为了方便改动数据库(mysql,oracle等)定义一个函数返回上面的接口，而接口实现在这个函数里面
func NewCategoryService(db *xorm.Engine) CategoryService {
	return &categoryService{
		engine: db,
	}
}

//结构体
type categoryService struct {
	engine *xorm.Engine
}

/**
保存食品记录
*/
//func (cs *categoryService) SaveFood(food model.Food) bool  {
//	_,err := cs.engine.Insert(&food)
//	return err == nil
//
//}


/**
保存商铺记录
*/
func (cs *categoryService) SaveShop(shop model.Shop)bool  {
	_,err := cs.engine.Insert(&shop)
	if err != nil{
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}

/**
删除商铺 软删除,修改删除字段0到1代表删除
*/
func (cs *categoryService) DeleteShop(restaurantId int) bool {
	shop := model.Shop{ShopId:restaurantId,Dele:1}
	_,err := cs.engine.Where("shop_id = ?",restaurantId).Cols("dele").Update(&shop)//只更新某列的数据
	if err != nil{
		iris.New().Logger().Info(err.Error())
	}
	return  err == nil
}

/**
删除食品记录
*/
func (cs *categoryService) DeleteFood(foodId int) bool  {
	food := model.Food{Id:int64(foodId),DelFlag:1}
	_, err := cs.engine.Where("id = ?",foodId).Cols("del_flag").Update(&food)
	if err != nil{
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}