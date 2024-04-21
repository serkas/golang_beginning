package main

import (
	"log"

	"proj/lessons/16_testing/lesson/functional/app"
	"proj/lessons/16_testing/lesson/functional/storage"
)

// curl localhost:8081/items -d '{"id": 1, "name": "item 1"}'
// curl localhost:8081/items
func main() {
	a := app.New(storage.NewMemStorage())

	err := a.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
