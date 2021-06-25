package xjredis

import (
	"gitee.com/xjieinfo/xjgo/xjcore/xjtypes"
	"github.com/go-redis/redis"
	"log"
	"os"
)

func RedisInit(conf xjtypes.Redis) *redis.Client {
	Redisdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.Db,
	})
	//心跳
	pong, err := Redisdb.Ping().Result()
	if err != nil {
		log.Println("连接redis出错了")
		log.Println(err)
		os.Exit(0)
	}
	log.Println("连接redis...", pong, err) // Output: PONG <nil>
	return Redisdb
}
