package service

import (
	"CmsProject/model"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)

/*
管理员服务
标准的开发模式将每个实体提供的功能以接口标准的形式定义,供控制层进行调用
*/
type AdminService interface {
	//通过管理员用户+密码 获取管理员实体 查询数据库，如果查询到，返回管理员实体并返回true
	//否则返回nil,false
	GetByAdminNameAndPassword(username,paaaword string)(model.Admin,bool)

	//获取管理员总数
	GetAdminCount()(int64,error)

	//查询管理员信息
	GetByAdminId(adminId int64)(model.Admin,bool)

	//保存管理员头像
	SaveAvatarImg(adminId int64,fileName string)bool

	//查询管理员列表
	GetAdminList(offset,limit int)[]*model.Admin
}

//为了方便改动数据库(mysql,oracle等)定义一个函数返回上面的接口，而接口实现在这个函数里面
func NewAdminService(db *xorm.Engine) AdminService {
	return &adminService{
		engine: db,
	}
}

/*
管理员服务真正实现的结构体
*/
type adminService struct {
	engine *xorm.Engine
}

/*
方法:查询管理员总数
*/
func (ac *adminService) GetAdminCount()(int64,error)  {
	count,err := ac.engine.Count(new(model.Admin))
	if err != nil{
		panic(err.Error())
		return 0,nil
	}
	return count,nil
}
/*
方法:通过用户名和密码查询管理员
*/
func (ac *adminService)GetByAdminNameAndPassword(username,paaaword string)(model.Admin,bool)  {
	var admin model.Admin
	ac.engine.Where("admin_name = ? and pwd = ?",username,paaaword).Get(&admin)//字段不对了，卡在这里了
	return admin,admin.AdminId != 0//是否有此对象
}

/**
查询管理员信息
*/

func (ac *adminService)GetByAdminId(adminId int64)(model.Admin,bool)  {
	var admin model.Admin
	ac.engine.Id(adminId).Get(&admin)//将数据库的信息存到admin实体
	return admin,admin.AdminId != 0

}

/**
保存头像信息
*/
func (ac *adminService)SaveAvatarImg(adminId int64,fileName string)bool  {
	admin := model.Admin{Avatar:fileName}
	_,err := ac.engine.Id(adminId).Cols("avatar").Update(&admin)//把admin实体更新到数据库中的avatar列
	return err != nil

}

/**
获取管理员列表
offset:管理员的偏移量
limit:请求管理员的条数
*/
func (ac adminService)GetAdminList(offset,limit int)[]*model.Admin{
	var adminList []*model.Admin

	err := ac.engine.Limit(limit,offset).Find(&adminList)
	if err != nil{
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return adminList
}


