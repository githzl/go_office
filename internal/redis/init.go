package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"go_office/config"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.YamlFile.Redis.Servers[0],
		Password: config.YamlFile.Redis.Auth,
		DB:       config.YamlFile.Redis.Db,
	})
	_, err := Client.Ping().Result()
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
}
