package handlers

import (
	"log"
	"net/http"

	"github.com/mohdrzu/gomicroservice/models"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func(p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request){
	productList := models.GetProducts()
	err := productList.ToJson(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}