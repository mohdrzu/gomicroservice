package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}


func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func(g *Goodbye) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	g.l.Printf( "request -> method: %s to: %q", r.Method, r.URL.Path)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Goodbye %s", data)
}