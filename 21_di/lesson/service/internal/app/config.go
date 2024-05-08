package app

import (
	"os"
	"strings"
)

type Config struct {
	DB            string
	RedisAddress  string
	ServerAddress string
}

func ReadConfig() *Config {
	return &Config{
		ServerAddress: readEnvStr("server_address", ":8081"),
		DB:            readEnvStr("mysql_user", "root:root@tcp(localhost:23306)/items_db"),
		RedisAddress:  readEnvStr("redis_address", "localhost:6379"),
	}
}

func readEnvStr(name, defaultValue string) string {
	val, exists := os.LookupEnv(strings.ToUpper(name))
	if exists {
		return val
	}

	return defaultValue
}
