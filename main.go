package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var emails = []string{
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
	"qbc1@gmail.com",
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://rabbitmq:rabbitmq@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"email", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, email := range emails {
		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(email),
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", email)
	}

}
