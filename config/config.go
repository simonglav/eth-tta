package config

import (
	"log"

	"github.com/pelletier/go-toml"
)

// Default values
var EtherScanAPIKey = "YourApiKeyToken"

var ServerAddr = "localhost:8080"

var RedisAddr = "localhost:6379"
var RedisPassword = ""
var RedisDB = 0

func init() {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		log.Println(err)
		return
	}
	EtherScanAPIKey = config.Get("EtherScanAPIKey").(string)

	ServerAddr = config.Get("ServerAddr").(string)

	RedisAddr = config.Get("RedisAddr").(string)
	RedisPassword = config.Get("RedisPassword").(string)
	RedisDB = int(config.Get("RedisDB").(int64))
}
