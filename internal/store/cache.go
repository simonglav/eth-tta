package store

import (
	"log"

	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func GetCache(block_number string) ([]byte, error) {
	val, err := client.Get(block_number).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("getcached")
	return []byte(val), nil

}

func SetCache(block_number string, json []byte) {
	err := client.Set(block_number, json, 0).Err()
	if err != nil {
		log.Println(err)
	}
	log.Println("setcached")
}
