package services

import (
	"context"
	"fmt"
	"proj/lessons/20_queues/lesson/queues/blogging/internal/models"
)

type ArticleEventPublisher interface {
	PublishArticleLike(ctx context.Context, article *models.Article, like *models.Like) error
}

type Articles struct {
	publisher ArticleEventPublisher
}

func NewArticles(p ArticleEventPublisher) *Articles {
	return &Articles{
		publisher: p,
	}
}

func (a *Articles) Like(ctx context.Context, article *models.Article, like *models.Like) error {
	// store like to DB

	// Notify other systems about like
	err := a.publisher.PublishArticleLike(ctx, article, like)
	if err != nil {
		return fmt.Errorf("publishing like: %w", err)
	}

	return nil
}
