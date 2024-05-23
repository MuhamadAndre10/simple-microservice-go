package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)

const webport = ":80"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {

	// try to connect to rabbitmq
	conn, err := connect()
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
		os.Exit(1)
	}

	defer conn.Close()

	app := Config{
		Rabbit: conn,
	}

	log.Println("Starting broker service on port", webport)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s", webport),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
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
