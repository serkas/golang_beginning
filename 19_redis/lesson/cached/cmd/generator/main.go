package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"proj/lessons/19_redis/lesson/cached/internal/model"
	"proj/lessons/19_redis/lesson/cached/internal/storage"
)

func main() {
	dbDSN := "root:root@tcp(localhost:23306)/items_db"

	sqlDB, err := sql.Open("mysql", dbDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	store := storage.New(sqlDB)

	n := 100000
	for i := 0; i < n; i++ {
		err := store.AddItem(context.Background(), &model.Item{
			Name:  fmt.Sprintf("generated item #%d", i),
			Views: int(rand.Int63n(100000)),
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}
