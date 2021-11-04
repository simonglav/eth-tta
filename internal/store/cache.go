package store

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/total-transactions-amount-eth/config"
)

var client = redis.NewClient(&redis.Options{
	Addr:     config.RedisAddr,
	Password: config.RedisPassword,
	DB:       config.RedisDB,
})

// GetCache is getting value of block_number.
// If the key does not exist nil is returned
func GetCache(block_number string) ([]byte, error) {
	val, err := client.Get(block_number).Result()
	if err != nil {
		if err != redis.Nil {
			log.Println(err)
		}
		return nil, err
	}
	return []byte(val), nil

}

// SetCache is setting block_number to hold the json value
// If key already holds a value, it is overwritten
func SetCache(block_number string, json []byte) {
	err := client.Set(block_number, json, 0).Err()
	if err != nil {
		log.Println(err)
	}
}
