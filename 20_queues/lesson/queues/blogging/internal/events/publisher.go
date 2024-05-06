package events

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"proj/lessons/20_queues/lesson/queues/blogging/internal/models"
	"sync/atomic"
	"time"
)

const (
	ArticleLikesQueueExchange = "blogging.article_likes"
)

// Important moment about github.com/rabbitmq/amqp091-go lib
// “Things not intended to be supported: … Auto reconnect …”
// https://github.com/rabbitmq/amqp091-go?tab=readme-ov-file#non-goals
// In case of connection error (network or server issues), the connection does not automatically recover.
// Option 1. Implement reconnection in publisher on your own (simplest example bellow)
// Option 2. Use some wrappers/libs to help with reconnects (one example, quite endorsed lib from one of my coworkers https://github.com/makasim/amqpextra)

// Publisher is an example of AMQP message sender with simplest reconnect logic
type Publisher struct {
	amqpURL    string
	connection *amqp.Connection
	channel    *amqp.Channel
	connected  atomic.Bool
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
	// It's important to call close on channel to ensure in-flight messages are completely delivered to server
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

	err := p.publishJSON(event, ArticleLikesQueueExchange)
	if err != nil {
		return fmt.Errorf("publishing article like: %w", err)
	}

	log.Printf("published event: %v", event)

	return nil
}

func (p *Publisher) publishJSON(event any, exchange string) error {
	data, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	//// Simplest logic for reconnects - if connection is absent, try to reconnect
	//if !p.connected.Load() {
	//	err = p.connect()
	//	if err != nil {
	//		return fmt.Errorf("reconnecting to queue: %w", err)
	//	}
	//}
	////

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = p.channel.PublishWithContext(ctx,
		exchange, // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		})
	if err != nil {
		//// Simplest logic for reconnects - mark connection as absent
		//p.connected.Store(false)
		//p.Close()
		////
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

	p.connected.Store(true)

	return nil
}

func (p *Publisher) ensureExchanges() error {
	// ensure exchange present
	// exchange - entry point and router for messages
	err := p.channel.ExchangeDeclare(
		ArticleLikesQueueExchange, // name
		"fanout",                  // type
		true,                      // durable
		false,                     // auto-deleted
		false,                     // internal
		false,                     // no-wait
		nil,                       // arguments
	)
	if err != nil {
		return fmt.Errorf("creating exchange: %w", err)
	}

	return nil
}
