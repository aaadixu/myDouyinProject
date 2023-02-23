package redis

import (
	"douyinProject/chat/cmd/consts"
	myredis "github.com/go-redis/redis"
)

var RedisCli *myredis.Client

func InitRedis() {

	RedisCli = myredis.NewClient(&myredis.Options{
		Addr:     consts.RedisAddr,
		Password: consts.RedisPassword,
		DB:       consts.RedisDB,
	})

}
