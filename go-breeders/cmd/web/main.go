package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct{}

func main() {

	//declare a variable "app" of type "application"
	app := application{}

	//Handler function for any requests for the "/" resource
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello Werld!")
		})
	*/

	srv := &http.Server{

		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting web application on port ", port)

	//err := http.ListenAndServe(port, nil)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
