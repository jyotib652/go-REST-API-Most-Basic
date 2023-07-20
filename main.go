package main

import (
	"fmt"
	"net/http"
)

type Config struct{}

const webPort = "8080"

// Problem: create a rest api

func main() {

	app := Config{}

	fmt.Printf("starting web server on port %s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("Couldn't start the http server")
	}

}
