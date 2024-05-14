package main

import (
	"fmt"
	"log"
	"net/http"
)

const webport = ":80"

type Config struct {
}

func main() {

	app := Config{}

	log.Println("Starting broker service on port", webport)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s", webport),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
