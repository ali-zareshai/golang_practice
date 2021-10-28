package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var cache *redis.Client
func ConnectToRedis()  {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cache = redis.NewClient(&redis.Options{
		Addr: os.Getenv("redis_address")+":"+os.Getenv("redis_port"),
		Password: os.Getenv("redis_password"),
		DB: 0,
	})
	fmt.Println("connect to redis")
}

func SetItem(key string ,val string,expireTime time.Duration)  {
	fmt.Println("set redis -> "+key+" :: "+val)
	error := cache.Set(key,val,expireTime).Err()
	if error != nil {
		fmt.Println(error)
	}
}

func GetItem(key string) string {
	val, err :=cache.Get(key).Result()
	fmt.Println("get redis -> "+val)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return val
}