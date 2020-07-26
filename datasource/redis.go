package datasource

import (
	"CmsProject/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions/sessiondb/redis"
)

/**
返回Redis实例
*/
func NewRedis() *redis.Database  {
	var database *redis.Database

	//项目配置
	cmsConfig := config.InitConfig()//这里没获取到 原因是文件的绝对路径错了,已改正
	//fmt.Println(cmsConfig)
	if cmsConfig != nil{
		iris.New().Logger().Info("hello")
		rd := cmsConfig.Redis//为空,已改正
		iris.New().Logger().Info(rd)
		database = redis.New(redis.Config{//最新版本的iris直接定义在了redis下的databse.go下，直接使用即可
			Network:   rd.NetWork,
			Addr:      rd.Addr + ":" + rd.Port,
			Password:  rd.Password,
			Database:  "",
			MaxActive: 10,
			Timeout:   redis.DefaultRedisTimeout,
			Prefix:    rd.Prefix,
		})
	}else{
		iris.New().Logger().Info("hello err... in redis.go")
	}
	return database
}
