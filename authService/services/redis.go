package services

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	_RedisCluster *redis.ClusterClient
)

func RedisCluster() *redis.ClusterClient {
	if _RedisCluster == nil {
		cli := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: []string{
				"47.243.46.134:6379",
				//"127.0.0.1:7001",
				//"127.0.0.1:7002",
				//"127.0.0.1:7003",
				//"127.0.0.1:7004",
				//"127.0.0.1:7005",
			},
			Password: "nieaowei",
		})
		cmd := cli.Ping(context.Background())
		if cmd.Err() != nil {
			log.Panicln(cmd.Err())
		}
		log.Println(cmd.Result())
		_RedisCluster = cli
	}
	return _RedisCluster
}

var (
	_RedisClient *redis.Client
)

func RedisClient() *redis.Client {
	if _RedisCluster == nil {
		cli := redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
		})
		cmd := cli.Ping(context.Background())
		if cmd.Err() != nil {
			log.Panicln(cmd.Err())
		}
		log.Println(cmd.Result())
		_RedisClient = cli
	}
	return _RedisClient
}
