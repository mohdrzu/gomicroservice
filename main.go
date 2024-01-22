package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mohdrzu/gomicroservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "gomicroservice-api -> ", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)

	serveMux := http.NewServeMux()
	serveMux.Handle("/hello", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)

	server := &http.Server{
		Addr: ":9090",
		Handler: serveMux,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	

	server.ListenAndServe()
}