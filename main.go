package main

import (
	"log"
	"microservice/handlers"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product=api", log.LstdFlags)
	hh := handlers.NewHello(l)
	bye := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", bye)

	http.ListenAndServe(":9090", sm)
}
