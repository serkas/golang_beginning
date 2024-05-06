package main

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	// connection - network transport abstraction
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("creating connection: %s", err)
	}
	defer conn.Close()

	// channel - more high level abstraction to use AMQP server
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("creating channel: %s", err)
	}
	defer ch.Close()

	// ensure consuming queue exists
	q, err := ch.QueueDeclare(
		"notification.article_likes", // name
		true,                         // durable
		false,                        // delete when unused
		false,                        // exclusive
		false,                        // no-wait
		nil,                          // arguments
	)
	if err != nil {
		log.Fatalf("creating channel: %s", err)
	}

	err = ch.QueueBind(
		q.Name,                   // queue name
		"",                       // routing key
		"blogging.article_likes", // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("binding channel: %s", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("registering consumer: %s", err)
	}
	var forever chan struct{}

	go func() {
		for d := range msgs {
			var event ArticleLikeEvent
			err := json.Unmarshal(d.Body, &event)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("got ArticleLikeEvent. Send push notification to author %d about like from %d", event.AuthorID, event.UserID)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

type ArticleLikeEvent struct {
	ArticleID int
	AuthorID  int
	UserID    int
	LikeType  string
}
