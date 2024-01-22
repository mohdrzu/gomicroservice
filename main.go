package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mohdrzu/gomicroservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "gomicroservice-api -> ", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)

	serveMux := http.NewServeMux()
	serveMux.Handle("/hello", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)

	
	http.ListenAndServe(":9090", serveMux)
}