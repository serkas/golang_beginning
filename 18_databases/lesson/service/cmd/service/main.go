package main

import (
	"context"
	"log"

	"proj/lessons/18_databases/lesson/service/app"
)

// curl localhost:8081/items -d '{"id": 1, "name": "item 1"}'
// curl localhost:8081/items
func main() {
	conf := app.Config{
		ServerAddress: ":8081",
		DB:            "root:root@tcp(localhost:13306)/items_db",
	}
	a, err := app.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("starting server at %s", conf)
	if err := a.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
