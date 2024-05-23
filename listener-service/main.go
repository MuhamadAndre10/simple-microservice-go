package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"listener/event"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	// try to connect to rabbitmq
	conn, err := connect()
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
		os.Exit(1)
	}

	defer conn.Close()

	// start listen the messages
	log.Println("Listening for and consuming RabbitMQ messages...")

	// create consumer
	consumer, err := event.NewCunsumer(conn)
	if err != nil {
		panic(err)
	}

	// watch the queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.ERROR", "log.WARNING"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {

	var counter int64
	var Backoff = 1 * time.Second
	var conn *amqp.Connection

	for {
		dial, err := amqp.Dial("amqp://rabbitmq:pass@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counter++
		} else {
			conn = dial
			log.Println("Connected to RabbitMQ")
			break
		}

		if counter > 5 {
			fmt.Println(err)
			return nil, err
		}

		Backoff = time.Duration(math.Pow(float64(counter), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(Backoff)
		continue
	}

	return conn, nil
}
