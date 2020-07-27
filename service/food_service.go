package service

import (
	"CmsProject/model"
	"github.com/go-xorm/xorm"
)

/**
 * 食品food的服务
 */
type FoodService interface {
	//查询食品总数，并返回
	GetFoodCount() (int64,error)
	//查询食品列表并返回
	GetFoodList(offset,limit int) []model.Food
	//查询食品类别列表

}

type foodService struct {
	Engine *xorm.Engine
}

/**
 * 新实例化一个商店模块服务对象结构体
 */
//func NewFoodService(engine *xorm.Engine) FoodService {
//	return &foodService{Engine: engine}
//}


