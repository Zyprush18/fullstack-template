package redisdb

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(host,port,username,pass string) *redis.Client {
	fmt.Println(username,host,port,pass)

	dbs :=  redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",host,port),
		DB: 0,
		Username: username,
		Password: pass,
	})	

	return dbs
}