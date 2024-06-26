package main

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

// This is a simples consumer. Check examples of more reliable consumer in lib https://github.com/rabbitmq/amqp091-go/blob/main/example_client_test.go#L53

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
		"",    // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("creating queue: %s", err)
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

	authorLikes := make(map[int]int)
	go func() {
		for d := range msgs {
			var event ArticleLikeEvent
			err := json.Unmarshal(d.Body, &event)
			if err != nil {
				log.Fatal(err)
			}

			authorLikes[event.AuthorID] += 1
			log.Printf("got ArticleLikeEvent. Author %d has %d likes", event.AuthorID, authorLikes[event.AuthorID])
		}
	}()

	log.Printf(" [*] Waiting for message. To exit press CTRL+C")
	<-forever
}

type ArticleLikeEvent struct {
	ArticleID int
	AuthorID  int
	LikeType  string
}
