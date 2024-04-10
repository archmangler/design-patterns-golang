package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
}

type appConfig struct {
	useCache bool
}

func main() {

	//declare a variable "app" of type "application"
	app := application{
		templateMap: make(map[string]*template.Template),
		config: appConfig{
			useCache: true,
		},
	}

	fmt.Println("DEBUG> Using the cache: ", app.config.useCache, " <DEBUG")

	flag.BoolVar(&app.config.useCache, "cache", true, "Use template cache")

	flag.Parse()

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
