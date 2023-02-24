package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
}

const PORT = "80"

func main() {
	app := Config{}

	log.Println("starting mail service on port: %s", PORT)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}