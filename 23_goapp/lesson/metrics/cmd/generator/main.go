package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/joho/godotenv"

	"proj/lessons/23_goapp/lesson/metrics/internal/model"
	"proj/lessons/23_goapp/lesson/metrics/internal/storage"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	sqlDB, err := sql.Open("mysql", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	store := storage.New(sqlDB)

	n := 50000
	for i := 0; i < n; i++ {
		err := store.AddItem(context.Background(), &model.Item{
			Name:  fmt.Sprintf("generated item #%d", i),
			Likes: int(rand.Int63n(1000000)),
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
