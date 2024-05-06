package main

import (
	"context"
	"log"
	"proj/lessons/20_queues/lesson/queues/blogging/internal/app"
)

func main() {
	conf := app.Config{
		AMQPUrl: "amqp://guest:guest@localhost:5672/",
	}
	err := app.Run(context.Background(), conf)
	if err != nil {
		log.Fatal(err)
	}

	//simple publishing example
	//connection - network transport abstraction
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	//if err != nil {
	//	log.Fatalf("creating connection: %s", err)
	//}
	//defer conn.Close()
	//
	//// channel - more high level abstraction to use AMQP server
	//ch, err := conn.Channel()
	//if err != nil {
	//	log.Fatalf("creating channel: %s", err)
	//}
	//defer ch.Close()
	//
	//// exchange - entry point and router for messages
	//err = ch.ExchangeDeclare(
	//	"blogging.article_likes", // name
	//	"fanout",                 // type
	//	true,                     // durable
	//	false,                    // auto-deleted
	//	false,                    // internal
	//	false,                    // no-wait
	//	nil,                      // arguments
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//body := "123"
	//err = ch.PublishWithContext(ctx,
	//	"blogging.article_likes", // exchange
	//	"",                       // routing key
	//	false,                    // mandatory
	//	false,                    // immediate
	//	amqp.Publishing{
	//		ContentType: "text/plain",
	//		Body:        []byte(body),
	//	})
	//if err != nil {
	//	log.Fatal(err)
	//}
}
