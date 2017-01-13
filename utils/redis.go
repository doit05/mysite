package utils

import (
	"github.com/astaxie/beego"
	"mysite/helper"
	"mysite/redis"
)

var Redis *redis.Redis

func init() {
	redisCache := redis.NewRedisCache()

	err := redisCache.Connect(beego.AppConfig.String(helper.GetConfigPrifix() + "redisconn"))

	if err == nil {
		Redis = redisCache
	} else {
		panic("redis连接失败: " + err.Error())
	}
}
