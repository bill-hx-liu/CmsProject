package service

import (
	"CmsProject/model"
	"github.com/go-xorm/xorm"
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
	ac.engine.Where("admin = ? and pwd = ?",username,paaaword).Get(&admin)
	return admin,admin.AdminId != 0//是否有此对象
}


