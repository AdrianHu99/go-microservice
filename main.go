package main

import (
	"log"
	"microservice/handlers"
	"net/http"
	"os"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product=api", log.LstdFlags)
	hh := handlers.NewHello(l)
	bye := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", bye)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()
}
