package xjredis

import (
	"github.com/go-redis/redis"
	"github.com/xjieinfo/xjgo/xjcore/xjtypes"
	"log"
	"testing"
	"time"
)

var redisdb *redis.Client

func init() {
	conf := xjtypes.Redis{
		Addr:     "127.0.0.1:6379",
		Password: "",
		Db:       1,
	}
	redisdb = RedisInit(conf)
}
func Test_setnx(t *testing.T) {
	ok, err := redisdb.SetNX("ORDER_CANCEL:"+"102021072910440002", 1, 120*time.Second).Result()
	if err != nil {
		log.Println(err)
	}
	if ok {
		log.Println("set ok.")
	} else {
		log.Println("set fail.")
	}
}

func Test_lock(t *testing.T) {
	ok := Lock(redisdb, "ORDER_CANCEL:"+"102021072910440002", time.Now().Format("2006-01-02 15:04:05"), 120)
	log.Printf("set %t \n.", ok)
}
