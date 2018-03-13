package db

import (
	"github.com/go-redis/redis"
	"log"
	"for_job/http/lib"
	"fmt"
)

var Cash *redis.Client

func Connect() {
	Cash = redis.NewClient(&redis.Options{
		Addr:     lib.Config.Redis,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := Cash.Ping().Err()
	if err != nil {
		log.Println("Error connect redis: ", err)
	}
}

func Check() bool {
	err := Cash.Ping().Err()
	if err != nil {
		return false
	}
	return true
}
