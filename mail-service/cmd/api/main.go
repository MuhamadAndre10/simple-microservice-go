package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Config struct {
	Mailer Mail
}

const (
	webPort = "80"
)

func main() {
	app := Config{
		Mailer: CreateMail(),
	}

	log.Println("Starting mail-service on port", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panicf("server failed to start: %v", err)
	}
}

func CreateMail() Mail {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	m := Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        port,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Encryption:  os.Getenv("MAIL_ENCRYPTION"),
		FromAddress: os.Getenv("MAIL_ADDRESS"),
		FromName:    os.Getenv("MAIL_NAME"),
	}

	//m := Mail{
	//	Domain:      "localhost",
	//	Host:        "mailpit",
	//	Port:        1025,
	//	Username:    "",
	//	Password:    "",
	//	Encryption:  "tls",
	//	FromAddress: "andrepriyanto95@gmail.com",
	//	FromName:    "john due",
	//}

	return m
}
