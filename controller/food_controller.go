package controller

import (
	"github.com/kataras/iris"
	//"github.com/kataras/iris/mvc"
)

type FoodController struct {
	Ctx iris.Context
}


/**
url:foods/count?
type:Get
desc:获取所有的食品记录总数
*/
//func (fc *FoodController)  GetCount() mvc.Result {
//	iris.New().Logger().Info("食品记录总数")
//	//result,err := fc.
//
//}
