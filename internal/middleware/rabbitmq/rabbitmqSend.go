package rabbitmq

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SendError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func SendInit() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	SendError(err, "Failed to connect to RabbitMQ")
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		SendError(err, "Failed to close connection")
	}(conn)

	ch, err := conn.Channel()
	SendError(err, "Failed to open a channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		SendError(err, "Failed to close channel")
	}(ch)

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	SendError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	SendError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
