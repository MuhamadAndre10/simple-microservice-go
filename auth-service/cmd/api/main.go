package main

import (
	"database/sql"
	"fmt"
	"github.com/MuhamadAndre10/simple-microservice/auth-service/data"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	// TODO connect to DB

	// set up config
	app := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
