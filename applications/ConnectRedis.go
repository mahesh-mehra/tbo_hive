package applications

import (
	"fmt"
	"tbo_backend/objects"
	"tbo_backend/utils"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func ConnectRedis() bool {

	defer utils.HandlePanic()

	fmt.Println("Initiating Redis connection on : ", objects.ConfigObj.Redis.Host)

	// connecting to redis servers
	rdb = redis.NewClient(&redis.Options{
		Addr:     objects.ConfigObj.Redis.Host,
		Password: objects.ConfigObj.Redis.Password,
		DB:       objects.ConfigObj.Redis.DB,
	})

	fmt.Println("Connected to Redis! on: ", objects.ConfigObj.Redis.Host)

	return true
}

// GetRDB returning redis DB object
func GetRDB() *redis.Client {
	defer utils.HandlePanic()
	return rdb
}
