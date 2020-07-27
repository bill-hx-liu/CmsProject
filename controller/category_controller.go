package controller

import (
	"CmsProject/model"
	"CmsProject/service"
	"CmsProject/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
)

/**
食品种类控制器
*/
type CategoryController struct {
	Ctx iris.Context
	Service service.CategoryService
}
/**
添加食品种类实体
*/
type CategoryEntity struct {
	Name 			string `json:"name"`
	Description		string	`json:"description"`
	RestaurantId	string`json:"restaurant_id"`
}
/**
在控制器处理http请求之前执行的方法,绑定方法
*/
func (cc *CategoryController) BeforeActivation(a mvc.BeforeActivation){
	//通过商铺Id获取对应的食品种类
	//a.Handle("GET","/getcategory/{shopId}","GetCategoryByShopId")

	//获取全部的食品种类
	//a.Handle("GET","/v2/restaurant/category","GetAllCategory")

	//添加商铺记录
	a.Handle("POST","/addShop","PostAddShop")

	//删除商铺记录
	a.Handle("DELETE","/restaurant/{restaurant_id}","DeleteRestaurant")

	//删除食品记录
	a.Handle("DELETE","/v2/food/{food_id}","DeleteFood")

	////获取某个商铺的信息
	//a.Handle("GET","/restaurant/{restaurant_id}","GetRestaurantInfo")
}

/**
添加商铺方法
url:/shopping/addshop
type:Post
desc:添加商铺记录
*/
func (cc *CategoryController)PostAddShop() mvc.Result  {
	iris.New().Logger().Info("PostAddShop,添加商铺记录")

	var shop model.Shop
	err := cc.Ctx.ReadJSON(&shop)
	iris.New().Logger().Info(shop)

	if err != nil{
		iris.New().Logger().Info(err.Error())

		cc.Ctx.Request()
		return mvc.Response{
			Object:map[string]interface{}{
				"status":utils.RECODE_FAIL,
				"message":utils.Recode2Text(utils.RESPMSG_FAIL_ADDREST),
			},
		}
	}

	//添加
	saveShop := cc.Service.SaveShop(shop)
	if !saveShop{//错误不为空，则处理错误
		return mvc.Response{
			Object:map[string]interface{}{
				"status":utils.RECODE_FAIL,
				"message":utils.Recode2Text(utils.RESPMSG_FAIL_ADDREST),
			},
		}

	}

	return mvc.Response{
		Object:map[string]interface{}{
			"status":utils.RECODE_OK,
			"message":utils.Recode2Text(utils.RESPMSG_SUCCESS_ADDREST),
			"shopDetail":shop,
		},
	}

}

/**
删除商铺记录
*/
func (cc *CategoryController) DeleteRestaurant() mvc.Result{
	restaurant_id := cc.Ctx.Params().Get("restaurant_id")
	shopId,err := strconv.Atoi(restaurant_id)

	if err != nil{
		return mvc.Response{
			Object:map[string]interface{}{
				"status":utils.RECODE_FAIL,
				"type":utils.RESPMSG_HASNOACCESS,
				"message":utils.Recode2Text(utils.RESPMSG_HASNOACCESS),
			},
		}
	}

	dele := cc.Service.DeleteShop(shopId)
	if !dele{//错误不为空
		return mvc.Response{
			Object:map[string]interface{}{
				"status":utils.RECODE_FAIL,
				"type":utils.RESPMSG_HASNOACCESS,
				"message":utils.Recode2Text(utils.RESPMSG_HASNOACCESS),
			},
		}

	}else{
		return mvc.Response{
			Object:map[string]interface{}{
				"status":utils.RECODE_OK,
				"type":utils.RESPMSG_SUCCESS_DELETESHOP,
				"message":utils.Recode2Text(utils.RESPMSG_SUCCESS_DELETESHOP),
			},
		}
	}
}

/**
删除食品记录
*/
func (cc *CategoryController)DeleteFood() mvc.Result  {
	food_id := cc.Ctx.Params().Get("food_id")

	foodID,err := strconv.Atoi(food_id)
	if err != nil{
		return mvc.Response{
			Object:map[string]interface{}{
				"status":utils.RECODE_FAIL,
				"type":utils.RESPMSG_HASNOACCESS,
				"message":utils.Recode2Text(utils.RESPMSG_HASNOACCESS),
			},
		}
	}
	dele := cc.Service.DeleteFood(foodID)
	if !dele{
		return mvc.Response{
			Object:map[string]interface{}{
				"status":utils.RECODE_FAIL,
				"type":utils.RESPMSG_HASNOACCESS,
				"message":utils.Recode2Text(utils.RESPMSG_HASNOACCESS),
			},
		}
	}else{
		return mvc.Response{
			Object:map[string]interface{}{
				"status":utils.RECODE_OK,
				"type":utils.RESPMSG_SUCCESS_FOODDELE,
				"message":utils.Recode2Text(utils.RESPMSG_SUCCESS_FOODDELE),
			},
		}
	}

}