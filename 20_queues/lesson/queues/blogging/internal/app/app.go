package app

import (
	"context"
	"log"
	"math/rand"
	"proj/lessons/20_queues/lesson/queues/blogging/internal/events"
	"proj/lessons/20_queues/lesson/queues/blogging/internal/models"
	"proj/lessons/20_queues/lesson/queues/blogging/internal/services"
	"time"
)

type Config struct {
	AMQPUrl string
}

func Run(ctx context.Context, conf Config) error {
	publisher, err := events.NewPublisher(conf.AMQPUrl)
	if err != nil {
		return err
	}
	defer publisher.Close()

	articlesService := services.NewArticles(publisher)
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		if ctx.Err() != nil {
			return nil
		}

		opCtx := context.Background()
		// This data can come from API request, notification message, etc.
		id := rand.Intn(10)
		article := &models.Article{
			ID:       id,
			Title:    "Title",
			AuthorID: id % 5,
		}
		like := &models.Like{
			UserID: rand.Intn(100),
			Type:   models.LikeTypeThumbUp,
		}

		// this call can be a part of API handler, event consumer, or a batch job, etc.
		err = articlesService.Like(opCtx, article, like)
		if err != nil {
			log.Printf("error on like article: %s", err)
		}
	}

	return nil
}
