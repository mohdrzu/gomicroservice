package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mohdrzu/gomicroservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "gomicroservice-api -> ", log.LstdFlags)
	productHandler := handlers.NewProduct(l)

	serveMux := http.NewServeMux()
	serveMux.Handle("/product", productHandler)

	server := &http.Server{
		Addr: ":9090",
		Handler: serveMux,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	
	go func ()  {
		l.Println("starting server on port 9090")

		err := server.ListenAndServe()
		if err != nil {
			l.Printf("error starting server at: %s\n", err)
			os.Exit(1)
		}	
	}()
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <- sigChan 
	l.Println("received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	
	server.Shutdown(tc)
}