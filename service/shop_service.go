package service

import (
	"CmsProject/model"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)

/**
 * 商店Shop的服务
 */
type ShopService interface {
	//查询商店总数，并返回
	GetShopCount() (int64,error)
	GetShopList(offset,limit int) []model.Shop
}

type shopService struct {
	Engine *xorm.Engine
}

/**
 * 新实例化一个商店模块服务对象结构体
 */
func NewShopService(engine *xorm.Engine) ShopService {
	return &shopService{Engine: engine}
}

/**
 * 查询商店的总数然后返回
 */
func (ss *shopService) GetShopCount() (int64,error) {
	result,err := ss.Engine.Where("dele = 0").Count(new(model.Shop))//软删除,不是真正的删除而是将"dele"置为1
	return result,err
}
/**
获取到商铺列表信息
*/
func (ss *shopService) GetShopList(offset,limit int) []model.Shop {
	shopList := make([]model.Shop,0)

	err := ss.Engine.Where("dele = 0").Limit(limit,offset).Find(&shopList)//limit:查询多少条,offset:起始位置
	if err != nil{
		iris.New().Logger().Info(err.Error())
	}
	return shopList
}