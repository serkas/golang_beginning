package events

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"proj/lessons/20_queues/lesson/queues/blogging/internal/models"
	"time"
)

type Publisher struct {
	amqpURL    string
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewPublisher(amqpURL string) (*Publisher, error) {
	p := &Publisher{amqpURL: amqpURL}

	err := p.connect()
	if err != nil {
		return nil, fmt.Errorf("connecting AMQP server: %w", err)
	}

	err = p.ensureExchanges()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Publisher) Close() {
	p.channel.Close()
	p.connection.Close()
}

func (p *Publisher) PublishArticleLike(ctx context.Context, article *models.Article, like *models.Like) error {
	event := ArticleLikeDTO{
		ArticleID: article.ID,
		AuthorID:  article.AuthorID,
		UserID:    like.UserID,
		LikeType:  string(like.Type),
	}

	data, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//
	err = p.channel.PublishWithContext(ctx,
		"blogging.article_likes", // exchange
		"",                       // routing key
		false,                    // mandatory
		false,                    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	if err != nil {
		return fmt.Errorf("publishing: %w", err)
	}

	return nil
}

func (p *Publisher) connect() error {
	// connection - network transport abstraction
	conn, err := amqp.Dial(p.amqpURL)
	if err != nil {
		return fmt.Errorf("creating connection: %w", err)
	}
	p.connection = conn

	// channel - more high level abstraction to use AMQP server
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("creating channel: %w", err)
	}
	p.channel = ch

	return nil
}

func (p *Publisher) ensureExchanges() error {
	// ensure exchange present
	// exchange - entry point and router for messages
	err := p.channel.ExchangeDeclare(
		"blogging.article_likes", // name
		"fanout",                 // type
		true,                     // durable
		false,                    // auto-deleted
		false,                    // internal
		false,                    // no-wait
		nil,                      // arguments
	)
	if err != nil {
		return fmt.Errorf("creating exchange: %w", err)
	}

	return nil
}
