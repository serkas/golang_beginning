package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	"proj/lessons/23_goapp/lesson/metrics/internal/app"
)

func main() {
	conf := getConfig()

	a, err := app.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("starting server at %s", conf.ServerAddress)
	if err := a.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func getConfig() app.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	// Also we can use a lib similar to this to parse envs directly into tagged struct
	// https://github.com/kelseyhightower/envconfig

	return app.Config{
		ServerAddress: os.Getenv("SERVER_ADDR"),
		DB:            os.Getenv("DB_DSN"),
		RedisAddress:  os.Getenv("REDIS_ADDR"),
	}
}
