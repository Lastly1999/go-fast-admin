package redis

import (
	"fast-admin-service/global"
	"fast-admin-service/pkg/setting"
	"fmt"
	"github.com/go-redis/redis"
)

var (
	Rdb *redis.Client
)

// Setup 初始化redis连接池
func Setup() {
	rsn := fmt.Sprintf("%v:%v", setting.RedisOptions.HOST, setting.RedisOptions.PORT)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     rsn,
		Password: setting.RedisOptions.PassWord,
	})
	_, err := Rdb.Ping().Result()
	if err != nil {
		global.ZAP_LOG.Info(err.Error())
		return
	}
	global.ZAP_LOG.Info("go-redis content succcess...")
}
